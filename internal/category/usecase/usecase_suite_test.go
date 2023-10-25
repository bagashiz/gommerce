package usecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCategoryUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Category Usecase Suite")
}
