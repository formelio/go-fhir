package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationRequestStatus is documented here http://hl7.org/fhir/ValueSet/medication-request-status
type MedicationRequestStatus int

const (
	MedicationRequestStatusActive MedicationRequestStatus = iota
	MedicationRequestStatusOnHold
	MedicationRequestStatusCancelled
	MedicationRequestStatusCompleted
	MedicationRequestStatusEnteredInError
	MedicationRequestStatusStopped
	MedicationRequestStatusDraft
	MedicationRequestStatusUnknown
)

func (code MedicationRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationRequestStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationRequestStatusActive
	case "on-hold":
		*code = MedicationRequestStatusOnHold
	case "cancelled":
		*code = MedicationRequestStatusCancelled
	case "completed":
		*code = MedicationRequestStatusCompleted
	case "entered-in-error":
		*code = MedicationRequestStatusEnteredInError
	case "stopped":
		*code = MedicationRequestStatusStopped
	case "draft":
		*code = MedicationRequestStatusDraft
	case "unknown":
		*code = MedicationRequestStatusUnknown
	default:
		return fmt.Errorf("unknown MedicationRequestStatus code `%s`", s)
	}
	return nil
}
func (code MedicationRequestStatus) String() string {
	return code.Code()
}
func (code MedicationRequestStatus) Code() string {
	switch code {
	case MedicationRequestStatusActive:
		return "active"
	case MedicationRequestStatusOnHold:
		return "on-hold"
	case MedicationRequestStatusCancelled:
		return "cancelled"
	case MedicationRequestStatusCompleted:
		return "completed"
	case MedicationRequestStatusEnteredInError:
		return "entered-in-error"
	case MedicationRequestStatusStopped:
		return "stopped"
	case MedicationRequestStatusDraft:
		return "draft"
	case MedicationRequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code MedicationRequestStatus) Display() string {
	switch code {
	case MedicationRequestStatusActive:
		return "Active"
	case MedicationRequestStatusOnHold:
		return "On Hold"
	case MedicationRequestStatusCancelled:
		return "Cancelled"
	case MedicationRequestStatusCompleted:
		return "Completed"
	case MedicationRequestStatusEnteredInError:
		return "Entered In Error"
	case MedicationRequestStatusStopped:
		return "Stopped"
	case MedicationRequestStatusDraft:
		return "Draft"
	case MedicationRequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code MedicationRequestStatus) Definition() string {
	switch code {
	case MedicationRequestStatusActive:
		return "The prescription is 'actionable', but not all actions that are implied by it have occurred yet."
	case MedicationRequestStatusOnHold:
		return "Actions implied by the prescription are to be temporarily halted, but are expected to continue later.  May also be called \"suspended\"."
	case MedicationRequestStatusCancelled:
		return "The prescription has been withdrawn."
	case MedicationRequestStatusCompleted:
		return "All actions that are implied by the prescription have occurred."
	case MedicationRequestStatusEnteredInError:
		return "The prescription was entered in error."
	case MedicationRequestStatusStopped:
		return "Actions implied by the prescription are to be permanently halted, before all of them occurred."
	case MedicationRequestStatusDraft:
		return "The prescription is not yet 'actionable', i.e. it is a work in progress, requires sign-off or verification, and needs to be run through decision support process."
	case MedicationRequestStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this request"
	}
	return "<unknown>"
}
