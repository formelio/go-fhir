package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// HTTPVerb is documented here http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb int

const (
	HTTPVerbGET HTTPVerb = iota
	HTTPVerbPOST
	HTTPVerbPUT
	HTTPVerbDELETE
)

func (code HTTPVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *HTTPVerb) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "GET":
		*code = HTTPVerbGET
	case "POST":
		*code = HTTPVerbPOST
	case "PUT":
		*code = HTTPVerbPUT
	case "DELETE":
		*code = HTTPVerbDELETE
	default:
		return fmt.Errorf("unknown HTTPVerb code `%s`", s)
	}
	return nil
}
func (code HTTPVerb) String() string {
	return code.Code()
}
func (code HTTPVerb) Code() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	}
	return "<unknown>"
}
func (code HTTPVerb) Display() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	}
	return "<unknown>"
}
func (code HTTPVerb) Definition() string {
	switch code {
	case HTTPVerbGET:
		return "HTTP GET"
	case HTTPVerbPOST:
		return "HTTP POST"
	case HTTPVerbPUT:
		return "HTTP PUT"
	case HTTPVerbDELETE:
		return "HTTP DELETE"
	}
	return "<unknown>"
}
