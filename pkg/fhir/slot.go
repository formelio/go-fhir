package fhir

import "encoding/json"

// Slot is documented here http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	ServiceCategory   *CodeableConcept  `bson:"serviceCategory" json:"serviceCategory"`
	ServiceType       []CodeableConcept `bson:"serviceType" json:"serviceType"`
	Specialty         []CodeableConcept `bson:"specialty" json:"specialty"`
	AppointmentType   *CodeableConcept  `bson:"appointmentType" json:"appointmentType"`
	Schedule          Reference         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Status            SlotStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Start             string            `bson:"start,omitempty" json:"start,omitempty"`
	End               string            `bson:"end,omitempty" json:"end,omitempty"`
	Overbooked        *bool             `bson:"overbooked" json:"overbooked"`
	Comment           *string           `bson:"comment" json:"comment"`
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

// UnmarshalSlot unmarshalls a Slot.
func UnmarshalSlot(b []byte) (Slot, error) {
	var slot Slot
	if err := json.Unmarshal(b, &slot); err != nil {
		return slot, err
	}
	return slot, nil
}
