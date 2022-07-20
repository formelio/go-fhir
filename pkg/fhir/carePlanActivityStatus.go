package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CarePlanActivityStatus is documented here http://hl7.org/fhir/ValueSet/care-plan-activity-status
type CarePlanActivityStatus int

const (
	CarePlanActivityStatusNotStarted CarePlanActivityStatus = iota
	CarePlanActivityStatusScheduled
	CarePlanActivityStatusInProgress
	CarePlanActivityStatusOnHold
	CarePlanActivityStatusCompleted
	CarePlanActivityStatusCancelled
	CarePlanActivityStatusUnknown
)

func (code CarePlanActivityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CarePlanActivityStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-started":
		*code = CarePlanActivityStatusNotStarted
	case "scheduled":
		*code = CarePlanActivityStatusScheduled
	case "in-progress":
		*code = CarePlanActivityStatusInProgress
	case "on-hold":
		*code = CarePlanActivityStatusOnHold
	case "completed":
		*code = CarePlanActivityStatusCompleted
	case "cancelled":
		*code = CarePlanActivityStatusCancelled
	case "unknown":
		*code = CarePlanActivityStatusUnknown
	default:
		return fmt.Errorf("unknown CarePlanActivityStatus code `%s`", s)
	}
	return nil
}
func (code CarePlanActivityStatus) String() string {
	return code.Code()
}
func (code CarePlanActivityStatus) Code() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "not-started"
	case CarePlanActivityStatusScheduled:
		return "scheduled"
	case CarePlanActivityStatusInProgress:
		return "in-progress"
	case CarePlanActivityStatusOnHold:
		return "on-hold"
	case CarePlanActivityStatusCompleted:
		return "completed"
	case CarePlanActivityStatusCancelled:
		return "cancelled"
	case CarePlanActivityStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code CarePlanActivityStatus) Display() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "Not Started"
	case CarePlanActivityStatusScheduled:
		return "Scheduled"
	case CarePlanActivityStatusInProgress:
		return "In Progress"
	case CarePlanActivityStatusOnHold:
		return "On Hold"
	case CarePlanActivityStatusCompleted:
		return "Completed"
	case CarePlanActivityStatusCancelled:
		return "Cancelled"
	case CarePlanActivityStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code CarePlanActivityStatus) Definition() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "Activity is planned but no action has yet been taken."
	case CarePlanActivityStatusScheduled:
		return "Appointment or other booking has occurred but activity has not yet begun."
	case CarePlanActivityStatusInProgress:
		return "Activity has been started but is not yet complete."
	case CarePlanActivityStatusOnHold:
		return "Activity was started but has temporarily ceased with an expectation of resumption at a future time."
	case CarePlanActivityStatusCompleted:
		return "The activities have been completed (more or less) as planned."
	case CarePlanActivityStatusCancelled:
		return "The activities have been ended prior to completion (perhaps even before they were started)."
	case CarePlanActivityStatusUnknown:
		return "The authoring system doesn't know the current state of the activity."
	}
	return "<unknown>"
}
