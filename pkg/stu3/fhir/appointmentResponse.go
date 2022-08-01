package fhir

import "encoding/json"

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource         `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Appointment       Reference           `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Start             *string             `bson:"start,omitempty" json:"start,omitempty"`
	End               *string             `bson:"end,omitempty" json:"end,omitempty"`
	ParticipantType   []CodeableConcept   `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Actor             *Reference          `bson:"actor,omitempty" json:"actor,omitempty"`
	ParticipantStatus ParticipationStatus `bson:"participantStatus,omitempty" json:"participantStatus,omitempty"`
	Comment           *string             `bson:"comment,omitempty" json:"comment,omitempty"`
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
