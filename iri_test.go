package iri_test

import (
	"github.com/jimsmart/iri"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Temporarily removed from the project - including these tests
// just seems to make it shout even more about the lack of coverage.
// They're machine-generated constants, dammit - how do I get coverage for consts?
// Do I really also have to generate tests that reference each and every one of them? :/

var _ = Describe("IRI", func() {
	Context("Sanity check for badly generated files", func() {
		It("should define various constants", func() {
			Expect(iri.RDF_NS).To(Equal("http://www.w3.org/1999/02/22-rdf-syntax-ns#"))
			Expect(iri.RDF_type).To(Equal("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"))
			Expect(iri.RDFS_NS).To(Equal("http://www.w3.org/2000/01/rdf-schema#"))
			Expect(iri.RDFS_Class).To(Equal("http://www.w3.org/2000/01/rdf-schema#Class"))
			Expect(iri.OWL_NS).To(Equal("http://www.w3.org/2002/07/owl#"))
			Expect(iri.OWL_Class).To(Equal("http://www.w3.org/2002/07/owl#Class"))
		})
	})
})
