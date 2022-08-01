package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConsentExceptType is documented here http://hl7.org/fhir/ValueSet/consent-except-type
type ConsentExceptType int

const (
	ConsentExceptTypeDeny ConsentExceptType = iota
	ConsentExceptTypePermit
)

func (code ConsentExceptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentExceptType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "deny":
		*code = ConsentExceptTypeDeny
	case "permit":
		*code = ConsentExceptTypePermit
	default:
		return fmt.Errorf("unknown ConsentExceptType code `%s`", s)
	}
	return nil
}
func (code ConsentExceptType) String() string {
	return code.Code()
}
func (code ConsentExceptType) Code() string {
	switch code {
	case ConsentExceptTypeDeny:
		return "deny"
	case ConsentExceptTypePermit:
		return "permit"
	}
	return "<unknown>"
}
func (code ConsentExceptType) Display() string {
	switch code {
	case ConsentExceptTypeDeny:
		return "Opt Out"
	case ConsentExceptTypePermit:
		return "Opt In"
	}
	return "<unknown>"
}
func (code ConsentExceptType) Definition() string {
	switch code {
	case ConsentExceptTypeDeny:
		return "Consent is denied for actions meeting these rules"
	case ConsentExceptTypePermit:
		return "Consent is provided for actions meeting these rules"
	}
	return "<unknown>"
}
