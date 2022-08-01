package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ContentType is documented here http://hl7.org/fhir/ValueSet/content-type
type ContentType int

const (
	ContentTypeXml ContentType = iota
	ContentTypeJson
	ContentTypeTtl
	ContentTypeNone
)

func (code ContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContentType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "xml":
		*code = ContentTypeXml
	case "json":
		*code = ContentTypeJson
	case "ttl":
		*code = ContentTypeTtl
	case "none":
		*code = ContentTypeNone
	default:
		return fmt.Errorf("unknown ContentType code `%s`", s)
	}
	return nil
}
func (code ContentType) String() string {
	return code.Code()
}
func (code ContentType) Code() string {
	switch code {
	case ContentTypeXml:
		return "xml"
	case ContentTypeJson:
		return "json"
	case ContentTypeTtl:
		return "ttl"
	case ContentTypeNone:
		return "none"
	}
	return "<unknown>"
}
func (code ContentType) Display() string {
	switch code {
	case ContentTypeXml:
		return "xml"
	case ContentTypeJson:
		return "json"
	case ContentTypeTtl:
		return "ttl"
	case ContentTypeNone:
		return "none"
	}
	return "<unknown>"
}
func (code ContentType) Definition() string {
	switch code {
	case ContentTypeXml:
		return "XML content-type corresponding to the application/fhir+xml mime-type."
	case ContentTypeJson:
		return "JSON content-type corresponding to the application/fhir+json mime-type."
	case ContentTypeTtl:
		return "RDF content-type corresponding to the text/turtle mime-type."
	case ContentTypeNone:
		return "Prevent the use of the corresponding http header."
	}
	return "<unknown>"
}
