package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SlotStatus is documented here http://hl7.org/fhir/ValueSet/slotstatus
type SlotStatus int

const (
	SlotStatusBusy SlotStatus = iota
	SlotStatusFree
	SlotStatusBusyUnavailable
	SlotStatusBusyTentative
	SlotStatusEnteredInError
)

func (code SlotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SlotStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "busy":
		*code = SlotStatusBusy
	case "free":
		*code = SlotStatusFree
	case "busy-unavailable":
		*code = SlotStatusBusyUnavailable
	case "busy-tentative":
		*code = SlotStatusBusyTentative
	case "entered-in-error":
		*code = SlotStatusEnteredInError
	default:
		return fmt.Errorf("unknown SlotStatus code `%s`", s)
	}
	return nil
}
func (code SlotStatus) String() string {
	return code.Code()
}
func (code SlotStatus) Code() string {
	switch code {
	case SlotStatusBusy:
		return "busy"
	case SlotStatusFree:
		return "free"
	case SlotStatusBusyUnavailable:
		return "busy-unavailable"
	case SlotStatusBusyTentative:
		return "busy-tentative"
	case SlotStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code SlotStatus) Display() string {
	switch code {
	case SlotStatusBusy:
		return "Busy"
	case SlotStatusFree:
		return "Free"
	case SlotStatusBusyUnavailable:
		return "Busy (Unavailable)"
	case SlotStatusBusyTentative:
		return "Busy (Tentative)"
	case SlotStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}
func (code SlotStatus) Definition() string {
	switch code {
	case SlotStatusBusy:
		return "Indicates that the time interval is busy because one  or more events have been scheduled for that interval."
	case SlotStatusFree:
		return "Indicates that the time interval is free for scheduling."
	case SlotStatusBusyUnavailable:
		return "Indicates that the time interval is busy and that the interval can not be scheduled."
	case SlotStatusBusyTentative:
		return "Indicates that the time interval is busy because one or more events have been tentatively scheduled for that interval."
	case SlotStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
