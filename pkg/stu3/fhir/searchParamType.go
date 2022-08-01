package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchParamType is documented here http://hl7.org/fhir/ValueSet/search-param-type
type SearchParamType int

const (
	SearchParamTypeNumber SearchParamType = iota
	SearchParamTypeDate
	SearchParamTypeString
	SearchParamTypeToken
	SearchParamTypeReference
	SearchParamTypeComposite
	SearchParamTypeQuantity
	SearchParamTypeUri
)

func (code SearchParamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SearchParamType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "number":
		*code = SearchParamTypeNumber
	case "date":
		*code = SearchParamTypeDate
	case "string":
		*code = SearchParamTypeString
	case "token":
		*code = SearchParamTypeToken
	case "reference":
		*code = SearchParamTypeReference
	case "composite":
		*code = SearchParamTypeComposite
	case "quantity":
		*code = SearchParamTypeQuantity
	case "uri":
		*code = SearchParamTypeUri
	default:
		return fmt.Errorf("unknown SearchParamType code `%s`", s)
	}
	return nil
}
func (code SearchParamType) String() string {
	return code.Code()
}
func (code SearchParamType) Code() string {
	switch code {
	case SearchParamTypeNumber:
		return "number"
	case SearchParamTypeDate:
		return "date"
	case SearchParamTypeString:
		return "string"
	case SearchParamTypeToken:
		return "token"
	case SearchParamTypeReference:
		return "reference"
	case SearchParamTypeComposite:
		return "composite"
	case SearchParamTypeQuantity:
		return "quantity"
	case SearchParamTypeUri:
		return "uri"
	}
	return "<unknown>"
}
func (code SearchParamType) Display() string {
	switch code {
	case SearchParamTypeNumber:
		return "Number"
	case SearchParamTypeDate:
		return "Date/DateTime"
	case SearchParamTypeString:
		return "String"
	case SearchParamTypeToken:
		return "Token"
	case SearchParamTypeReference:
		return "Reference"
	case SearchParamTypeComposite:
		return "Composite"
	case SearchParamTypeQuantity:
		return "Quantity"
	case SearchParamTypeUri:
		return "URI"
	}
	return "<unknown>"
}
func (code SearchParamType) Definition() string {
	switch code {
	case SearchParamTypeNumber:
		return "Search parameter SHALL be a number (a whole number, or a decimal)."
	case SearchParamTypeDate:
		return "Search parameter is on a date/time. The date format is the standard XML format, though other formats may be supported."
	case SearchParamTypeString:
		return "Search parameter is a simple string, like a name part. Search is case-insensitive and accent-insensitive. May match just the start of a string. String parameters may contain spaces."
	case SearchParamTypeToken:
		return "Search parameter on a coded element or identifier. May be used to search through the text, displayname, code and code/codesystem (for codes) and label, system and key (for identifier). Its value is either a string or a pair of namespace and value, separated by a \"|\", depending on the modifier used."
	case SearchParamTypeReference:
		return "A reference to another resource."
	case SearchParamTypeComposite:
		return "A composite search parameter that combines a search on two values together."
	case SearchParamTypeQuantity:
		return "A search parameter that searches on a quantity."
	case SearchParamTypeUri:
		return "A search parameter that searches on a URI (RFC 3986)."
	}
	return "<unknown>"
}
