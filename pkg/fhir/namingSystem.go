package fhir

import "encoding/json"

// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                 `bson:"name" json:"name"`
	Status            string                 `bson:"status" json:"status"`
	Kind              string                 `bson:"kind" json:"kind"`
	Date              string                 `bson:"date" json:"date"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible       *string                `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Type              *CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `bson:"uniqueId" json:"uniqueId"`
	ReplacedBy        *Reference             `bson:"replacedBy,omitempty" json:"replacedBy,omitempty"`
}
type NamingSystemUniqueId struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Value             string      `bson:"value" json:"value"`
	Preferred         *bool       `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherNamingSystem NamingSystem

// MarshalJSON marshals the given NamingSystem as JSON into a byte slice
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}

// UnmarshalNamingSystem unmarshals a NamingSystem.
func UnmarshalNamingSystem(b []byte) (NamingSystem, error) {
	var namingSystem NamingSystem
	if err := json.Unmarshal(b, &namingSystem); err != nil {
		return namingSystem, err
	}
	return namingSystem, nil
}
