package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CodeSystemContentMode is documented here http://hl7.org/fhir/ValueSet/codesystem-content-mode
type CodeSystemContentMode int

const (
	CodeSystemContentModeNotPresent CodeSystemContentMode = iota
	CodeSystemContentModeExample
	CodeSystemContentModeFragment
	CodeSystemContentModeComplete
)

func (code CodeSystemContentMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSystemContentMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-present":
		*code = CodeSystemContentModeNotPresent
	case "example":
		*code = CodeSystemContentModeExample
	case "fragment":
		*code = CodeSystemContentModeFragment
	case "complete":
		*code = CodeSystemContentModeComplete
	default:
		return fmt.Errorf("unknown CodeSystemContentMode code `%s`", s)
	}
	return nil
}
func (code CodeSystemContentMode) String() string {
	return code.Code()
}
func (code CodeSystemContentMode) Code() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "not-present"
	case CodeSystemContentModeExample:
		return "example"
	case CodeSystemContentModeFragment:
		return "fragment"
	case CodeSystemContentModeComplete:
		return "complete"
	}
	return "<unknown>"
}
func (code CodeSystemContentMode) Display() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "Not Present"
	case CodeSystemContentModeExample:
		return "Example"
	case CodeSystemContentModeFragment:
		return "Fragment"
	case CodeSystemContentModeComplete:
		return "Complete"
	}
	return "<unknown>"
}
func (code CodeSystemContentMode) Definition() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "None of the concepts defined by the code system are included in the code system resource"
	case CodeSystemContentModeExample:
		return "A few representative concepts are included in the code system resource"
	case CodeSystemContentModeFragment:
		return "A subset of the code system concepts are included in the code system resource"
	case CodeSystemContentModeComplete:
		return "All the concepts defined by the code system are included in the code system resource"
	}
	return "<unknown>"
}
