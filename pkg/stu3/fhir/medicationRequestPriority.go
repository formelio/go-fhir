package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationRequestPriority is documented here http://hl7.org/fhir/ValueSet/medication-request-priority
type MedicationRequestPriority int

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = iota
	MedicationRequestPriorityUrgent
	MedicationRequestPriorityStat
	MedicationRequestPriorityAsap
)

func (code MedicationRequestPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationRequestPriority) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "routine":
		*code = MedicationRequestPriorityRoutine
	case "urgent":
		*code = MedicationRequestPriorityUrgent
	case "stat":
		*code = MedicationRequestPriorityStat
	case "asap":
		*code = MedicationRequestPriorityAsap
	default:
		return fmt.Errorf("unknown MedicationRequestPriority code `%s`", s)
	}
	return nil
}
func (code MedicationRequestPriority) String() string {
	return code.Code()
}
func (code MedicationRequestPriority) Code() string {
	switch code {
	case MedicationRequestPriorityRoutine:
		return "routine"
	case MedicationRequestPriorityUrgent:
		return "urgent"
	case MedicationRequestPriorityStat:
		return "stat"
	case MedicationRequestPriorityAsap:
		return "asap"
	}
	return "<unknown>"
}
func (code MedicationRequestPriority) Display() string {
	switch code {
	case MedicationRequestPriorityRoutine:
		return "Routine"
	case MedicationRequestPriorityUrgent:
		return "Urgent"
	case MedicationRequestPriorityStat:
		return "Stat"
	case MedicationRequestPriorityAsap:
		return "ASAP"
	}
	return "<unknown>"
}
func (code MedicationRequestPriority) Definition() string {
	switch code {
	case MedicationRequestPriorityRoutine:
		return "The order has a normal priority ."
	case MedicationRequestPriorityUrgent:
		return "The order should be urgently."
	case MedicationRequestPriorityStat:
		return "The order is time-critical."
	case MedicationRequestPriorityAsap:
		return "The order should be acted on as soon as possible."
	}
	return "<unknown>"
}
