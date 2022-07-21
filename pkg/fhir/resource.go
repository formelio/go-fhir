package fhir

import "encoding/json"

// Resource is documented here http://hl7.org/fhir/StructureDefinition/Resource
type Resource struct {
	Id            *string      `bson:"id" json:"id"`
	Meta          *Meta        `bson:"meta" json:"meta"`
	ImplicitRules *string      `bson:"implicitRules" json:"implicitRules"`
	Language      *string      `bson:"language" json:"language"`
	ResourceType  ResourceType `bson:"resourceType" json:"resourceType"`
}

// UnmarshalResource unmarshalls a Resource.
func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}

// type OtherResource Resource
type IResource interface {
	GetResourceType() ResourceType
}

func (r Resource) GetResourceType() ResourceType {
	return r.ResourceType
}
