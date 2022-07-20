package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationDispenseStatus is documented here http://hl7.org/fhir/ValueSet/medication-dispense-status
type MedicationDispenseStatus int

const (
	MedicationDispenseStatusPreparation MedicationDispenseStatus = iota
	MedicationDispenseStatusInProgress
	MedicationDispenseStatusOnHold
	MedicationDispenseStatusCompleted
	MedicationDispenseStatusEnteredInError
	MedicationDispenseStatusStopped
)

func (code MedicationDispenseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationDispenseStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = MedicationDispenseStatusPreparation
	case "in-progress":
		*code = MedicationDispenseStatusInProgress
	case "on-hold":
		*code = MedicationDispenseStatusOnHold
	case "completed":
		*code = MedicationDispenseStatusCompleted
	case "entered-in-error":
		*code = MedicationDispenseStatusEnteredInError
	case "stopped":
		*code = MedicationDispenseStatusStopped
	default:
		return fmt.Errorf("unknown MedicationDispenseStatus code `%s`", s)
	}
	return nil
}
func (code MedicationDispenseStatus) String() string {
	return code.Code()
}
func (code MedicationDispenseStatus) Code() string {
	switch code {
	case MedicationDispenseStatusPreparation:
		return "preparation"
	case MedicationDispenseStatusInProgress:
		return "in-progress"
	case MedicationDispenseStatusOnHold:
		return "on-hold"
	case MedicationDispenseStatusCompleted:
		return "completed"
	case MedicationDispenseStatusEnteredInError:
		return "entered-in-error"
	case MedicationDispenseStatusStopped:
		return "stopped"
	}
	return "<unknown>"
}
func (code MedicationDispenseStatus) Display() string {
	switch code {
	case MedicationDispenseStatusPreparation:
		return "Preparation"
	case MedicationDispenseStatusInProgress:
		return "In Progress"
	case MedicationDispenseStatusOnHold:
		return "On Hold"
	case MedicationDispenseStatusCompleted:
		return "Completed"
	case MedicationDispenseStatusEnteredInError:
		return "Entered in-Error"
	case MedicationDispenseStatusStopped:
		return "Stopped"
	}
	return "<unknown>"
}
func (code MedicationDispenseStatus) Definition() string {
	switch code {
	case MedicationDispenseStatusPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. initial compounding or packaging of medication). Preparation stages may be tracked for billing purposes."
	case MedicationDispenseStatusInProgress:
		return "The dispense has started but has not yet completed."
	case MedicationDispenseStatusOnHold:
		return "Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called \"suspended\""
	case MedicationDispenseStatusCompleted:
		return "All actions that are implied by the dispense have occurred."
	case MedicationDispenseStatusEnteredInError:
		return "The dispense was entered in error and therefore nullified."
	case MedicationDispenseStatusStopped:
		return "Actions implied by the dispense have been permanently halted, before all of them occurred."
	}
	return "<unknown>"
}
