package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AllergyIntoleranceClinicalStatus is documented here http://hl7.org/fhir/ValueSet/allergy-clinical-status
type AllergyIntoleranceClinicalStatus int

const (
	AllergyIntoleranceClinicalStatusActive AllergyIntoleranceClinicalStatus = iota
	AllergyIntoleranceClinicalStatusInactive
	AllergyIntoleranceClinicalStatusResolved
)

func (code AllergyIntoleranceClinicalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceClinicalStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = AllergyIntoleranceClinicalStatusActive
	case "inactive":
		*code = AllergyIntoleranceClinicalStatusInactive
	case "resolved":
		*code = AllergyIntoleranceClinicalStatusResolved
	default:
		return fmt.Errorf("unknown AllergyIntoleranceClinicalStatus code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceClinicalStatus) String() string {
	return code.Code()
}
func (code AllergyIntoleranceClinicalStatus) Code() string {
	switch code {
	case AllergyIntoleranceClinicalStatusActive:
		return "active"
	case AllergyIntoleranceClinicalStatusInactive:
		return "inactive"
	case AllergyIntoleranceClinicalStatusResolved:
		return "resolved"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceClinicalStatus) Display() string {
	switch code {
	case AllergyIntoleranceClinicalStatusActive:
		return "Active"
	case AllergyIntoleranceClinicalStatusInactive:
		return "Inactive"
	case AllergyIntoleranceClinicalStatusResolved:
		return "Resolved"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceClinicalStatus) Definition() string {
	switch code {
	case AllergyIntoleranceClinicalStatusActive:
		return "An active record of a risk of a reaction to the identified substance."
	case AllergyIntoleranceClinicalStatusInactive:
		return "An inactivated record of a risk of a reaction to the identified substance."
	case AllergyIntoleranceClinicalStatusResolved:
		return "A reaction to the identified substance has been clinically reassessed by testing or re-exposure and considered to be resolved."
	}
	return "<unknown>"
}
