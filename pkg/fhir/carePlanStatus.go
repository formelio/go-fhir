package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CarePlanStatus is documented here http://hl7.org/fhir/ValueSet/care-plan-status
type CarePlanStatus int

const (
	CarePlanStatusDraft CarePlanStatus = iota
	CarePlanStatusActive
	CarePlanStatusSuspended
	CarePlanStatusCompleted
	CarePlanStatusEnteredInError
	CarePlanStatusCancelled
	CarePlanStatusUnknown
)

func (code CarePlanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CarePlanStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = CarePlanStatusDraft
	case "active":
		*code = CarePlanStatusActive
	case "suspended":
		*code = CarePlanStatusSuspended
	case "completed":
		*code = CarePlanStatusCompleted
	case "entered-in-error":
		*code = CarePlanStatusEnteredInError
	case "cancelled":
		*code = CarePlanStatusCancelled
	case "unknown":
		*code = CarePlanStatusUnknown
	default:
		return fmt.Errorf("unknown CarePlanStatus code `%s`", s)
	}
	return nil
}
func (code CarePlanStatus) String() string {
	return code.Code()
}
func (code CarePlanStatus) Code() string {
	switch code {
	case CarePlanStatusDraft:
		return "draft"
	case CarePlanStatusActive:
		return "active"
	case CarePlanStatusSuspended:
		return "suspended"
	case CarePlanStatusCompleted:
		return "completed"
	case CarePlanStatusEnteredInError:
		return "entered-in-error"
	case CarePlanStatusCancelled:
		return "cancelled"
	case CarePlanStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code CarePlanStatus) Display() string {
	switch code {
	case CarePlanStatusDraft:
		return "Pending"
	case CarePlanStatusActive:
		return "Active"
	case CarePlanStatusSuspended:
		return "Suspended"
	case CarePlanStatusCompleted:
		return "Completed"
	case CarePlanStatusEnteredInError:
		return "Entered In Error"
	case CarePlanStatusCancelled:
		return "Cancelled"
	case CarePlanStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code CarePlanStatus) Definition() string {
	switch code {
	case CarePlanStatusDraft:
		return "The plan is in development or awaiting use but is not yet intended to be acted upon."
	case CarePlanStatusActive:
		return "The plan is intended to be followed and used as part of patient care."
	case CarePlanStatusSuspended:
		return "The plan has been temporarily stopped but is expected to resume in the future."
	case CarePlanStatusCompleted:
		return "The plan is no longer in use and is not expected to be followed or used in patient care."
	case CarePlanStatusEnteredInError:
		return "The plan was entered in error and voided."
	case CarePlanStatusCancelled:
		return "The plan has been terminated prior to reaching completion (though it may have been replaced by a new plan)."
	case CarePlanStatusUnknown:
		return "The authoring system doesn't know the current state of the care plan."
	}
	return "<unknown>"
}
