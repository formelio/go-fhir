package fhir

import "encoding/json"

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherDomainResource DomainResource

// MarshalJSON marshals the given DomainResource as JSON into a byte slice
func (r DomainResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDomainResource
		ResourceType string `json:"resourceType"`
	}{
		OtherDomainResource: OtherDomainResource(r),
		ResourceType:        "DomainResource",
	})
}

// UnmarshalDomainResource unmarshals a DomainResource.
func UnmarshalDomainResource(b []byte) (DomainResource, error) {
	var domainResource DomainResource
	if err := json.Unmarshal(b, &domainResource); err != nil {
		return domainResource, err
	}
	return domainResource, nil
}
