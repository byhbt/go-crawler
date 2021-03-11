package controllers_test

import (
	"testing"

	"github.com/byhbt/go-crawler/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
