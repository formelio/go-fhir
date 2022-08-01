package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationStatementTaken is documented here http://hl7.org/fhir/ValueSet/medication-statement-taken
type MedicationStatementTaken int

const (
	MedicationStatementTakenY MedicationStatementTaken = iota
	MedicationStatementTakenN
	MedicationStatementTakenUnk
	MedicationStatementTakenNa
)

func (code MedicationStatementTaken) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationStatementTaken) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "y":
		*code = MedicationStatementTakenY
	case "n":
		*code = MedicationStatementTakenN
	case "unk":
		*code = MedicationStatementTakenUnk
	case "na":
		*code = MedicationStatementTakenNa
	default:
		return fmt.Errorf("unknown MedicationStatementTaken code `%s`", s)
	}
	return nil
}
func (code MedicationStatementTaken) String() string {
	return code.Code()
}
func (code MedicationStatementTaken) Code() string {
	switch code {
	case MedicationStatementTakenY:
		return "y"
	case MedicationStatementTakenN:
		return "n"
	case MedicationStatementTakenUnk:
		return "unk"
	case MedicationStatementTakenNa:
		return "na"
	}
	return "<unknown>"
}
func (code MedicationStatementTaken) Display() string {
	switch code {
	case MedicationStatementTakenY:
		return "Yes"
	case MedicationStatementTakenN:
		return "No"
	case MedicationStatementTakenUnk:
		return "Unknown"
	case MedicationStatementTakenNa:
		return "Not Applicable"
	}
	return "<unknown>"
}
func (code MedicationStatementTaken) Definition() string {
	switch code {
	case MedicationStatementTakenY:
		return "Positive assertion that patient has taken medication"
	case MedicationStatementTakenN:
		return "Negative assertion that patient has not taken medication"
	case MedicationStatementTakenUnk:
		return "Unknown assertion if patient has taken medication"
	case MedicationStatementTakenNa:
		return "Patient reporting does not apply"
	}
	return "<unknown>"
}
