package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TriggerType is documented here http://hl7.org/fhir/ValueSet/trigger-type
type TriggerType int

const (
	TriggerTypeNamedEvent TriggerType = iota
	TriggerTypePeriodic
	TriggerTypeDataAdded
	TriggerTypeDataModified
	TriggerTypeDataRemoved
	TriggerTypeDataAccessed
	TriggerTypeDataAccessEnded
)

func (code TriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TriggerType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "named-event":
		*code = TriggerTypeNamedEvent
	case "periodic":
		*code = TriggerTypePeriodic
	case "data-added":
		*code = TriggerTypeDataAdded
	case "data-modified":
		*code = TriggerTypeDataModified
	case "data-removed":
		*code = TriggerTypeDataRemoved
	case "data-accessed":
		*code = TriggerTypeDataAccessed
	case "data-access-ended":
		*code = TriggerTypeDataAccessEnded
	default:
		return fmt.Errorf("unknown TriggerType code `%s`", s)
	}
	return nil
}
func (code TriggerType) String() string {
	return code.Code()
}
func (code TriggerType) Code() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "named-event"
	case TriggerTypePeriodic:
		return "periodic"
	case TriggerTypeDataAdded:
		return "data-added"
	case TriggerTypeDataModified:
		return "data-modified"
	case TriggerTypeDataRemoved:
		return "data-removed"
	case TriggerTypeDataAccessed:
		return "data-accessed"
	case TriggerTypeDataAccessEnded:
		return "data-access-ended"
	}
	return "<unknown>"
}
func (code TriggerType) Display() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "Named Event"
	case TriggerTypePeriodic:
		return "Periodic"
	case TriggerTypeDataAdded:
		return "Data Added"
	case TriggerTypeDataModified:
		return "Data Modified"
	case TriggerTypeDataRemoved:
		return "Data Removed"
	case TriggerTypeDataAccessed:
		return "Data Accessed"
	case TriggerTypeDataAccessEnded:
		return "Data Access Ended"
	}
	return "<unknown>"
}
func (code TriggerType) Definition() string {
	switch code {
	case TriggerTypeNamedEvent:
		return "The trigger occurs in response to a specific named event"
	case TriggerTypePeriodic:
		return "The trigger occurs at a specific time or periodically as described by a timing or schedule"
	case TriggerTypeDataAdded:
		return "The trigger occurs whenever data of a particular type is added"
	case TriggerTypeDataModified:
		return "The trigger occurs whenever data of a particular type is modified"
	case TriggerTypeDataRemoved:
		return "The trigger occurs whenever data of a particular type is removed"
	case TriggerTypeDataAccessed:
		return "The trigger occurs whenever data of a particular type is accessed"
	case TriggerTypeDataAccessEnded:
		return "The trigger occurs whenever access to data of a particular type is completed"
	}
	return "<unknown>"
}
