package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// VisionEyes is documented here http://hl7.org/fhir/ValueSet/vision-eye-codes
type VisionEyes int

const (
	VisionEyesRight VisionEyes = iota
	VisionEyesLeft
)

func (code VisionEyes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *VisionEyes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "right":
		*code = VisionEyesRight
	case "left":
		*code = VisionEyesLeft
	default:
		return fmt.Errorf("unknown VisionEyes code `%s`", s)
	}
	return nil
}
func (code VisionEyes) String() string {
	return code.Code()
}
func (code VisionEyes) Code() string {
	switch code {
	case VisionEyesRight:
		return "right"
	case VisionEyesLeft:
		return "left"
	}
	return "<unknown>"
}
func (code VisionEyes) Display() string {
	switch code {
	case VisionEyesRight:
		return "Right Eye"
	case VisionEyesLeft:
		return "Left Eye"
	}
	return "<unknown>"
}
func (code VisionEyes) Definition() string {
	switch code {
	case VisionEyesRight:
		return "Right Eye"
	case VisionEyesLeft:
		return "Left Eye"
	}
	return "<unknown>"
}
