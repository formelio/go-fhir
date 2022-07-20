package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ResearchSubjectStatus is documented here http://hl7.org/fhir/ValueSet/research-subject-status
type ResearchSubjectStatus int

const (
	ResearchSubjectStatusCandidate ResearchSubjectStatus = iota
	ResearchSubjectStatusEnrolled
	ResearchSubjectStatusActive
	ResearchSubjectStatusSuspended
	ResearchSubjectStatusWithdrawn
	ResearchSubjectStatusCompleted
)

func (code ResearchSubjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchSubjectStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "candidate":
		*code = ResearchSubjectStatusCandidate
	case "enrolled":
		*code = ResearchSubjectStatusEnrolled
	case "active":
		*code = ResearchSubjectStatusActive
	case "suspended":
		*code = ResearchSubjectStatusSuspended
	case "withdrawn":
		*code = ResearchSubjectStatusWithdrawn
	case "completed":
		*code = ResearchSubjectStatusCompleted
	default:
		return fmt.Errorf("unknown ResearchSubjectStatus code `%s`", s)
	}
	return nil
}
func (code ResearchSubjectStatus) String() string {
	return code.Code()
}
func (code ResearchSubjectStatus) Code() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "candidate"
	case ResearchSubjectStatusEnrolled:
		return "enrolled"
	case ResearchSubjectStatusActive:
		return "active"
	case ResearchSubjectStatusSuspended:
		return "suspended"
	case ResearchSubjectStatusWithdrawn:
		return "withdrawn"
	case ResearchSubjectStatusCompleted:
		return "completed"
	}
	return "<unknown>"
}
func (code ResearchSubjectStatus) Display() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "Candidate"
	case ResearchSubjectStatusEnrolled:
		return "Enrolled"
	case ResearchSubjectStatusActive:
		return "Active"
	case ResearchSubjectStatusSuspended:
		return "Suspended"
	case ResearchSubjectStatusWithdrawn:
		return "Withdrawn"
	case ResearchSubjectStatusCompleted:
		return "Completed"
	}
	return "<unknown>"
}
func (code ResearchSubjectStatus) Definition() string {
	switch code {
	case ResearchSubjectStatusCandidate:
		return "The subject has been identified as a potential participant in the study but has not yet agreed to participate"
	case ResearchSubjectStatusEnrolled:
		return "The subject has agreed to participate in the study but has not yet begun performing any action within the study"
	case ResearchSubjectStatusActive:
		return "The subject is currently being monitored and/or subject to treatment as part of the study"
	case ResearchSubjectStatusSuspended:
		return "The subject has temporarily discontinued monitoring/treatment as part of the study"
	case ResearchSubjectStatusWithdrawn:
		return "The subject has permanently ended participation in the study prior to completion of the intended monitoring/treatment"
	case ResearchSubjectStatusCompleted:
		return "All intended monitoring/treatment of the subject has been completed and their engagement with the study is now ended"
	}
	return "<unknown>"
}
