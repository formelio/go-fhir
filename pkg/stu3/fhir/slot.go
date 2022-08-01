package fhir

import "encoding/json"

// Slot is documented here http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ServiceCategory   *CodeableConcept  `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType   *CodeableConcept  `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Schedule          Reference         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Status            SlotStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Start             string            `bson:"start,omitempty" json:"start,omitempty"`
	End               string            `bson:"end,omitempty" json:"end,omitempty"`
	Overbooked        *bool             `bson:"overbooked,omitempty" json:"overbooked,omitempty"`
	Comment           *string           `bson:"comment,omitempty" json:"comment,omitempty"`
}

// OtherSlot is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSlot Slot

// MarshalJSON marshals the given Slot as JSON into a byte slice
func (r Slot) MarshalJSON() ([]byte, error) {
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
		OtherSlot
		ResourceType string `json:"resourceType"`
	}{
		OtherSlot:    OtherSlot(r),
		ResourceType: "Slot",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Slot
func (r *Slot) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSlot)(r)); err != nil {
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
func (r Slot) GetResourceType() ResourceType {
	return ResourceTypeSlot
}
