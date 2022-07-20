package fhir

import "encoding/json"

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
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

// UnmarshalDomainResource unmarshalls a DomainResource.
func UnmarshalDomainResource(b []byte) (DomainResource, error) {
	var domainResource DomainResource
	if err := json.Unmarshal(b, &domainResource); err != nil {
		return domainResource, err
	}
	return domainResource, nil
}
