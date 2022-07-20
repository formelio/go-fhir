package fhir

import "encoding/json"

// Linkage is documented here http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string       `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative    `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Active            *bool         `bson:"active,omitempty" json:"active,omitempty"`
	Item              []LinkageItem `bson:"item" json:"item"`
}
type LinkageItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Resource          Reference   `bson:"resource" json:"resource"`
}
type OtherLinkage Linkage

// MarshalJSON marshals the given Linkage as JSON into a byte slice
func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType"`
	}{
		OtherLinkage: OtherLinkage(r),
		ResourceType: "Linkage",
	})
}

// UnmarshalLinkage unmarshals a Linkage.
func UnmarshalLinkage(b []byte) (Linkage, error) {
	var linkage Linkage
	if err := json.Unmarshal(b, &linkage); err != nil {
		return linkage, err
	}
	return linkage, nil
}
