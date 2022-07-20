package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchComparator is documented here http://hl7.org/fhir/ValueSet/search-comparator
type SearchComparator int

const (
	SearchComparatorEq SearchComparator = iota
	SearchComparatorNe
	SearchComparatorGt
	SearchComparatorLt
	SearchComparatorGe
	SearchComparatorLe
	SearchComparatorSa
	SearchComparatorEb
	SearchComparatorAp
)

func (code SearchComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SearchComparator) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "eq":
		*code = SearchComparatorEq
	case "ne":
		*code = SearchComparatorNe
	case "gt":
		*code = SearchComparatorGt
	case "lt":
		*code = SearchComparatorLt
	case "ge":
		*code = SearchComparatorGe
	case "le":
		*code = SearchComparatorLe
	case "sa":
		*code = SearchComparatorSa
	case "eb":
		*code = SearchComparatorEb
	case "ap":
		*code = SearchComparatorAp
	default:
		return fmt.Errorf("unknown SearchComparator code `%s`", s)
	}
	return nil
}
func (code SearchComparator) String() string {
	return code.Code()
}
func (code SearchComparator) Code() string {
	switch code {
	case SearchComparatorEq:
		return "eq"
	case SearchComparatorNe:
		return "ne"
	case SearchComparatorGt:
		return "gt"
	case SearchComparatorLt:
		return "lt"
	case SearchComparatorGe:
		return "ge"
	case SearchComparatorLe:
		return "le"
	case SearchComparatorSa:
		return "sa"
	case SearchComparatorEb:
		return "eb"
	case SearchComparatorAp:
		return "ap"
	}
	return "<unknown>"
}
func (code SearchComparator) Display() string {
	switch code {
	case SearchComparatorEq:
		return "Equals"
	case SearchComparatorNe:
		return "Not Equals"
	case SearchComparatorGt:
		return "Greater Than"
	case SearchComparatorLt:
		return "Less Then"
	case SearchComparatorGe:
		return "Greater or Equals"
	case SearchComparatorLe:
		return "Less of Equal"
	case SearchComparatorSa:
		return "Starts After"
	case SearchComparatorEb:
		return "Ends Before"
	case SearchComparatorAp:
		return "Approximately"
	}
	return "<unknown>"
}
func (code SearchComparator) Definition() string {
	switch code {
	case SearchComparatorEq:
		return "the value for the parameter in the resource is equal to the provided value"
	case SearchComparatorNe:
		return "the value for the parameter in the resource is not equal to the provided value"
	case SearchComparatorGt:
		return "the value for the parameter in the resource is greater than the provided value"
	case SearchComparatorLt:
		return "the value for the parameter in the resource is less than the provided value"
	case SearchComparatorGe:
		return "the value for the parameter in the resource is greater or equal to the provided value"
	case SearchComparatorLe:
		return "the value for the parameter in the resource is less or equal to the provided value"
	case SearchComparatorSa:
		return "the value for the parameter in the resource starts after the provided value"
	case SearchComparatorEb:
		return "the value for the parameter in the resource ends before the provided value"
	case SearchComparatorAp:
		return "the value for the parameter in the resource is approximately the same to the provided value."
	}
	return "<unknown>"
}
