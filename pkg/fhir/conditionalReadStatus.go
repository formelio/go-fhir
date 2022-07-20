package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConditionalReadStatus is documented here http://hl7.org/fhir/ValueSet/conditional-read-status
type ConditionalReadStatus int

const (
	ConditionalReadStatusNotSupported ConditionalReadStatus = iota
	ConditionalReadStatusModifiedSince
	ConditionalReadStatusNotMatch
	ConditionalReadStatusFullSupport
)

func (code ConditionalReadStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConditionalReadStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-supported":
		*code = ConditionalReadStatusNotSupported
	case "modified-since":
		*code = ConditionalReadStatusModifiedSince
	case "not-match":
		*code = ConditionalReadStatusNotMatch
	case "full-support":
		*code = ConditionalReadStatusFullSupport
	default:
		return fmt.Errorf("unknown ConditionalReadStatus code `%s`", s)
	}
	return nil
}
func (code ConditionalReadStatus) String() string {
	return code.Code()
}
func (code ConditionalReadStatus) Code() string {
	switch code {
	case ConditionalReadStatusNotSupported:
		return "not-supported"
	case ConditionalReadStatusModifiedSince:
		return "modified-since"
	case ConditionalReadStatusNotMatch:
		return "not-match"
	case ConditionalReadStatusFullSupport:
		return "full-support"
	}
	return "<unknown>"
}
func (code ConditionalReadStatus) Display() string {
	switch code {
	case ConditionalReadStatusNotSupported:
		return "Not Supported"
	case ConditionalReadStatusModifiedSince:
		return "If-Modified-Since"
	case ConditionalReadStatusNotMatch:
		return "If-None-Match"
	case ConditionalReadStatusFullSupport:
		return "Full Support"
	}
	return "<unknown>"
}
func (code ConditionalReadStatus) Definition() string {
	switch code {
	case ConditionalReadStatusNotSupported:
		return "No support for conditional deletes."
	case ConditionalReadStatusModifiedSince:
		return "Conditional reads are supported, but only with the If-Modified-Since HTTP Header."
	case ConditionalReadStatusNotMatch:
		return "Conditional reads are supported, but only with the If-None-Match HTTP Header."
	case ConditionalReadStatusFullSupport:
		return "Conditional reads are supported, with both If-Modified-Since and If-None-Match HTTP Headers."
	}
	return "<unknown>"
}
