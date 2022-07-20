package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructureDefinitionKind is documented here http://hl7.org/fhir/ValueSet/structure-definition-kind
type StructureDefinitionKind int

const (
	StructureDefinitionKindPrimitiveType StructureDefinitionKind = iota
	StructureDefinitionKindComplexType
	StructureDefinitionKindResource
	StructureDefinitionKindLogical
)

func (code StructureDefinitionKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureDefinitionKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "primitive-type":
		*code = StructureDefinitionKindPrimitiveType
	case "complex-type":
		*code = StructureDefinitionKindComplexType
	case "resource":
		*code = StructureDefinitionKindResource
	case "logical":
		*code = StructureDefinitionKindLogical
	default:
		return fmt.Errorf("unknown StructureDefinitionKind code `%s`", s)
	}
	return nil
}
func (code StructureDefinitionKind) String() string {
	return code.Code()
}
func (code StructureDefinitionKind) Code() string {
	switch code {
	case StructureDefinitionKindPrimitiveType:
		return "primitive-type"
	case StructureDefinitionKindComplexType:
		return "complex-type"
	case StructureDefinitionKindResource:
		return "resource"
	case StructureDefinitionKindLogical:
		return "logical"
	}
	return "<unknown>"
}
func (code StructureDefinitionKind) Display() string {
	switch code {
	case StructureDefinitionKindPrimitiveType:
		return "Primitive Data Type"
	case StructureDefinitionKindComplexType:
		return "Complex Data Type"
	case StructureDefinitionKindResource:
		return "Resource"
	case StructureDefinitionKindLogical:
		return "Logical Model"
	}
	return "<unknown>"
}
func (code StructureDefinitionKind) Definition() string {
	switch code {
	case StructureDefinitionKindPrimitiveType:
		return "A primitive type that has a value and an extension. These can be used throughout Resource and extension definitions. Only the base specification can define primitive types."
	case StructureDefinitionKindComplexType:
		return "A  complex structure that defines a set of data elements. These can be used throughout Resource and extension definitions, and in logical models."
	case StructureDefinitionKindResource:
		return "A resource defined by the FHIR specification."
	case StructureDefinitionKindLogical:
		return "A conceptual package of data that will be mapped to resources for implementation."
	}
	return "<unknown>"
}
