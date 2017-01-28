package iri

// TODO(js) XSD constants list is incomplete, it currently covers only the basics — see http://www.w3.org/TR/xmlschema-2/

// The RDF-compatible XML Schema datatypes.
const (
	// The RDF-compatible XML Schema datatypes.
	XSD_NS                 = "http://www.w3.org/2001/XMLSchema#"
	XSD_NCName             = XSD_NS + "NCName"
	XSD_NMTOKEN            = XSD_NS + "NMTOKEN"
	XSD_Name               = XSD_NS + "Name"
	XSD_anyURI             = XSD_NS + "anyURI"
	XSD_base64Binary       = XSD_NS + "base64Binary"
	XSD_boolean            = XSD_NS + "boolean"
	XSD_byte               = XSD_NS + "byte"
	XSD_date               = XSD_NS + "date"
	XSD_dateTime           = XSD_NS + "dateTime"
	XSD_dateTimeStamp      = XSD_NS + "dateTimeStamp"
	XSD_dayTimeDuration    = XSD_NS + "dayTimeDuration"
	XSD_decimal            = XSD_NS + "decimal"
	XSD_double             = XSD_NS + "double"
	XSD_duration           = XSD_NS + "duration"
	XSD_float              = XSD_NS + "float"
	XSD_gDay               = XSD_NS + "gDay"
	XSD_gMonth             = XSD_NS + "gMonth"
	XSD_gMonthDay          = XSD_NS + "gMonthDay"
	XSD_gYear              = XSD_NS + "gYear"
	XSD_gYearMonth         = XSD_NS + "gYearMonth"
	XSD_hexBinary          = XSD_NS + "hexBinary"
	XSD_int                = XSD_NS + "int"
	XSD_integer            = XSD_NS + "integer"
	XSD_language           = XSD_NS + "language"
	XSD_long               = XSD_NS + "long"
	XSD_negativeInteger    = XSD_NS + "negativeInteger"
	XSD_nonNegativeInteger = XSD_NS + "nonNegativeInteger"
	XSD_nonPositiveInteger = XSD_NS + "nonPositiveInteger"
	XSD_normalizedString   = XSD_NS + "normalizedString"
	XSD_positiveInteger    = XSD_NS + "positiveInteger"
	XSD_short              = XSD_NS + "short"
	XSD_string             = XSD_NS + "string"
	XSD_time               = XSD_NS + "time"
	XSD_token              = XSD_NS + "token"
	XSD_unsignedByte       = XSD_NS + "unsignedByte"
	XSD_unsignedInt        = XSD_NS + "unsignedInt"
	XSD_unsignedLong       = XSD_NS + "unsignedLong"
	XSD_unsignedShort      = XSD_NS + "unsignedShort"
	XSD_yearMonthDuration  = XSD_NS + "yearMonthDuration"
)
