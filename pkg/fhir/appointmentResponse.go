package fhir

import "encoding/json"

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string             `bson:"id" json:"id"`
	Meta              *Meta               `bson:"meta" json:"meta"`
	ImplicitRules     *string             `bson:"implicitRules" json:"implicitRules"`
	Language          *string             `bson:"language" json:"language"`
	Text              *Narrative          `bson:"text" json:"text"`
	RawContained      []json.RawMessage   `bson:"contained" json:"contained"`
	Contained         []IResource         `bson:"-" json:"-"`
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

// OtherAppointmentResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAppointmentResponse AppointmentResponse

// MarshalJSON marshals the given AppointmentResponse as JSON into a byte slice
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherAppointmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointmentResponse: OtherAppointmentResponse(r),
		ResourceType:             "AppointmentResponse",
	})
}

// UnmarshalJSON unmarshals the given byte slice into AppointmentResponse
func (r *AppointmentResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAppointmentResponse)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r AppointmentResponse) GetResourceType() ResourceType {
	return ResourceTypeAppointmentResponse
}
