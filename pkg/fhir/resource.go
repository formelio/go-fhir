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
func UnmarshalResource(b []byte) (IResource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return nil, err
	}
	// Get the correct implementation of IResource from the ResourceType
	var resourceInterface = ResourceTypeToIResource(resource.ResourceType)
	// Unmarshal the raw resource into the correct implementation of IResource
	if err := json.Unmarshal(b, resourceInterface); err != nil {
		return nil, err
	}
	return resourceInterface, nil
}

// type OtherResource Resource
type IResource interface {
	GetResourceType() ResourceType
}

func (r Resource) GetResourceType() ResourceType {
	return r.ResourceType
}
