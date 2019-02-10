# iri

iri is a [Go](https://golang.org) package providing IRI constants for common RDF namespaces.

More than 99% of this package is machine-generated from ontologies.

Remember: Go doesn't include unreferenced consts in the build output.

Note: the full set of consts for DBpedia and Schema.org have been effectively 'commented out' (their files have been renamed to `_gen_dbo.go` and `_gen_schema.go`). This is because when they were included in the package they slowed my text editor's autocomplete to a crawl. 

## Installation
```bash
$ go get github.com/jimsmart/iri
```

```go
import "github.com/jimsmart/iri"
```

## Example

```go
import "github.com/jimsmart/iri"

x := iri.FOAF_knows
y := iri.RDFS_Class
z := iri.DCE_title
a := iri.RDF_type
```

## License

Package iri is copyright 2017-2019 by Jim Smart and released under the [MIT License](LICENSE.md)
