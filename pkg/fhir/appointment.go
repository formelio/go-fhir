package fhir

import "encoding/json"

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
	Id                    *string                  `bson:"id" json:"id"`
	Meta                  *Meta                    `bson:"meta" json:"meta"`
	ImplicitRules         *string                  `bson:"implicitRules" json:"implicitRules"`
	Language              *string                  `bson:"language" json:"language"`
	Text                  *Narrative               `bson:"text" json:"text"`
	Contained             []json.RawMessage        `bson:"contained" json:"contained"`
	Extension             []Extension              `bson:"extension" json:"extension"`
	ModifierExtension     []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            []Identifier             `bson:"identifier" json:"identifier"`
	Status                AppointmentStatus        `bson:"status,omitempty" json:"status,omitempty"`
	ServiceCategory       *CodeableConcept         `bson:"serviceCategory" json:"serviceCategory"`
	ServiceType           []CodeableConcept        `bson:"serviceType" json:"serviceType"`
	Specialty             []CodeableConcept        `bson:"specialty" json:"specialty"`
	AppointmentType       *CodeableConcept         `bson:"appointmentType" json:"appointmentType"`
	Reason                []CodeableConcept        `bson:"reason" json:"reason"`
	Indication            []Reference              `bson:"indication" json:"indication"`
	Priority              *int                     `bson:"priority" json:"priority"`
	Description           *string                  `bson:"description" json:"description"`
	SupportingInformation []Reference              `bson:"supportingInformation" json:"supportingInformation"`
	Start                 *string                  `bson:"start" json:"start"`
	End                   *string                  `bson:"end" json:"end"`
	MinutesDuration       *int                     `bson:"minutesDuration" json:"minutesDuration"`
	Slot                  []Reference              `bson:"slot" json:"slot"`
	Created               *string                  `bson:"created" json:"created"`
	Comment               *string                  `bson:"comment" json:"comment"`
	IncomingReferral      []Reference              `bson:"incomingReferral" json:"incomingReferral"`
	Participant           []AppointmentParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	RequestedPeriod       []Period                 `bson:"requestedPeriod" json:"requestedPeriod"`
}
type AppointmentParticipant struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Type              []CodeableConcept    `bson:"type" json:"type"`
	Actor             *Reference           `bson:"actor" json:"actor"`
	Required          *ParticipantRequired `bson:"required" json:"required"`
	Status            ParticipationStatus  `bson:"status,omitempty" json:"status,omitempty"`
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

// UnmarshalAppointment unmarshalls a Appointment.
func UnmarshalAppointment(b []byte) (Appointment, error) {
	var appointment Appointment
	if err := json.Unmarshal(b, &appointment); err != nil {
		return appointment, err
	}
	return appointment, nil
}
