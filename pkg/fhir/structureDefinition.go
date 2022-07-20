package fhir

import "encoding/json"

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `bson:"id" json:"id"`
	Meta              *Meta                            `bson:"meta" json:"meta"`
	ImplicitRules     *string                          `bson:"implicitRules" json:"implicitRules"`
	Language          *string                          `bson:"language" json:"language"`
	Text              *Narrative                       `bson:"text" json:"text"`
	Contained         []json.RawMessage                `bson:"contained" json:"contained"`
	Extension         []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                     `bson:"identifier" json:"identifier"`
	Version           *string                          `bson:"version" json:"version"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                          `bson:"title" json:"title"`
	Status            PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                            `bson:"experimental" json:"experimental"`
	Date              *string                          `bson:"date" json:"date"`
	Publisher         *string                          `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                  `bson:"contact" json:"contact"`
	Description       *string                          `bson:"description" json:"description"`
	UseContext        []UsageContext                   `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept                `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                          `bson:"purpose" json:"purpose"`
	Copyright         *string                          `bson:"copyright" json:"copyright"`
	Keyword           []Coding                         `bson:"keyword" json:"keyword"`
	FhirVersion       *string                          `bson:"fhirVersion" json:"fhirVersion"`
	Mapping           []StructureDefinitionMapping     `bson:"mapping" json:"mapping"`
	Kind              StructureDefinitionKind          `bson:"kind,omitempty" json:"kind,omitempty"`
	Abstract          bool                             `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType       *ExtensionContext                `bson:"contextType" json:"contextType"`
	Context           []string                         `bson:"context" json:"context"`
	ContextInvariant  []string                         `bson:"contextInvariant" json:"contextInvariant"`
	Type              string                           `bson:"type,omitempty" json:"type,omitempty"`
	BaseDefinition    *string                          `bson:"baseDefinition" json:"baseDefinition"`
	Derivation        *TypeDerivationRule              `bson:"derivation" json:"derivation"`
	Snapshot          *StructureDefinitionSnapshot     `bson:"snapshot" json:"snapshot"`
	Differential      *StructureDefinitionDifferential `bson:"differential" json:"differential"`
}
type StructureDefinitionMapping struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identity          string      `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri               *string     `bson:"uri" json:"uri"`
	Name              *string     `bson:"name" json:"name"`
	Comment           *string     `bson:"comment" json:"comment"`
}
type StructureDefinitionSnapshot struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Element           []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}
type StructureDefinitionDifferential struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Element           []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}
type OtherStructureDefinition StructureDefinition

// MarshalJSON marshals the given StructureDefinition as JSON into a byte slice
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

// UnmarshalStructureDefinition unmarshalls a StructureDefinition.
func UnmarshalStructureDefinition(b []byte) (StructureDefinition, error) {
	var structureDefinition StructureDefinition
	if err := json.Unmarshal(b, &structureDefinition); err != nil {
		return structureDefinition, err
	}
	return structureDefinition, nil
}
