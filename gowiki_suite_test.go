package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGowiki(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gowiki Suite")
}
