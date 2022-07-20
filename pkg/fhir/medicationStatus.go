package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationStatus is documented here http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatus int

const (
	MedicationStatusActive MedicationStatus = iota
	MedicationStatusInactive
	MedicationStatusEnteredInError
)

func (code MedicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationStatusActive
	case "inactive":
		*code = MedicationStatusInactive
	case "entered-in-error":
		*code = MedicationStatusEnteredInError
	default:
		return fmt.Errorf("unknown MedicationStatus code `%s`", s)
	}
	return nil
}
func (code MedicationStatus) String() string {
	return code.Code()
}
func (code MedicationStatus) Code() string {
	switch code {
	case MedicationStatusActive:
		return "active"
	case MedicationStatusInactive:
		return "inactive"
	case MedicationStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code MedicationStatus) Display() string {
	switch code {
	case MedicationStatusActive:
		return "Active"
	case MedicationStatusInactive:
		return "Inactive"
	case MedicationStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code MedicationStatus) Definition() string {
	switch code {
	case MedicationStatusActive:
		return "The medication is available for use"
	case MedicationStatusInactive:
		return "The medication is not available for use"
	case MedicationStatusEnteredInError:
		return "The medication was entered in error"
	}
	return "<unknown>"
}
