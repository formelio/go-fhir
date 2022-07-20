package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FHIRDeviceStatus is documented here http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus int

const (
	FHIRDeviceStatusActive FHIRDeviceStatus = iota
	FHIRDeviceStatusInactive
	FHIRDeviceStatusEnteredInError
	FHIRDeviceStatusUnknown
)

func (code FHIRDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRDeviceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FHIRDeviceStatusActive
	case "inactive":
		*code = FHIRDeviceStatusInactive
	case "entered-in-error":
		*code = FHIRDeviceStatusEnteredInError
	case "unknown":
		*code = FHIRDeviceStatusUnknown
	default:
		return fmt.Errorf("unknown FHIRDeviceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRDeviceStatus) String() string {
	return code.Code()
}
func (code FHIRDeviceStatus) Code() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "active"
	case FHIRDeviceStatusInactive:
		return "inactive"
	case FHIRDeviceStatusEnteredInError:
		return "entered-in-error"
	case FHIRDeviceStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Display() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "Active"
	case FHIRDeviceStatusInactive:
		return "Inactive"
	case FHIRDeviceStatusEnteredInError:
		return "Entered in Error"
	case FHIRDeviceStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Definition() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "The Device is available for use.  Note: This means for *implanted devices*  the device is implanted in the patient."
	case FHIRDeviceStatusInactive:
		return "The Device is no longer available for use (e.g. lost, expired, damaged).  Note: This means for *implanted devices*  the device has been removed from the patient."
	case FHIRDeviceStatusEnteredInError:
		return "The Device was entered in error and voided."
	case FHIRDeviceStatusUnknown:
		return "The status of the device has not been determined."
	}
	return "<unknown>"
}
