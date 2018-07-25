package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCoffeed(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coffeed Suite")
}
