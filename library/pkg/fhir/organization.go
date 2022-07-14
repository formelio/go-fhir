package fhir

import "encoding/json"

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Type              []CodeableConcept     `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string               `bson:"name,omitempty" json:"name,omitempty"`
	Alias             []string              `bson:"alias,omitempty" json:"alias,omitempty"`
	Telecom           []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []Address             `bson:"address,omitempty" json:"address,omitempty"`
	PartOf            *Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Contact           []OrganizationContact `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          []Reference           `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type OrganizationContact struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name              *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address         `bson:"address,omitempty" json:"address,omitempty"`
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

// UnmarshalOrganization unmarshals a Organization.
func UnmarshalOrganization(b []byte) (Organization, error) {
	var organization Organization
	if err := json.Unmarshal(b, &organization); err != nil {
		return organization, err
	}
	return organization, nil
}
