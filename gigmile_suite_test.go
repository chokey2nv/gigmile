package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGigmile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gigmile Suite")
}
