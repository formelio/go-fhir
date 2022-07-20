package fhir

import "encoding/json"

// BodySite is documented here http://hl7.org/fhir/StructureDefinition/BodySite
type BodySite struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool             `bson:"active,omitempty" json:"active,omitempty"`
	Code              *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Qualifier         []CodeableConcept `bson:"qualifier,omitempty" json:"qualifier,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	Image             []Attachment      `bson:"image,omitempty" json:"image,omitempty"`
	Patient           Reference         `bson:"patient" json:"patient"`
}
type OtherBodySite BodySite

// MarshalJSON marshals the given BodySite as JSON into a byte slice
func (r BodySite) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBodySite
		ResourceType string `json:"resourceType"`
	}{
		OtherBodySite: OtherBodySite(r),
		ResourceType:  "BodySite",
	})
}

// UnmarshalBodySite unmarshals a BodySite.
func UnmarshalBodySite(b []byte) (BodySite, error) {
	var bodySite BodySite
	if err := json.Unmarshal(b, &bodySite); err != nil {
		return bodySite, err
	}
	return bodySite, nil
}
