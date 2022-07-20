package fhir

import "encoding/json"

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Active            *bool                 `bson:"active" json:"active"`
	Type              []CodeableConcept     `bson:"type" json:"type"`
	Name              *string               `bson:"name" json:"name"`
	Alias             []string              `bson:"alias" json:"alias"`
	Telecom           []ContactPoint        `bson:"telecom" json:"telecom"`
	Address           []Address             `bson:"address" json:"address"`
	PartOf            *Reference            `bson:"partOf" json:"partOf"`
	Contact           []OrganizationContact `bson:"contact" json:"contact"`
	Endpoint          []Reference           `bson:"endpoint" json:"endpoint"`
}
type OrganizationContact struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Purpose           *CodeableConcept `bson:"purpose" json:"purpose"`
	Name              *HumanName       `bson:"name" json:"name"`
	Telecom           []ContactPoint   `bson:"telecom" json:"telecom"`
	Address           *Address         `bson:"address" json:"address"`
}
type OtherOrganization Organization

// MarshalJSON marshals the given Organization as JSON into a byte slice
func (r Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganization
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganization: OtherOrganization(r),
		ResourceType:      "Organization",
	})
}

// UnmarshalOrganization unmarshalls a Organization.
func UnmarshalOrganization(b []byte) (Organization, error) {
	var organization Organization
	if err := json.Unmarshal(b, &organization); err != nil {
		return organization, err
	}
	return organization, nil
}
