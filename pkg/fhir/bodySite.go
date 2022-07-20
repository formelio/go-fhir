package fhir

import "encoding/json"

// BodySite is documented here http://hl7.org/fhir/StructureDefinition/BodySite
type BodySite struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Active            *bool             `bson:"active" json:"active"`
	Code              *CodeableConcept  `bson:"code" json:"code"`
	Qualifier         []CodeableConcept `bson:"qualifier" json:"qualifier"`
	Description       *string           `bson:"description" json:"description"`
	Image             []Attachment      `bson:"image" json:"image"`
	Patient           Reference         `bson:"patient,omitempty" json:"patient,omitempty"`
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

// UnmarshalBodySite unmarshalls a BodySite.
func UnmarshalBodySite(b []byte) (BodySite, error) {
	var bodySite BodySite
	if err := json.Unmarshal(b, &bodySite); err != nil {
		return bodySite, err
	}
	return bodySite, nil
}
