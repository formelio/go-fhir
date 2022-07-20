package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ActionParticipantType is documented here http://hl7.org/fhir/ValueSet/action-participant-type
type ActionParticipantType int

const (
	ActionParticipantTypePatient ActionParticipantType = iota
	ActionParticipantTypePractitioner
	ActionParticipantTypeRelatedPerson
)

func (code ActionParticipantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionParticipantType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "patient":
		*code = ActionParticipantTypePatient
	case "practitioner":
		*code = ActionParticipantTypePractitioner
	case "related-person":
		*code = ActionParticipantTypeRelatedPerson
	default:
		return fmt.Errorf("unknown ActionParticipantType code `%s`", s)
	}
	return nil
}
func (code ActionParticipantType) String() string {
	return code.Code()
}
func (code ActionParticipantType) Code() string {
	switch code {
	case ActionParticipantTypePatient:
		return "patient"
	case ActionParticipantTypePractitioner:
		return "practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "related-person"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Display() string {
	switch code {
	case ActionParticipantTypePatient:
		return "Patient"
	case ActionParticipantTypePractitioner:
		return "Practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "Related Person"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Definition() string {
	switch code {
	case ActionParticipantTypePatient:
		return "The participant is the patient under evaluation"
	case ActionParticipantTypePractitioner:
		return "The participant is a practitioner involved in the patient's care"
	case ActionParticipantTypeRelatedPerson:
		return "The participant is a person related to the patient"
	}
	return "<unknown>"
}
