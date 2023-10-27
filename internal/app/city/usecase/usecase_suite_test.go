package usecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCityUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "City Usecase Suite")
}
