package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PropertyType is documented here http://hl7.org/fhir/ValueSet/concept-property-type
type PropertyType int

const (
	PropertyTypeCode PropertyType = iota
	PropertyTypeCoding
	PropertyTypeString
	PropertyTypeInteger
	PropertyTypeBoolean
	PropertyTypeDateTime
)

func (code PropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *PropertyType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "code":
		*code = PropertyTypeCode
	case "Coding":
		*code = PropertyTypeCoding
	case "string":
		*code = PropertyTypeString
	case "integer":
		*code = PropertyTypeInteger
	case "boolean":
		*code = PropertyTypeBoolean
	case "dateTime":
		*code = PropertyTypeDateTime
	default:
		return fmt.Errorf("unknown PropertyType code `%s`", s)
	}
	return nil
}
func (code PropertyType) String() string {
	return code.Code()
}
func (code PropertyType) Code() string {
	switch code {
	case PropertyTypeCode:
		return "code"
	case PropertyTypeCoding:
		return "Coding"
	case PropertyTypeString:
		return "string"
	case PropertyTypeInteger:
		return "integer"
	case PropertyTypeBoolean:
		return "boolean"
	case PropertyTypeDateTime:
		return "dateTime"
	}
	return "<unknown>"
}
func (code PropertyType) Display() string {
	switch code {
	case PropertyTypeCode:
		return "code (internal reference)"
	case PropertyTypeCoding:
		return "Coding (external reference)"
	case PropertyTypeString:
		return "string"
	case PropertyTypeInteger:
		return "integer"
	case PropertyTypeBoolean:
		return "boolean"
	case PropertyTypeDateTime:
		return "dateTime"
	}
	return "<unknown>"
}
func (code PropertyType) Definition() string {
	switch code {
	case PropertyTypeCode:
		return "The property value is a code that identifies a concept defined in the code system"
	case PropertyTypeCoding:
		return "The property  value is a code defined in an external code system. This may be used for translations, but is not the intent"
	case PropertyTypeString:
		return "The property value is a string"
	case PropertyTypeInteger:
		return "The property value is a string (often used to assign ranking values to concepts for supporting score assessments)"
	case PropertyTypeBoolean:
		return "The property value is a boolean true | false"
	case PropertyTypeDateTime:
		return "The property is a date or a date + time"
	}
	return "<unknown>"
}
