package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DaysOfWeek is documented here http://hl7.org/fhir/ValueSet/days-of-week
type DaysOfWeek int

const (
	DaysOfWeekMon DaysOfWeek = iota
	DaysOfWeekTue
	DaysOfWeekWed
	DaysOfWeekThu
	DaysOfWeekFri
	DaysOfWeekSat
	DaysOfWeekSun
)

func (code DaysOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DaysOfWeek) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "mon":
		*code = DaysOfWeekMon
	case "tue":
		*code = DaysOfWeekTue
	case "wed":
		*code = DaysOfWeekWed
	case "thu":
		*code = DaysOfWeekThu
	case "fri":
		*code = DaysOfWeekFri
	case "sat":
		*code = DaysOfWeekSat
	case "sun":
		*code = DaysOfWeekSun
	default:
		return fmt.Errorf("unknown DaysOfWeek code `%s`", s)
	}
	return nil
}
func (code DaysOfWeek) String() string {
	return code.Code()
}
func (code DaysOfWeek) Code() string {
	switch code {
	case DaysOfWeekMon:
		return "mon"
	case DaysOfWeekTue:
		return "tue"
	case DaysOfWeekWed:
		return "wed"
	case DaysOfWeekThu:
		return "thu"
	case DaysOfWeekFri:
		return "fri"
	case DaysOfWeekSat:
		return "sat"
	case DaysOfWeekSun:
		return "sun"
	}
	return "<unknown>"
}
func (code DaysOfWeek) Display() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday"
	case DaysOfWeekTue:
		return "Tuesday"
	case DaysOfWeekWed:
		return "Wednesday"
	case DaysOfWeekThu:
		return "Thursday"
	case DaysOfWeekFri:
		return "Friday"
	case DaysOfWeekSat:
		return "Saturday"
	case DaysOfWeekSun:
		return "Sunday"
	}
	return "<unknown>"
}
func (code DaysOfWeek) Definition() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday"
	case DaysOfWeekTue:
		return "Tuesday"
	case DaysOfWeekWed:
		return "Wednesday"
	case DaysOfWeekThu:
		return "Thursday"
	case DaysOfWeekFri:
		return "Friday"
	case DaysOfWeekSat:
		return "Saturday"
	case DaysOfWeekSun:
		return "Sunday"
	}
	return "<unknown>"
}
