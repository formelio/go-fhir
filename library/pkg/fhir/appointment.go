package fhir

import "encoding/json"

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
	Id                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                   `bson:"status" json:"status"`
	ServiceCategory       *CodeableConcept         `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept        `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty             []CodeableConcept        `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept         `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Reason                []CodeableConcept        `bson:"reason,omitempty" json:"reason,omitempty"`
	Priority              *int                     `bson:"priority,omitempty" json:"priority,omitempty"`
	Description           *string                  `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInformation []Reference              `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Start                 *string                  `bson:"start,omitempty" json:"start,omitempty"`
	End                   *string                  `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration       *int                     `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot                  []Reference              `bson:"slot,omitempty" json:"slot,omitempty"`
	Created               *string                  `bson:"created,omitempty" json:"created,omitempty"`
	Comment               *string                  `bson:"comment,omitempty" json:"comment,omitempty"`
	IncomingReferral      []Reference              `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant           []AppointmentParticipant `bson:"participant" json:"participant"`
	RequestedPeriod       []Period                 `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}
type AppointmentParticipant struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Required          *string           `bson:"required,omitempty" json:"required,omitempty"`
	Status            string            `bson:"status" json:"status"`
}
type OtherAppointment Appointment

// MarshalJSON marshals the given Appointment as JSON into a byte slice
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}

// UnmarshalAppointment unmarshals a Appointment.
func UnmarshalAppointment(b []byte) (Appointment, error) {
	var appointment Appointment
	if err := json.Unmarshal(b, &appointment); err != nil {
		return appointment, err
	}
	return appointment, nil
}
