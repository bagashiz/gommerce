package usecase_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/bagashiz/gommerce/internal/app/category/domain"
	"github.com/bagashiz/gommerce/internal/app/category/domain/mock"
	"github.com/bagashiz/gommerce/internal/app/category/usecase"
)

var _ = Describe("CategoryUsecase", func() {
	var (
		mockCtrl        *gomock.Controller
		categoryRepo    *mock.MockCategoryRepository
		categoryUsecase domain.CategoryUsecase
		ctx             context.Context
		category        *domain.Category
		categories      []domain.Category
		page, limit     int
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		categoryRepo = mock.NewMockCategoryRepository(mockCtrl)
		categoryUsecase = usecase.New(categoryRepo)

		ctx = context.Background()

		page = 1
		limit = 5

		category = &domain.Category{
			ID:   1,
			Name: "Category 1",
		}

		categories = []domain.Category{
			{
				ID:   1,
				Name: "Category 1",
			},
			{
				ID:   2,
				Name: "Category 2",
			},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Create", func() {
		It("should create a new category", func() {
			Expect(categoryUsecase).ShouldNot(BeNil())

			categoryRepo.EXPECT().Create(ctx, category).Return(nil)

			err := categoryUsecase.Create(ctx, category)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("GetAll", func() {
		It("should return all categories", func() {
			Expect(categoryUsecase).ShouldNot(BeNil())

			categoryRepo.EXPECT().GetAll(ctx, page, limit).Return(categories, nil)

			res, err := categoryUsecase.GetAll(ctx, page, limit)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(categories))
		})
	})

	Describe("GetByID", func() {
		It("should return a category by ID", func() {
			Expect(categoryUsecase).ShouldNot(BeNil())

			categoryRepo.EXPECT().GetByID(ctx, category.ID).Return(category, nil)

			res, err := categoryUsecase.GetByID(ctx, category.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(category))
		})
	})

	Describe("Update", func() {
		It("should update a category", func() {
			Expect(categoryUsecase).ShouldNot(BeNil())

			categoryRepo.EXPECT().GetByID(ctx, category.ID).Return(category, nil)
			categoryRepo.EXPECT().Update(ctx, category).Return(nil)

			err := categoryUsecase.Update(ctx, category)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Delete", func() {
		It("should delete a category", func() {
			Expect(categoryUsecase).ShouldNot(BeNil())

			categoryRepo.EXPECT().GetByID(ctx, category.ID).Return(category, nil)
			categoryRepo.EXPECT().Delete(ctx, category.ID).Return(nil)

			err := categoryUsecase.Delete(ctx, category.ID)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
