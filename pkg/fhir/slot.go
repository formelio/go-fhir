package fhir

import "encoding/json"

// Slot is documented here http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ServiceCategory   *CodeableConcept  `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType   *CodeableConcept  `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Schedule          Reference         `bson:"schedule" json:"schedule"`
	Status            string            `bson:"status" json:"status"`
	Start             string            `bson:"start" json:"start"`
	End               string            `bson:"end" json:"end"`
	Overbooked        *bool             `bson:"overbooked,omitempty" json:"overbooked,omitempty"`
	Comment           *string           `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherSlot Slot

// MarshalJSON marshals the given Slot as JSON into a byte slice
func (r Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSlot
		ResourceType string `json:"resourceType"`
	}{
		OtherSlot:    OtherSlot(r),
		ResourceType: "Slot",
	})
}

// UnmarshalSlot unmarshals a Slot.
func UnmarshalSlot(b []byte) (Slot, error) {
	var slot Slot
	if err := json.Unmarshal(b, &slot); err != nil {
		return slot, err
	}
	return slot, nil
}
