package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TestScriptRequestMethodCode is documented here http://hl7.org/fhir/ValueSet/http-operations
type TestScriptRequestMethodCode int

const (
	TestScriptRequestMethodCodeDelete TestScriptRequestMethodCode = iota
	TestScriptRequestMethodCodeGet
	TestScriptRequestMethodCodeOptions
	TestScriptRequestMethodCodePatch
	TestScriptRequestMethodCodePost
	TestScriptRequestMethodCodePut
)

func (code TestScriptRequestMethodCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestScriptRequestMethodCode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "delete":
		*code = TestScriptRequestMethodCodeDelete
	case "get":
		*code = TestScriptRequestMethodCodeGet
	case "options":
		*code = TestScriptRequestMethodCodeOptions
	case "patch":
		*code = TestScriptRequestMethodCodePatch
	case "post":
		*code = TestScriptRequestMethodCodePost
	case "put":
		*code = TestScriptRequestMethodCodePut
	default:
		return fmt.Errorf("unknown TestScriptRequestMethodCode code `%s`", s)
	}
	return nil
}
func (code TestScriptRequestMethodCode) String() string {
	return code.Code()
}
func (code TestScriptRequestMethodCode) Code() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "delete"
	case TestScriptRequestMethodCodeGet:
		return "get"
	case TestScriptRequestMethodCodeOptions:
		return "options"
	case TestScriptRequestMethodCodePatch:
		return "patch"
	case TestScriptRequestMethodCodePost:
		return "post"
	case TestScriptRequestMethodCodePut:
		return "put"
	}
	return "<unknown>"
}
func (code TestScriptRequestMethodCode) Display() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "DELETE"
	case TestScriptRequestMethodCodeGet:
		return "GET"
	case TestScriptRequestMethodCodeOptions:
		return "OPTIONS"
	case TestScriptRequestMethodCodePatch:
		return "PATCH"
	case TestScriptRequestMethodCodePost:
		return "POST"
	case TestScriptRequestMethodCodePut:
		return "PUT"
	}
	return "<unknown>"
}
func (code TestScriptRequestMethodCode) Definition() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "HTTP DELETE operation"
	case TestScriptRequestMethodCodeGet:
		return "HTTP GET operation"
	case TestScriptRequestMethodCodeOptions:
		return "HTTP OPTIONS operation"
	case TestScriptRequestMethodCodePatch:
		return "HTTP PATCH operation"
	case TestScriptRequestMethodCodePost:
		return "HTTP POST operation"
	case TestScriptRequestMethodCodePut:
		return "HTTP PUT operation"
	}
	return "<unknown>"
}
