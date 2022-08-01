package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GuideDependencyType is documented here http://hl7.org/fhir/ValueSet/guide-dependency-type
type GuideDependencyType int

const (
	GuideDependencyTypeReference GuideDependencyType = iota
	GuideDependencyTypeInclusion
)

func (code GuideDependencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GuideDependencyType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "reference":
		*code = GuideDependencyTypeReference
	case "inclusion":
		*code = GuideDependencyTypeInclusion
	default:
		return fmt.Errorf("unknown GuideDependencyType code `%s`", s)
	}
	return nil
}
func (code GuideDependencyType) String() string {
	return code.Code()
}
func (code GuideDependencyType) Code() string {
	switch code {
	case GuideDependencyTypeReference:
		return "reference"
	case GuideDependencyTypeInclusion:
		return "inclusion"
	}
	return "<unknown>"
}
func (code GuideDependencyType) Display() string {
	switch code {
	case GuideDependencyTypeReference:
		return "Reference"
	case GuideDependencyTypeInclusion:
		return "Inclusion"
	}
	return "<unknown>"
}
func (code GuideDependencyType) Definition() string {
	switch code {
	case GuideDependencyTypeReference:
		return "The guide is referred to by URL."
	case GuideDependencyTypeInclusion:
		return "The guide is embedded in this guide when published."
	}
	return "<unknown>"
}
