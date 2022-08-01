package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructureMapInputMode is documented here http://hl7.org/fhir/ValueSet/map-input-mode
type StructureMapInputMode int

const (
	StructureMapInputModeSource StructureMapInputMode = iota
	StructureMapInputModeTarget
)

func (code StructureMapInputMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapInputMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "source":
		*code = StructureMapInputModeSource
	case "target":
		*code = StructureMapInputModeTarget
	default:
		return fmt.Errorf("unknown StructureMapInputMode code `%s`", s)
	}
	return nil
}
func (code StructureMapInputMode) String() string {
	return code.Code()
}
func (code StructureMapInputMode) Code() string {
	switch code {
	case StructureMapInputModeSource:
		return "source"
	case StructureMapInputModeTarget:
		return "target"
	}
	return "<unknown>"
}
func (code StructureMapInputMode) Display() string {
	switch code {
	case StructureMapInputModeSource:
		return "Source Instance"
	case StructureMapInputModeTarget:
		return "Target Instance"
	}
	return "<unknown>"
}
func (code StructureMapInputMode) Definition() string {
	switch code {
	case StructureMapInputModeSource:
		return "Names an input instance used a source for mapping"
	case StructureMapInputModeTarget:
		return "Names an instance that is being populated"
	}
	return "<unknown>"
}
