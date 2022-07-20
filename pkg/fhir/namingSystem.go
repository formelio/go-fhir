package fhir

import "encoding/json"

// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Kind              NamingSystemType       `bson:"kind,omitempty" json:"kind,omitempty"`
	Date              string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail        `bson:"contact" json:"contact"`
	Responsible       *string                `bson:"responsible" json:"responsible"`
	Type              *CodeableConcept       `bson:"type" json:"type"`
	Description       *string                `bson:"description" json:"description"`
	UseContext        []UsageContext         `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction" json:"jurisdiction"`
	Usage             *string                `bson:"usage" json:"usage"`
	UniqueId          []NamingSystemUniqueId `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
	ReplacedBy        *Reference             `bson:"replacedBy" json:"replacedBy"`
}
type NamingSystemUniqueId struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Type              NamingSystemIdentifierType `bson:"type,omitempty" json:"type,omitempty"`
	Value             string                     `bson:"value,omitempty" json:"value,omitempty"`
	Preferred         *bool                      `bson:"preferred" json:"preferred"`
	Comment           *string                    `bson:"comment" json:"comment"`
	Period            *Period                    `bson:"period" json:"period"`
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

// UnmarshalNamingSystem unmarshalls a NamingSystem.
func UnmarshalNamingSystem(b []byte) (NamingSystem, error) {
	var namingSystem NamingSystem
	if err := json.Unmarshal(b, &namingSystem); err != nil {
		return namingSystem, err
	}
	return namingSystem, nil
}
