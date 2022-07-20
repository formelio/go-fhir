package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ResearchStudyStatus is documented here http://hl7.org/fhir/ValueSet/research-study-status
type ResearchStudyStatus int

const (
	ResearchStudyStatusDraft ResearchStudyStatus = iota
	ResearchStudyStatusInProgress
	ResearchStudyStatusSuspended
	ResearchStudyStatusStopped
	ResearchStudyStatusCompleted
	ResearchStudyStatusEnteredInError
)

func (code ResearchStudyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchStudyStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = ResearchStudyStatusDraft
	case "in-progress":
		*code = ResearchStudyStatusInProgress
	case "suspended":
		*code = ResearchStudyStatusSuspended
	case "stopped":
		*code = ResearchStudyStatusStopped
	case "completed":
		*code = ResearchStudyStatusCompleted
	case "entered-in-error":
		*code = ResearchStudyStatusEnteredInError
	default:
		return fmt.Errorf("unknown ResearchStudyStatus code `%s`", s)
	}
	return nil
}
func (code ResearchStudyStatus) String() string {
	return code.Code()
}
func (code ResearchStudyStatus) Code() string {
	switch code {
	case ResearchStudyStatusDraft:
		return "draft"
	case ResearchStudyStatusInProgress:
		return "in-progress"
	case ResearchStudyStatusSuspended:
		return "suspended"
	case ResearchStudyStatusStopped:
		return "stopped"
	case ResearchStudyStatusCompleted:
		return "completed"
	case ResearchStudyStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ResearchStudyStatus) Display() string {
	switch code {
	case ResearchStudyStatusDraft:
		return "Draft"
	case ResearchStudyStatusInProgress:
		return "In-progress"
	case ResearchStudyStatusSuspended:
		return "Suspended"
	case ResearchStudyStatusStopped:
		return "Stopped"
	case ResearchStudyStatusCompleted:
		return "Completed"
	case ResearchStudyStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}
func (code ResearchStudyStatus) Definition() string {
	switch code {
	case ResearchStudyStatusDraft:
		return "The study is undergoing design but the process of selecting study subjects and capturing data has not yet begun."
	case ResearchStudyStatusInProgress:
		return "The study is currently being executed"
	case ResearchStudyStatusSuspended:
		return "Execution of the study has been temporarily paused"
	case ResearchStudyStatusStopped:
		return "The study was terminated prior to the final determination of results"
	case ResearchStudyStatusCompleted:
		return "The information sought by the study has been gathered and compiled and no further work is being performed"
	case ResearchStudyStatusEnteredInError:
		return "This study never actually existed.  The record is retained for tracking purposes in the event decisions may have been made based on this erroneous information."
	}
	return "<unknown>"
}
