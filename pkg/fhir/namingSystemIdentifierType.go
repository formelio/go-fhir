package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NamingSystemIdentifierType is documented here http://hl7.org/fhir/ValueSet/namingsystem-identifier-type
type NamingSystemIdentifierType int

const (
	NamingSystemIdentifierTypeOid NamingSystemIdentifierType = iota
	NamingSystemIdentifierTypeUuid
	NamingSystemIdentifierTypeUri
	NamingSystemIdentifierTypeOther
)

func (code NamingSystemIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NamingSystemIdentifierType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "oid":
		*code = NamingSystemIdentifierTypeOid
	case "uuid":
		*code = NamingSystemIdentifierTypeUuid
	case "uri":
		*code = NamingSystemIdentifierTypeUri
	case "other":
		*code = NamingSystemIdentifierTypeOther
	default:
		return fmt.Errorf("unknown NamingSystemIdentifierType code `%s`", s)
	}
	return nil
}
func (code NamingSystemIdentifierType) String() string {
	return code.Code()
}
func (code NamingSystemIdentifierType) Code() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "oid"
	case NamingSystemIdentifierTypeUuid:
		return "uuid"
	case NamingSystemIdentifierTypeUri:
		return "uri"
	case NamingSystemIdentifierTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code NamingSystemIdentifierType) Display() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "OID"
	case NamingSystemIdentifierTypeUuid:
		return "UUID"
	case NamingSystemIdentifierTypeUri:
		return "URI"
	case NamingSystemIdentifierTypeOther:
		return "Other"
	}
	return "<unknown>"
}
func (code NamingSystemIdentifierType) Definition() string {
	switch code {
	case NamingSystemIdentifierTypeOid:
		return "An ISO object identifier; e.g. 1.2.3.4.5."
	case NamingSystemIdentifierTypeUuid:
		return "A universally unique identifier of the form a5afddf4-e880-459b-876e-e4591b0acc11."
	case NamingSystemIdentifierTypeUri:
		return "A uniform resource identifier (ideally a URL - uniform resource locator); e.g. http://unitsofmeasure.org."
	case NamingSystemIdentifierTypeOther:
		return "Some other type of unique identifier; e.g. HL7-assigned reserved string such as LN for LOINC."
	}
	return "<unknown>"
}
