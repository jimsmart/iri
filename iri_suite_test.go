package iri_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIri(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IRI Suite")
}
