package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GoalStatus is documented here http://hl7.org/fhir/ValueSet/goal-status
type GoalStatus int

const (
	GoalStatusProposed GoalStatus = iota
	GoalStatusAccepted
	GoalStatusPlanned
	GoalStatusInProgress
	GoalStatusOnTarget
	GoalStatusAheadOfTarget
	GoalStatusBehindTarget
	GoalStatusSustaining
	GoalStatusAchieved
	GoalStatusOnHold
	GoalStatusCancelled
	GoalStatusEnteredInError
	GoalStatusRejected
)

func (code GoalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GoalStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = GoalStatusProposed
	case "accepted":
		*code = GoalStatusAccepted
	case "planned":
		*code = GoalStatusPlanned
	case "in-progress":
		*code = GoalStatusInProgress
	case "on-target":
		*code = GoalStatusOnTarget
	case "ahead-of-target":
		*code = GoalStatusAheadOfTarget
	case "behind-target":
		*code = GoalStatusBehindTarget
	case "sustaining":
		*code = GoalStatusSustaining
	case "achieved":
		*code = GoalStatusAchieved
	case "on-hold":
		*code = GoalStatusOnHold
	case "cancelled":
		*code = GoalStatusCancelled
	case "entered-in-error":
		*code = GoalStatusEnteredInError
	case "rejected":
		*code = GoalStatusRejected
	default:
		return fmt.Errorf("unknown GoalStatus code `%s`", s)
	}
	return nil
}
func (code GoalStatus) String() string {
	return code.Code()
}
func (code GoalStatus) Code() string {
	switch code {
	case GoalStatusProposed:
		return "proposed"
	case GoalStatusAccepted:
		return "accepted"
	case GoalStatusPlanned:
		return "planned"
	case GoalStatusInProgress:
		return "in-progress"
	case GoalStatusOnTarget:
		return "on-target"
	case GoalStatusAheadOfTarget:
		return "ahead-of-target"
	case GoalStatusBehindTarget:
		return "behind-target"
	case GoalStatusSustaining:
		return "sustaining"
	case GoalStatusAchieved:
		return "achieved"
	case GoalStatusOnHold:
		return "on-hold"
	case GoalStatusCancelled:
		return "cancelled"
	case GoalStatusEnteredInError:
		return "entered-in-error"
	case GoalStatusRejected:
		return "rejected"
	}
	return "<unknown>"
}
func (code GoalStatus) Display() string {
	switch code {
	case GoalStatusProposed:
		return "Proposed"
	case GoalStatusAccepted:
		return "Accepted"
	case GoalStatusPlanned:
		return "Planned"
	case GoalStatusInProgress:
		return "In Progress"
	case GoalStatusOnTarget:
		return "On Target"
	case GoalStatusAheadOfTarget:
		return "Ahead of Target"
	case GoalStatusBehindTarget:
		return "Behind Target"
	case GoalStatusSustaining:
		return "Sustaining"
	case GoalStatusAchieved:
		return "Achieved"
	case GoalStatusOnHold:
		return "On Hold"
	case GoalStatusCancelled:
		return "Cancelled"
	case GoalStatusEnteredInError:
		return "Entered In Error"
	case GoalStatusRejected:
		return "Rejected"
	}
	return "<unknown>"
}
func (code GoalStatus) Definition() string {
	switch code {
	case GoalStatusProposed:
		return "A goal is proposed for this patient"
	case GoalStatusAccepted:
		return "A proposed goal was accepted or acknowledged"
	case GoalStatusPlanned:
		return "A goal is planned for this patient"
	case GoalStatusInProgress:
		return "The goal is being sought but has not yet been reached.  (Also applies if goal was reached in the past but there has been regression and goal is being sought again)"
	case GoalStatusOnTarget:
		return "The goal is on schedule for the planned timelines"
	case GoalStatusAheadOfTarget:
		return "The goal is ahead of the planned timelines"
	case GoalStatusBehindTarget:
		return "The goal is behind the planned timelines"
	case GoalStatusSustaining:
		return "The goal has been met, but ongoing activity is needed to sustain the goal objective"
	case GoalStatusAchieved:
		return "The goal has been met and no further action is needed"
	case GoalStatusOnHold:
		return "The goal remains a long term objective but is no longer being actively pursued for a temporary period of time."
	case GoalStatusCancelled:
		return "The previously accepted goal is no longer being sought"
	case GoalStatusEnteredInError:
		return "The goal was entered in error and voided."
	case GoalStatusRejected:
		return "A proposed goal was rejected"
	}
	return "<unknown>"
}
