package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// EventStatus is documented here http://hl7.org/fhir/ValueSet/event-status
type EventStatus int

const (
	EventStatusPreparation EventStatus = iota
	EventStatusInProgress
	EventStatusSuspended
	EventStatusAborted
	EventStatusCompleted
	EventStatusEnteredInError
	EventStatusUnknown
)

func (code EventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EventStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preparation":
		*code = EventStatusPreparation
	case "in-progress":
		*code = EventStatusInProgress
	case "suspended":
		*code = EventStatusSuspended
	case "aborted":
		*code = EventStatusAborted
	case "completed":
		*code = EventStatusCompleted
	case "entered-in-error":
		*code = EventStatusEnteredInError
	case "unknown":
		*code = EventStatusUnknown
	default:
		return fmt.Errorf("unknown EventStatus code `%s`", s)
	}
	return nil
}
func (code EventStatus) String() string {
	return code.Code()
}
func (code EventStatus) Code() string {
	switch code {
	case EventStatusPreparation:
		return "preparation"
	case EventStatusInProgress:
		return "in-progress"
	case EventStatusSuspended:
		return "suspended"
	case EventStatusAborted:
		return "aborted"
	case EventStatusCompleted:
		return "completed"
	case EventStatusEnteredInError:
		return "entered-in-error"
	case EventStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code EventStatus) Display() string {
	switch code {
	case EventStatusPreparation:
		return "Preparation"
	case EventStatusInProgress:
		return "In Progress"
	case EventStatusSuspended:
		return "Suspended"
	case EventStatusAborted:
		return "Aborted"
	case EventStatusCompleted:
		return "Completed"
	case EventStatusEnteredInError:
		return "Entered in Error"
	case EventStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code EventStatus) Definition() string {
	switch code {
	case EventStatusPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes."
	case EventStatusInProgress:
		return "The event is currently occurring"
	case EventStatusSuspended:
		return "The event has been temporarily stopped but is expected to resume in the future"
	case EventStatusAborted:
		return "The event was  prior to the full completion of the intended actions"
	case EventStatusCompleted:
		return "The event has now concluded"
	case EventStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)"
	case EventStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
