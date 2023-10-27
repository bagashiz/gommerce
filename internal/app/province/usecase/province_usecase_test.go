package usecase_test

import (
	"github.com/bagashiz/gommerce/internal/app/province/domain"
	"github.com/bagashiz/gommerce/internal/app/province/domain/mock"
	"github.com/bagashiz/gommerce/internal/app/province/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("ProvinceUsecase", func() {
	var (
		mockCtrl        *gomock.Controller
		provinceRepo    *mock.MockProvinceRepository
		provinceUsecase domain.ProvinceUsecase
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		provinceRepo = mock.NewMockProvinceRepository(mockCtrl)
		provinceUsecase = usecase.New(provinceRepo)
		Expect(provinceUsecase).ShouldNot(BeNil())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetAll", func() {
		provinces := []domain.Province{
			{
				ID:   "11",
				Name: "Province 1",
			},
			{
				ID:   "12",
				Name: "Province 2",
			},
		}

		It("should return all provinces", func() {
			provinceRepo.EXPECT().GetAll().Return(provinces, nil)

			result, err := provinceUsecase.GetAll()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(provinces))
		})
	})

	Describe("GetByID", func() {
		When("valid province ID parameters", func() {
			province := &domain.Province{
				ID:   "11",
				Name: "Province 1",
			}

			It("should return a province by ID", func() {
				id := "11"

				provinceRepo.EXPECT().GetByID(id).Return(province, nil)

				result, err := provinceUsecase.GetByID(id)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(province))
			})
		})

		When("data does not exist", func() {
			It("should return data not found error", func() {
				id := "69420"

				provinceRepo.EXPECT().GetByID(id).Return(nil, helper.ErrDataNotFound)

				result, err := provinceUsecase.GetByID(id)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
				Expect(result).Should(BeNil())
			})
		})
	})
})
