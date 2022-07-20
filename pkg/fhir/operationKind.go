package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// OperationKind is documented here http://hl7.org/fhir/ValueSet/operation-kind
type OperationKind int

const (
	OperationKindOperation OperationKind = iota
	OperationKindQuery
)

func (code OperationKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OperationKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "operation":
		*code = OperationKindOperation
	case "query":
		*code = OperationKindQuery
	default:
		return fmt.Errorf("unknown OperationKind code `%s`", s)
	}
	return nil
}
func (code OperationKind) String() string {
	return code.Code()
}
func (code OperationKind) Code() string {
	switch code {
	case OperationKindOperation:
		return "operation"
	case OperationKindQuery:
		return "query"
	}
	return "<unknown>"
}
func (code OperationKind) Display() string {
	switch code {
	case OperationKindOperation:
		return "Operation"
	case OperationKindQuery:
		return "Query"
	}
	return "<unknown>"
}
func (code OperationKind) Definition() string {
	switch code {
	case OperationKindOperation:
		return "This operation is invoked as an operation."
	case OperationKindQuery:
		return "This operation is a named query, invoked using the search mechanism."
	}
	return "<unknown>"
}
