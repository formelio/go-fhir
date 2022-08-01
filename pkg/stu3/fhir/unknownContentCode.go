package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// UnknownContentCode is documented here http://hl7.org/fhir/ValueSet/unknown-content-code
type UnknownContentCode int

const (
	UnknownContentCodeNo UnknownContentCode = iota
	UnknownContentCodeExtensions
	UnknownContentCodeElements
	UnknownContentCodeBoth
)

func (code UnknownContentCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *UnknownContentCode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "no":
		*code = UnknownContentCodeNo
	case "extensions":
		*code = UnknownContentCodeExtensions
	case "elements":
		*code = UnknownContentCodeElements
	case "both":
		*code = UnknownContentCodeBoth
	default:
		return fmt.Errorf("unknown UnknownContentCode code `%s`", s)
	}
	return nil
}
func (code UnknownContentCode) String() string {
	return code.Code()
}
func (code UnknownContentCode) Code() string {
	switch code {
	case UnknownContentCodeNo:
		return "no"
	case UnknownContentCodeExtensions:
		return "extensions"
	case UnknownContentCodeElements:
		return "elements"
	case UnknownContentCodeBoth:
		return "both"
	}
	return "<unknown>"
}
func (code UnknownContentCode) Display() string {
	switch code {
	case UnknownContentCodeNo:
		return "Neither Elements or Extensions"
	case UnknownContentCodeExtensions:
		return "Unknown Extensions"
	case UnknownContentCodeElements:
		return "Unknown Elements"
	case UnknownContentCodeBoth:
		return "Unknown Elements and Extensions"
	}
	return "<unknown>"
}
func (code UnknownContentCode) Definition() string {
	switch code {
	case UnknownContentCodeNo:
		return "The application does not accept either unknown elements or extensions."
	case UnknownContentCodeExtensions:
		return "The application accepts unknown extensions, but not unknown elements."
	case UnknownContentCodeElements:
		return "The application accepts unknown elements, but not unknown extensions."
	case UnknownContentCodeBoth:
		return "The application accepts unknown elements and extensions."
	}
	return "<unknown>"
}
