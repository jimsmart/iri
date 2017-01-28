package iri

// The following is a lightweight replacement for:
//  gen_dbo.go
//  gen_schema.go
//
// - which, when included in the package, are so big
// that they slow down autocomplete in Sublime Text.
//
// Dunno if the autocomplete tool can be improved,
// or if this is something in ST - or if such files
// are simply too big for autocomplete, no matter what?

const (
	DBO_NS    = "http://dbpedia.org/ontology/"
	SCHEMA_NS = "http://schema.org/"
)
