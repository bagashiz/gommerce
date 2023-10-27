package usecase_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/bagashiz/gommerce/internal/app/category/domain"
	"github.com/bagashiz/gommerce/internal/app/category/domain/mock"
	"github.com/bagashiz/gommerce/internal/app/category/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

var _ = Describe("CategoryUsecase", func() {
	var (
		mockCtrl        *gomock.Controller
		categoryRepo    *mock.MockCategoryRepository
		category        *domain.Category
		categoryUsecase domain.CategoryUsecase
		ctx             context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		categoryRepo = mock.NewMockCategoryRepository(mockCtrl)
		categoryUsecase = usecase.New(categoryRepo)
		Expect(categoryUsecase).ShouldNot(BeNil())

		ctx = context.Background()

		category = &domain.Category{
			ID:   1,
			Name: "Category 1",
		}

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Create", func() {
		When("valid category data", func() {
			It("should create a new category", func() {
				categoryRepo.EXPECT().Create(ctx, category).Return(nil)

				err := categoryUsecase.Create(ctx, category)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("duplicate category data", func() {
			It("should return data already exists error", func() {
				categoryRepo.EXPECT().Create(ctx, category).Return(helper.ErrDataAlreadyExists)

				err := categoryUsecase.Create(ctx, category)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataAlreadyExists))
			})
		})
	})

	Describe("GetAll", func() {
		When("valid page and limit parameter", func() {
			page := 1
			limit := 5

			categories := []domain.Category{
				{
					ID:   1,
					Name: "Category 1",
				},
				{
					ID:   2,
					Name: "Category 2",
				},
			}

			It("should return all categories", func() {
				categoryRepo.EXPECT().GetAll(ctx, page, limit).Return(categories, nil)

				res, err := categoryUsecase.GetAll(ctx, page, limit)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res).To(Equal(categories))
			})
		})
	})

	Describe("GetByID", func() {
		When("valid category ID parameter", func() {
			id := uint(1)

			It("should return a category by ID", func() {
				categoryRepo.EXPECT().GetByID(ctx, id).Return(category, nil)

				res, err := categoryUsecase.GetByID(ctx, id)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res).To(Equal(category))
				Expect(res.ID).To(Equal(category.ID))
			})
		})

		When("data does not exist", func() {
			It("should return data not found error", func() {
				id := uint(1)

				categoryRepo.EXPECT().GetByID(ctx, id).Return(nil, helper.ErrDataNotFound)

				res, err := categoryUsecase.GetByID(ctx, id)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
				Expect(res).Should(BeNil())
			})
		})
	})

	Describe("Update", func() {
		When("valid category data", func() {
			id := uint(1)

			It("should update a category", func() {
				categoryRepo.EXPECT().GetByID(ctx, id).Return(category, nil)
				categoryRepo.EXPECT().Update(ctx, category).Return(nil)

				err := categoryUsecase.Update(ctx, category)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("data does not exist", func() {
			It("should return data not found error", func() {
				id := uint(1)

				categoryRepo.EXPECT().GetByID(ctx, id).Return(nil, helper.ErrDataNotFound)

				err := categoryUsecase.Update(ctx, category)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
			})
		})
	})

	Describe("Delete", func() {
		When("valid category ID", func() {
			It("should delete a category", func() {
				categoryRepo.EXPECT().GetByID(ctx, category.ID).Return(category, nil)
				categoryRepo.EXPECT().Delete(ctx, category.ID).Return(nil)

				err := categoryUsecase.Delete(ctx, category.ID)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("data does not exist", func() {
			It("should return data not found error", func() {
				id := uint(1)

				categoryRepo.EXPECT().GetByID(ctx, id).Return(nil, helper.ErrDataNotFound)

				err := categoryUsecase.Delete(ctx, id)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(Equal(helper.ErrDataNotFound))
			})
		})
	})
})
