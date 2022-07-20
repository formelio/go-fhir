package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AllergyIntoleranceCriticality is documented here http://hl7.org/fhir/ValueSet/allergy-intolerance-criticality
type AllergyIntoleranceCriticality int

const (
	AllergyIntoleranceCriticalityLow AllergyIntoleranceCriticality = iota
	AllergyIntoleranceCriticalityHigh
	AllergyIntoleranceCriticalityUnableToAssess
)

func (code AllergyIntoleranceCriticality) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceCriticality) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "low":
		*code = AllergyIntoleranceCriticalityLow
	case "high":
		*code = AllergyIntoleranceCriticalityHigh
	case "unable-to-assess":
		*code = AllergyIntoleranceCriticalityUnableToAssess
	default:
		return fmt.Errorf("unknown AllergyIntoleranceCriticality code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceCriticality) String() string {
	return code.Code()
}
func (code AllergyIntoleranceCriticality) Code() string {
	switch code {
	case AllergyIntoleranceCriticalityLow:
		return "low"
	case AllergyIntoleranceCriticalityHigh:
		return "high"
	case AllergyIntoleranceCriticalityUnableToAssess:
		return "unable-to-assess"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCriticality) Display() string {
	switch code {
	case AllergyIntoleranceCriticalityLow:
		return "Low Risk"
	case AllergyIntoleranceCriticalityHigh:
		return "High Risk"
	case AllergyIntoleranceCriticalityUnableToAssess:
		return "Unable to Assess Risk"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCriticality) Definition() string {
	switch code {
	case AllergyIntoleranceCriticalityLow:
		return "Worst case result of a future exposure is not assessed to be life-threatening or having high potential for organ system failure."
	case AllergyIntoleranceCriticalityHigh:
		return "Worst case result of a future exposure is assessed to be life-threatening or having high potential for organ system failure."
	case AllergyIntoleranceCriticalityUnableToAssess:
		return "Unable to assess the worst case result of a future exposure."
	}
	return "<unknown>"
}
