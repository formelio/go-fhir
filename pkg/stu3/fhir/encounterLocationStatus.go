package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// EncounterLocationStatus is documented here http://hl7.org/fhir/ValueSet/encounter-location-status
type EncounterLocationStatus int

const (
	EncounterLocationStatusPlanned EncounterLocationStatus = iota
	EncounterLocationStatusActive
	EncounterLocationStatusReserved
	EncounterLocationStatusCompleted
)

func (code EncounterLocationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EncounterLocationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "planned":
		*code = EncounterLocationStatusPlanned
	case "active":
		*code = EncounterLocationStatusActive
	case "reserved":
		*code = EncounterLocationStatusReserved
	case "completed":
		*code = EncounterLocationStatusCompleted
	default:
		return fmt.Errorf("unknown EncounterLocationStatus code `%s`", s)
	}
	return nil
}
func (code EncounterLocationStatus) String() string {
	return code.Code()
}
func (code EncounterLocationStatus) Code() string {
	switch code {
	case EncounterLocationStatusPlanned:
		return "planned"
	case EncounterLocationStatusActive:
		return "active"
	case EncounterLocationStatusReserved:
		return "reserved"
	case EncounterLocationStatusCompleted:
		return "completed"
	}
	return "<unknown>"
}
func (code EncounterLocationStatus) Display() string {
	switch code {
	case EncounterLocationStatusPlanned:
		return "Planned"
	case EncounterLocationStatusActive:
		return "Active"
	case EncounterLocationStatusReserved:
		return "Reserved"
	case EncounterLocationStatusCompleted:
		return "Completed"
	}
	return "<unknown>"
}
func (code EncounterLocationStatus) Definition() string {
	switch code {
	case EncounterLocationStatusPlanned:
		return "The patient is planned to be moved to this location at some point in the future."
	case EncounterLocationStatusActive:
		return "The patient is currently at this location, or was between the period specified.\r\rA system may update these records when the patient leaves the location to either reserved, or completed"
	case EncounterLocationStatusReserved:
		return "This location is held empty for this patient."
	case EncounterLocationStatusCompleted:
		return "The patient was at this location during the period specified.\r\rNot to be used when the patient is currently at the location"
	}
	return "<unknown>"
}
