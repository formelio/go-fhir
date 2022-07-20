package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationAdministrationStatus is documented here http://hl7.org/fhir/ValueSet/medication-admin-status
type MedicationAdministrationStatus int

const (
	MedicationAdministrationStatusInProgress MedicationAdministrationStatus = iota
	MedicationAdministrationStatusOnHold
	MedicationAdministrationStatusCompleted
	MedicationAdministrationStatusEnteredInError
	MedicationAdministrationStatusStopped
	MedicationAdministrationStatusUnknown
)

func (code MedicationAdministrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationAdministrationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-progress":
		*code = MedicationAdministrationStatusInProgress
	case "on-hold":
		*code = MedicationAdministrationStatusOnHold
	case "completed":
		*code = MedicationAdministrationStatusCompleted
	case "entered-in-error":
		*code = MedicationAdministrationStatusEnteredInError
	case "stopped":
		*code = MedicationAdministrationStatusStopped
	case "unknown":
		*code = MedicationAdministrationStatusUnknown
	default:
		return fmt.Errorf("unknown MedicationAdministrationStatus code `%s`", s)
	}
	return nil
}
func (code MedicationAdministrationStatus) String() string {
	return code.Code()
}
func (code MedicationAdministrationStatus) Code() string {
	switch code {
	case MedicationAdministrationStatusInProgress:
		return "in-progress"
	case MedicationAdministrationStatusOnHold:
		return "on-hold"
	case MedicationAdministrationStatusCompleted:
		return "completed"
	case MedicationAdministrationStatusEnteredInError:
		return "entered-in-error"
	case MedicationAdministrationStatusStopped:
		return "stopped"
	case MedicationAdministrationStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code MedicationAdministrationStatus) Display() string {
	switch code {
	case MedicationAdministrationStatusInProgress:
		return "In Progress"
	case MedicationAdministrationStatusOnHold:
		return "On Hold"
	case MedicationAdministrationStatusCompleted:
		return "Completed"
	case MedicationAdministrationStatusEnteredInError:
		return "Entered in Error"
	case MedicationAdministrationStatusStopped:
		return "Stopped"
	case MedicationAdministrationStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code MedicationAdministrationStatus) Definition() string {
	switch code {
	case MedicationAdministrationStatusInProgress:
		return "The administration has started but has not yet completed."
	case MedicationAdministrationStatusOnHold:
		return "Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	case MedicationAdministrationStatusCompleted:
		return "All actions that are implied by the administration have occurred."
	case MedicationAdministrationStatusEnteredInError:
		return "The administration was entered in error and therefore nullified."
	case MedicationAdministrationStatusStopped:
		return "Actions implied by the administration have been permanently halted, before all of them occurred."
	case MedicationAdministrationStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
