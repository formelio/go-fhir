package fhir

import "encoding/json"

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Appointment       Reference         `bson:"appointment" json:"appointment"`
	Start             *string           `bson:"start,omitempty" json:"start,omitempty"`
	End               *string           `bson:"end,omitempty" json:"end,omitempty"`
	ParticipantType   []CodeableConcept `bson:"participantType,omitempty" json:"participantType,omitempty"`
	ParticipantStatus string            `bson:"participantStatus" json:"participantStatus"`
	Comment           *string           `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherAppointmentResponse AppointmentResponse

// MarshalJSON marshals the given AppointmentResponse as JSON into a byte slice
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointmentResponse: OtherAppointmentResponse(r),
		ResourceType:             "AppointmentResponse",
	})
}

// UnmarshalAppointmentResponse unmarshals a AppointmentResponse.
func UnmarshalAppointmentResponse(b []byte) (AppointmentResponse, error) {
	var appointmentResponse AppointmentResponse
	if err := json.Unmarshal(b, &appointmentResponse); err != nil {
		return appointmentResponse, err
	}
	return appointmentResponse, nil
}
