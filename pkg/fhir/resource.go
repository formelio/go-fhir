package fhir

import "encoding/json"

// Resource is documented here http://hl7.org/fhir/StructureDefinition/Resource
type Resource struct {
	Id            *string `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string `bson:"language,omitempty" json:"language,omitempty"`
}
type OtherResource Resource

// MarshalJSON marshals the given Resource as JSON into a byte slice
func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResource
		ResourceType string `json:"resourceType"`
	}{
		OtherResource: OtherResource(r),
		ResourceType:  "Resource",
	})
}

// UnmarshalResource unmarshals a Resource.
func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}
