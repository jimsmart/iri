package iri

// Unique IRIs for File Formats.
//
// For further information, see http://www.w3.org/ns/formats/
const (
	// Unique IRIs for File Formats.
	FORMAT_NS                  = "http://www.w3.org/ns/formats/"
	FORMAT_JSON_LD             = FORMAT_NS + "JSON-LD"
	FORMAT_N3                  = FORMAT_NS + "N3"
	FORMAT_N_Triples           = FORMAT_NS + "N-Triples"
	FORMAT_N_Quads             = FORMAT_NS + "N-Quads"
	FORMAT_LD_Patch            = FORMAT_NS + "LD_Patch"
	FORMAT_microdata           = FORMAT_NS + "microdata"
	FORMAT_OWL_XML             = FORMAT_NS + "OWL_XML"
	FORMAT_OWL_Functional      = FORMAT_NS + "OWL_Functional"
	FORMAT_OWL_Manchester      = FORMAT_NS + "OWL_Manchester"
	FORMAT_POWDER              = FORMAT_NS + "POWDER"
	FORMAT_POWDER_S            = FORMAT_NS + "POWDER-S"
	FORMAT_PROV_N              = FORMAT_NS + "PROV-N"
	FORMAT_PROV_XML            = FORMAT_NS + "PROV-XML"
	FORMAT_RDFa                = FORMAT_NS + "RDFa"
	FORMAT_RDF_JSON            = FORMAT_NS + "RDF_JSON"
	FORMAT_RDF_XML             = FORMAT_NS + "RDF_XML"
	FORMAT_RIF_XML             = FORMAT_NS + "RIF_XML"
	FORMAT_SPARQL_Results_XML  = FORMAT_NS + "SPARQL_Results_XML"
	FORMAT_SPARQL_Results_JSON = FORMAT_NS + "SPARQL_Results_JSON"
	FORMAT_SPARQL_Results_CSV  = FORMAT_NS + "SPARQL_Results_CSV"
	FORMAT_SPARQL_Results_TSV  = FORMAT_NS + "SPARQL_Results_TSV"
	FORMAT_Turtle              = FORMAT_NS + "Turtle"
	FORMAT_TriG                = FORMAT_NS + "TriG"

	// FORMAT_TriX IRI = FORMAT_NS + "TriX"
)

// TODO(js) Should format.go include TriX?
