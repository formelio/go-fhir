package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SystemVersionProcessingMode is documented here http://hl7.org/fhir/ValueSet/system-version-processing-mode
type SystemVersionProcessingMode int

const (
	SystemVersionProcessingModeDefault SystemVersionProcessingMode = iota
	SystemVersionProcessingModeCheck
	SystemVersionProcessingModeOverride
)

func (code SystemVersionProcessingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SystemVersionProcessingMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "default":
		*code = SystemVersionProcessingModeDefault
	case "check":
		*code = SystemVersionProcessingModeCheck
	case "override":
		*code = SystemVersionProcessingModeOverride
	default:
		return fmt.Errorf("unknown SystemVersionProcessingMode code `%s`", s)
	}
	return nil
}
func (code SystemVersionProcessingMode) String() string {
	return code.Code()
}
func (code SystemVersionProcessingMode) Code() string {
	switch code {
	case SystemVersionProcessingModeDefault:
		return "default"
	case SystemVersionProcessingModeCheck:
		return "check"
	case SystemVersionProcessingModeOverride:
		return "override"
	}
	return "<unknown>"
}
func (code SystemVersionProcessingMode) Display() string {
	switch code {
	case SystemVersionProcessingModeDefault:
		return "Default Version"
	case SystemVersionProcessingModeCheck:
		return "Check ValueSet Version"
	case SystemVersionProcessingModeOverride:
		return "Override ValueSet Version"
	}
	return "<unknown>"
}
func (code SystemVersionProcessingMode) Definition() string {
	switch code {
	case SystemVersionProcessingModeDefault:
		return "Use this version of the code system if a value set doesn't specify a version"
	case SystemVersionProcessingModeCheck:
		return "Use this version of the code system. If a value set specifies a different version, the expansion operation should fail"
	case SystemVersionProcessingModeOverride:
		return "Use this version of the code system irrespective of which version is specified by a value set. Note that this has obvious safety issues, in that it may result in a value set expansion giving a different list of codes that is both wrong and unsafe, and implementers should only use this capability reluctantly. It primarily exists to deal with situations where specifications have fallen into decay as time passes. If a  version is override, the version used SHALL explicitly be represented in the expansion parameters"
	}
	return "<unknown>"
}
