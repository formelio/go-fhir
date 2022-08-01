package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// RequestStatus is documented here http://hl7.org/fhir/ValueSet/request-status
type RequestStatus int

const (
	RequestStatusDraft RequestStatus = iota
	RequestStatusActive
	RequestStatusSuspended
	RequestStatusCancelled
	RequestStatusCompleted
	RequestStatusEnteredInError
	RequestStatusUnknown
)

func (code RequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = RequestStatusDraft
	case "active":
		*code = RequestStatusActive
	case "suspended":
		*code = RequestStatusSuspended
	case "cancelled":
		*code = RequestStatusCancelled
	case "completed":
		*code = RequestStatusCompleted
	case "entered-in-error":
		*code = RequestStatusEnteredInError
	case "unknown":
		*code = RequestStatusUnknown
	default:
		return fmt.Errorf("unknown RequestStatus code `%s`", s)
	}
	return nil
}
func (code RequestStatus) String() string {
	return code.Code()
}
func (code RequestStatus) Code() string {
	switch code {
	case RequestStatusDraft:
		return "draft"
	case RequestStatusActive:
		return "active"
	case RequestStatusSuspended:
		return "suspended"
	case RequestStatusCancelled:
		return "cancelled"
	case RequestStatusCompleted:
		return "completed"
	case RequestStatusEnteredInError:
		return "entered-in-error"
	case RequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code RequestStatus) Display() string {
	switch code {
	case RequestStatusDraft:
		return "Draft"
	case RequestStatusActive:
		return "Active"
	case RequestStatusSuspended:
		return "Suspended"
	case RequestStatusCancelled:
		return "Cancelled"
	case RequestStatusCompleted:
		return "Completed"
	case RequestStatusEnteredInError:
		return "Entered in Error"
	case RequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code RequestStatus) Definition() string {
	switch code {
	case RequestStatusDraft:
		return "The request has been created but is not yet complete or ready for action"
	case RequestStatusActive:
		return "The request is ready to be acted upon"
	case RequestStatusSuspended:
		return "The authorization/request to act has been temporarily withdrawn but is expected to resume in the future"
	case RequestStatusCancelled:
		return "The authorization/request to act has been terminated prior to the full completion of the intended actions.  No further activity should occur."
	case RequestStatusCompleted:
		return "Activity against the request has been sufficiently completed to the satisfaction of the requester"
	case RequestStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)"
	case RequestStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for \"other\" . One of the listed statuses is presumed to apply,  but the system creating the request doesn't know."
	}
	return "<unknown>"
}
