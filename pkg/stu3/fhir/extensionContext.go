package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ExtensionContext is documented here http://hl7.org/fhir/ValueSet/extension-context
type ExtensionContext int

const (
	ExtensionContextResource ExtensionContext = iota
	ExtensionContextDatatype
	ExtensionContextExtension
)

func (code ExtensionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExtensionContext) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "resource":
		*code = ExtensionContextResource
	case "datatype":
		*code = ExtensionContextDatatype
	case "extension":
		*code = ExtensionContextExtension
	default:
		return fmt.Errorf("unknown ExtensionContext code `%s`", s)
	}
	return nil
}
func (code ExtensionContext) String() string {
	return code.Code()
}
func (code ExtensionContext) Code() string {
	switch code {
	case ExtensionContextResource:
		return "resource"
	case ExtensionContextDatatype:
		return "datatype"
	case ExtensionContextExtension:
		return "extension"
	}
	return "<unknown>"
}
func (code ExtensionContext) Display() string {
	switch code {
	case ExtensionContextResource:
		return "Resource"
	case ExtensionContextDatatype:
		return "Datatype"
	case ExtensionContextExtension:
		return "Extension"
	}
	return "<unknown>"
}
func (code ExtensionContext) Definition() string {
	switch code {
	case ExtensionContextResource:
		return "The context is all elements matching a particular resource element path."
	case ExtensionContextDatatype:
		return "The context is all nodes matching a particular data type element path (root or repeating element) or all elements referencing a particular primitive data type (expressed as the datatype name)."
	case ExtensionContextExtension:
		return "The context is a particular extension from a particular profile, a uri that identifies the extension definition."
	}
	return "<unknown>"
}
