package usecase_test

import (
	"github.com/bagashiz/gommerce/internal/app/city/domain"
	"github.com/bagashiz/gommerce/internal/app/city/domain/mock"
	"github.com/bagashiz/gommerce/internal/app/city/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("CityUsecase", func() {
	var (
		mockCtrl    *gomock.Controller
		cityRepo    *mock.MockCityRepository
		cityUsecase domain.CityUsecase
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		cityRepo = mock.NewMockCityRepository(mockCtrl)
		cityUsecase = usecase.New(cityRepo)
		Expect(cityUsecase).ShouldNot(BeNil())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetAll", func() {
		cities := []domain.City{
			{
				ID:         "1101",
				ProvinceID: "11",
				Name:       "City 1",
			},
			{
				ID:         "1102",
				ProvinceID: "11",
				Name:       "City 2",
			},
		}

		When("valid province ID parameter", func() {
			It("should return all cities", func() {
				reqParam := "11"

				cityRepo.EXPECT().GetAll(reqParam).Return(cities, nil)

				result, err := cityUsecase.GetAll(reqParam)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(cities))
			})
		})
	})

	Describe("GetByID", func() {
		When("valid province and city ID parameters", func() {
			city := &domain.City{
				ID:         "1101",
				ProvinceID: "11",
				Name:       "City 1",
			}

			It("should return a city by ID", func() {
				provID := "11"
				cityID := "1101"

				cityRepo.EXPECT().GetByID(cityID).Return(city, nil)

				result, err := cityUsecase.GetByID(provID, cityID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(city))
			})
		})

		When("city ID length is less than 2", func() {
			It("should return data not found error", func() {
				provID := "11"
				cityID := "1"

				result, err := cityUsecase.GetByID(provID, cityID)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
				Expect(result).Should(BeNil())
			})
		})

		When("province ID is not equal to city ID[:2]", func() {
			It("should return data not found error", func() {

				provID := "11"
				cityID := "1201"

				result, err := cityUsecase.GetByID(provID, cityID)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
				Expect(result).Should(BeNil())
			})
		})
	})
})
