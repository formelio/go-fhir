package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AppointmentStatus is documented here http://hl7.org/fhir/ValueSet/appointmentstatus
type AppointmentStatus int

const (
	AppointmentStatusProposed AppointmentStatus = iota
	AppointmentStatusPending
	AppointmentStatusBooked
	AppointmentStatusArrived
	AppointmentStatusFulfilled
	AppointmentStatusCancelled
	AppointmentStatusNoshow
	AppointmentStatusEnteredInError
)

func (code AppointmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AppointmentStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = AppointmentStatusProposed
	case "pending":
		*code = AppointmentStatusPending
	case "booked":
		*code = AppointmentStatusBooked
	case "arrived":
		*code = AppointmentStatusArrived
	case "fulfilled":
		*code = AppointmentStatusFulfilled
	case "cancelled":
		*code = AppointmentStatusCancelled
	case "noshow":
		*code = AppointmentStatusNoshow
	case "entered-in-error":
		*code = AppointmentStatusEnteredInError
	default:
		return fmt.Errorf("unknown AppointmentStatus code `%s`", s)
	}
	return nil
}
func (code AppointmentStatus) String() string {
	return code.Code()
}
func (code AppointmentStatus) Code() string {
	switch code {
	case AppointmentStatusProposed:
		return "proposed"
	case AppointmentStatusPending:
		return "pending"
	case AppointmentStatusBooked:
		return "booked"
	case AppointmentStatusArrived:
		return "arrived"
	case AppointmentStatusFulfilled:
		return "fulfilled"
	case AppointmentStatusCancelled:
		return "cancelled"
	case AppointmentStatusNoshow:
		return "noshow"
	case AppointmentStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code AppointmentStatus) Display() string {
	switch code {
	case AppointmentStatusProposed:
		return "Proposed"
	case AppointmentStatusPending:
		return "Pending"
	case AppointmentStatusBooked:
		return "Booked"
	case AppointmentStatusArrived:
		return "Arrived"
	case AppointmentStatusFulfilled:
		return "Fulfilled"
	case AppointmentStatusCancelled:
		return "Cancelled"
	case AppointmentStatusNoshow:
		return "No Show"
	case AppointmentStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}
func (code AppointmentStatus) Definition() string {
	switch code {
	case AppointmentStatusProposed:
		return "None of the participant(s) have finalized their acceptance of the appointment request, and the start/end time may not be set yet."
	case AppointmentStatusPending:
		return "Some or all of the participant(s) have not finalized their acceptance of the appointment request."
	case AppointmentStatusBooked:
		return "All participant(s) have been considered and the appointment is confirmed to go ahead at the date/times specified."
	case AppointmentStatusArrived:
		return "Some of the patients have arrived."
	case AppointmentStatusFulfilled:
		return "This appointment has completed and may have resulted in an encounter."
	case AppointmentStatusCancelled:
		return "The appointment has been cancelled."
	case AppointmentStatusNoshow:
		return "Some or all of the participant(s) have not/did not appear for the appointment (usually the patient)."
	case AppointmentStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
