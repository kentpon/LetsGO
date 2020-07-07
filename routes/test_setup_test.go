package routes

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var Router *gin.Engine

func TestSetup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
	Router = Setup()
	fmt.Println("Test router setup")
})

var _ = AfterSuite(func() {
	fmt.Println("Test cleanup")
})
