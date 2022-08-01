package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationStatementStatus is documented here http://hl7.org/fhir/ValueSet/medication-statement-status
type MedicationStatementStatus int

const (
	MedicationStatementStatusActive MedicationStatementStatus = iota
	MedicationStatementStatusCompleted
	MedicationStatementStatusEnteredInError
	MedicationStatementStatusIntended
	MedicationStatementStatusStopped
	MedicationStatementStatusOnHold
)

func (code MedicationStatementStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationStatementStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationStatementStatusActive
	case "completed":
		*code = MedicationStatementStatusCompleted
	case "entered-in-error":
		*code = MedicationStatementStatusEnteredInError
	case "intended":
		*code = MedicationStatementStatusIntended
	case "stopped":
		*code = MedicationStatementStatusStopped
	case "on-hold":
		*code = MedicationStatementStatusOnHold
	default:
		return fmt.Errorf("unknown MedicationStatementStatus code `%s`", s)
	}
	return nil
}
func (code MedicationStatementStatus) String() string {
	return code.Code()
}
func (code MedicationStatementStatus) Code() string {
	switch code {
	case MedicationStatementStatusActive:
		return "active"
	case MedicationStatementStatusCompleted:
		return "completed"
	case MedicationStatementStatusEnteredInError:
		return "entered-in-error"
	case MedicationStatementStatusIntended:
		return "intended"
	case MedicationStatementStatusStopped:
		return "stopped"
	case MedicationStatementStatusOnHold:
		return "on-hold"
	}
	return "<unknown>"
}
func (code MedicationStatementStatus) Display() string {
	switch code {
	case MedicationStatementStatusActive:
		return "Active"
	case MedicationStatementStatusCompleted:
		return "Completed"
	case MedicationStatementStatusEnteredInError:
		return "Entered in Error"
	case MedicationStatementStatusIntended:
		return "Intended"
	case MedicationStatementStatusStopped:
		return "Stopped"
	case MedicationStatementStatusOnHold:
		return "On Hold"
	}
	return "<unknown>"
}
func (code MedicationStatementStatus) Definition() string {
	switch code {
	case MedicationStatementStatusActive:
		return "The medication is still being taken."
	case MedicationStatementStatusCompleted:
		return "The medication is no longer being taken."
	case MedicationStatementStatusEnteredInError:
		return "The statement was recorded incorrectly."
	case MedicationStatementStatusIntended:
		return "The medication may be taken at some time in the future."
	case MedicationStatementStatusStopped:
		return "Actions implied by the statement have been permanently halted, before all of them occurred."
	case MedicationStatementStatusOnHold:
		return "Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	}
	return "<unknown>"
}
