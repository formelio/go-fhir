package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructureMapContextType is documented here http://hl7.org/fhir/ValueSet/map-context-type
type StructureMapContextType int

const (
	StructureMapContextTypeType StructureMapContextType = iota
	StructureMapContextTypeVariable
)

func (code StructureMapContextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapContextType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "type":
		*code = StructureMapContextTypeType
	case "variable":
		*code = StructureMapContextTypeVariable
	default:
		return fmt.Errorf("unknown StructureMapContextType code `%s`", s)
	}
	return nil
}
func (code StructureMapContextType) String() string {
	return code.Code()
}
func (code StructureMapContextType) Code() string {
	switch code {
	case StructureMapContextTypeType:
		return "type"
	case StructureMapContextTypeVariable:
		return "variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Display() string {
	switch code {
	case StructureMapContextTypeType:
		return "Type"
	case StructureMapContextTypeVariable:
		return "Variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Definition() string {
	switch code {
	case StructureMapContextTypeType:
		return "The context specifies a type"
	case StructureMapContextTypeVariable:
		return "The context specifies a variable"
	}
	return "<unknown>"
}
