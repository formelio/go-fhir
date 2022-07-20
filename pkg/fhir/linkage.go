package fhir

import "encoding/json"

// Linkage is documented here http://hl7.org/fhir/StructureDefinition/Linkage
type Linkage struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Active            *bool             `bson:"active" json:"active"`
	Author            *Reference        `bson:"author" json:"author"`
	Item              []LinkageItem     `bson:"item,omitempty" json:"item,omitempty"`
}
type LinkageItem struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Type              LinkageType `bson:"type,omitempty" json:"type,omitempty"`
	Resource          Reference   `bson:"resource,omitempty" json:"resource,omitempty"`
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

// UnmarshalLinkage unmarshalls a Linkage.
func UnmarshalLinkage(b []byte) (Linkage, error) {
	var linkage Linkage
	if err := json.Unmarshal(b, &linkage); err != nil {
		return linkage, err
	}
	return linkage, nil
}
