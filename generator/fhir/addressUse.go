package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AddressUse is documented here http://hl7.org/fhir/ValueSet/address-use
type AddressUse int

const (
	AddressUseHome AddressUse = iota
	AddressUseWork
	AddressUseTemp
	AddressUseOld
)

func (code AddressUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AddressUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "home":
		*code = AddressUseHome
	case "work":
		*code = AddressUseWork
	case "temp":
		*code = AddressUseTemp
	case "old":
		*code = AddressUseOld
	default:
		return fmt.Errorf("unknown AddressUse code `%s`", s)
	}
	return nil
}
func (code AddressUse) String() string {
	return code.Code()
}
func (code AddressUse) Code() string {
	switch code {
	case AddressUseHome:
		return "home"
	case AddressUseWork:
		return "work"
	case AddressUseTemp:
		return "temp"
	case AddressUseOld:
		return "old"
	}
	return "<unknown>"
}
func (code AddressUse) Display() string {
	switch code {
	case AddressUseHome:
		return "Home"
	case AddressUseWork:
		return "Work"
	case AddressUseTemp:
		return "Temporary"
	case AddressUseOld:
		return "Old / Incorrect"
	}
	return "<unknown>"
}
func (code AddressUse) Definition() string {
	switch code {
	case AddressUseHome:
		return "A communication address at a home."
	case AddressUseWork:
		return "An office address. First choice for business related contacts during business hours."
	case AddressUseTemp:
		return "A temporary address. The period can provide more detailed information."
	case AddressUseOld:
		return "This address is no longer in use (or was never correct, but retained for records)."
	}
	return "<unknown>"
}
