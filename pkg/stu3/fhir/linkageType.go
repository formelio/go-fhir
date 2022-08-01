package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// LinkageType is documented here http://hl7.org/fhir/ValueSet/linkage-type
type LinkageType int

const (
	LinkageTypeSource LinkageType = iota
	LinkageTypeAlternate
	LinkageTypeHistorical
)

func (code LinkageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *LinkageType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "source":
		*code = LinkageTypeSource
	case "alternate":
		*code = LinkageTypeAlternate
	case "historical":
		*code = LinkageTypeHistorical
	default:
		return fmt.Errorf("unknown LinkageType code `%s`", s)
	}
	return nil
}
func (code LinkageType) String() string {
	return code.Code()
}
func (code LinkageType) Code() string {
	switch code {
	case LinkageTypeSource:
		return "source"
	case LinkageTypeAlternate:
		return "alternate"
	case LinkageTypeHistorical:
		return "historical"
	}
	return "<unknown>"
}
func (code LinkageType) Display() string {
	switch code {
	case LinkageTypeSource:
		return "Source of truth"
	case LinkageTypeAlternate:
		return "Alternate record"
	case LinkageTypeHistorical:
		return "Historical/obsolete record"
	}
	return "<unknown>"
}
func (code LinkageType) Definition() string {
	switch code {
	case LinkageTypeSource:
		return "The record represents the \"source of truth\" (from the perspective of this Linkage resource) for the underlying event/condition/etc."
	case LinkageTypeAlternate:
		return "The record represents the alternative view of the underlying event/condition/etc.  The record may still be actively maintained, even though it is not considered to be the source of truth."
	case LinkageTypeHistorical:
		return "The record represents an obsolete record of the underlyng event/condition/etc.  It is not expected to be actively maintained."
	}
	return "<unknown>"
}
