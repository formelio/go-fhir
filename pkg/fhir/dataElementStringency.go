package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DataElementStringency is documented here http://hl7.org/fhir/ValueSet/dataelement-stringency
type DataElementStringency int

const (
	DataElementStringencyComparable DataElementStringency = iota
	DataElementStringencyFullySpecified
	DataElementStringencyEquivalent
	DataElementStringencyConvertable
	DataElementStringencyScaleable
	DataElementStringencyFlexible
)

func (code DataElementStringency) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DataElementStringency) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "comparable":
		*code = DataElementStringencyComparable
	case "fully-specified":
		*code = DataElementStringencyFullySpecified
	case "equivalent":
		*code = DataElementStringencyEquivalent
	case "convertable":
		*code = DataElementStringencyConvertable
	case "scaleable":
		*code = DataElementStringencyScaleable
	case "flexible":
		*code = DataElementStringencyFlexible
	default:
		return fmt.Errorf("unknown DataElementStringency code `%s`", s)
	}
	return nil
}
func (code DataElementStringency) String() string {
	return code.Code()
}
func (code DataElementStringency) Code() string {
	switch code {
	case DataElementStringencyComparable:
		return "comparable"
	case DataElementStringencyFullySpecified:
		return "fully-specified"
	case DataElementStringencyEquivalent:
		return "equivalent"
	case DataElementStringencyConvertable:
		return "convertable"
	case DataElementStringencyScaleable:
		return "scaleable"
	case DataElementStringencyFlexible:
		return "flexible"
	}
	return "<unknown>"
}
func (code DataElementStringency) Display() string {
	switch code {
	case DataElementStringencyComparable:
		return "Comparable"
	case DataElementStringencyFullySpecified:
		return "Fully Specified"
	case DataElementStringencyEquivalent:
		return "Equivalent"
	case DataElementStringencyConvertable:
		return "Convertable"
	case DataElementStringencyScaleable:
		return "Scaleable"
	case DataElementStringencyFlexible:
		return "Flexible"
	}
	return "<unknown>"
}
func (code DataElementStringency) Definition() string {
	switch code {
	case DataElementStringencyComparable:
		return "The data element is sufficiently well-constrained that multiple pieces of data captured according to the constraints of the data element will be comparable (though in some cases, a degree of automated conversion/normalization may be required)."
	case DataElementStringencyFullySpecified:
		return "The data element is fully specified down to a single value set, single unit of measure, single data type, etc.  Multiple pieces of data associated with this data element are fully comparable."
	case DataElementStringencyEquivalent:
		return "The data element allows multiple units of measure having equivalent meaning; e.g. \"cc\" (cubic centimeter) and \"mL\" (milliliter)."
	case DataElementStringencyConvertable:
		return "The data element allows multiple units of measure that are convertable between each other (e.g. inches and centimeters) and/or allows data to be captured in multiple value sets for which a known mapping exists allowing conversion of meaning."
	case DataElementStringencyScaleable:
		return "A convertable data element where unit conversions are different only by a power of 10; e.g. g, mg, kg."
	case DataElementStringencyFlexible:
		return "The data element is unconstrained in units, choice of data types and/or choice of vocabulary such that automated comparison of data captured using the data element is not possible."
	}
	return "<unknown>"
}
