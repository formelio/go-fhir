package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IdentifierUse is documented here http://hl7.org/fhir/ValueSet/identifier-use
type IdentifierUse int

const (
	IdentifierUseUsual IdentifierUse = iota
	IdentifierUseOfficial
	IdentifierUseTemp
	IdentifierUseSecondary
)

func (code IdentifierUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IdentifierUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "usual":
		*code = IdentifierUseUsual
	case "official":
		*code = IdentifierUseOfficial
	case "temp":
		*code = IdentifierUseTemp
	case "secondary":
		*code = IdentifierUseSecondary
	default:
		return fmt.Errorf("unknown IdentifierUse code `%s`", s)
	}
	return nil
}
func (code IdentifierUse) String() string {
	return code.Code()
}
func (code IdentifierUse) Code() string {
	switch code {
	case IdentifierUseUsual:
		return "usual"
	case IdentifierUseOfficial:
		return "official"
	case IdentifierUseTemp:
		return "temp"
	case IdentifierUseSecondary:
		return "secondary"
	}
	return "<unknown>"
}
func (code IdentifierUse) Display() string {
	switch code {
	case IdentifierUseUsual:
		return "Usual"
	case IdentifierUseOfficial:
		return "Official"
	case IdentifierUseTemp:
		return "Temp"
	case IdentifierUseSecondary:
		return "Secondary"
	}
	return "<unknown>"
}
func (code IdentifierUse) Definition() string {
	switch code {
	case IdentifierUseUsual:
		return "The identifier recommended for display and use in real-world interactions."
	case IdentifierUseOfficial:
		return "The identifier considered to be most trusted for the identification of this item."
	case IdentifierUseTemp:
		return "A temporary identifier."
	case IdentifierUseSecondary:
		return "An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context."
	}
	return "<unknown>"
}
