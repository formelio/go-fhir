package fhir

import "encoding/json"

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string             `bson:"id" json:"id"`
	Meta              *Meta               `bson:"meta" json:"meta"`
	ImplicitRules     *string             `bson:"implicitRules" json:"implicitRules"`
	Language          *string             `bson:"language" json:"language"`
	Text              *Narrative          `bson:"text" json:"text"`
	Contained         []json.RawMessage   `bson:"contained" json:"contained"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier        `bson:"identifier" json:"identifier"`
	Appointment       Reference           `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Start             *string             `bson:"start" json:"start"`
	End               *string             `bson:"end" json:"end"`
	ParticipantType   []CodeableConcept   `bson:"participantType" json:"participantType"`
	Actor             *Reference          `bson:"actor" json:"actor"`
	ParticipantStatus ParticipationStatus `bson:"participantStatus,omitempty" json:"participantStatus,omitempty"`
	Comment           *string             `bson:"comment" json:"comment"`
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

// UnmarshalAppointmentResponse unmarshalls a AppointmentResponse.
func UnmarshalAppointmentResponse(b []byte) (AppointmentResponse, error) {
	var appointmentResponse AppointmentResponse
	if err := json.Unmarshal(b, &appointmentResponse); err != nil {
		return appointmentResponse, err
	}
	return appointmentResponse, nil
}
