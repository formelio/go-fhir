package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AuditEventOutcome is documented here http://hl7.org/fhir/ValueSet/audit-event-outcome
type AuditEventOutcome int

const (
	AuditEventOutcome0 AuditEventOutcome = iota
	AuditEventOutcome4
	AuditEventOutcome8
	AuditEventOutcome12
)

func (code AuditEventOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AuditEventOutcome) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "0":
		*code = AuditEventOutcome0
	case "4":
		*code = AuditEventOutcome4
	case "8":
		*code = AuditEventOutcome8
	case "12":
		*code = AuditEventOutcome12
	default:
		return fmt.Errorf("unknown AuditEventOutcome code `%s`", s)
	}
	return nil
}
func (code AuditEventOutcome) String() string {
	return code.Code()
}
func (code AuditEventOutcome) Code() string {
	switch code {
	case AuditEventOutcome0:
		return "0"
	case AuditEventOutcome4:
		return "4"
	case AuditEventOutcome8:
		return "8"
	case AuditEventOutcome12:
		return "12"
	}
	return "<unknown>"
}
func (code AuditEventOutcome) Display() string {
	switch code {
	case AuditEventOutcome0:
		return "Success"
	case AuditEventOutcome4:
		return "Minor failure"
	case AuditEventOutcome8:
		return "Serious failure"
	case AuditEventOutcome12:
		return "Major failure"
	}
	return "<unknown>"
}
func (code AuditEventOutcome) Definition() string {
	switch code {
	case AuditEventOutcome0:
		return "The operation completed successfully (whether with warnings or not)."
	case AuditEventOutcome4:
		return "The action was not successful due to some kind of catered for error (often equivalent to an HTTP 400 response)."
	case AuditEventOutcome8:
		return "The action was not successful due to some kind of unexpected error (often equivalent to an HTTP 500 response)."
	case AuditEventOutcome12:
		return "An error of such magnitude occurred that the system is no longer available for use (i.e. the system died)."
	}
	return "<unknown>"
}
