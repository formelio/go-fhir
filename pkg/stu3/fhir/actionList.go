package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ActionList is documented here http://hl7.org/fhir/ValueSet/actionlist
type ActionList int

const (
	ActionListCancel ActionList = iota
	ActionListPoll
	ActionListReprocess
	ActionListStatus
)

func (code ActionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionList) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "cancel":
		*code = ActionListCancel
	case "poll":
		*code = ActionListPoll
	case "reprocess":
		*code = ActionListReprocess
	case "status":
		*code = ActionListStatus
	default:
		return fmt.Errorf("unknown ActionList code `%s`", s)
	}
	return nil
}
func (code ActionList) String() string {
	return code.Code()
}
func (code ActionList) Code() string {
	switch code {
	case ActionListCancel:
		return "cancel"
	case ActionListPoll:
		return "poll"
	case ActionListReprocess:
		return "reprocess"
	case ActionListStatus:
		return "status"
	}
	return "<unknown>"
}
func (code ActionList) Display() string {
	switch code {
	case ActionListCancel:
		return "Cancel, Reverse or Nullify"
	case ActionListPoll:
		return "Poll"
	case ActionListReprocess:
		return "Re-Process"
	case ActionListStatus:
		return "Status Check"
	}
	return "<unknown>"
}
func (code ActionList) Definition() string {
	switch code {
	case ActionListCancel:
		return "Cancel, reverse or nullify the target resource."
	case ActionListPoll:
		return "Check for previously un-read/ not-retrieved resources."
	case ActionListReprocess:
		return "Re-process the target resource."
	case ActionListStatus:
		return "Retrieve the processing status of the target resource."
	}
	return "<unknown>"
}
