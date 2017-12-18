package main_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/kellimohr/gowiki"
)

var _ = Describe("Gowiki", func() {

	var (
		page Page
	)

	BeforeEach(func() {
		page.Title = "Test"
		page.Body = []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	})
})
