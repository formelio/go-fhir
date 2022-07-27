package fhir

import "encoding/json"

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	Id              *string    `bson:"id" json:"id"`
	Meta            *Meta      `bson:"meta" json:"meta"`
	ImplicitRules   *string    `bson:"implicitRules" json:"implicitRules"`
	Language        *string    `bson:"language" json:"language"`
	ContentType     string     `bson:"contentType,omitempty" json:"contentType,omitempty"`
	SecurityContext *Reference `bson:"securityContext" json:"securityContext"`
	Content         string     `bson:"content,omitempty" json:"content,omitempty"`
}

// OtherBinary is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBinary Binary

// MarshalJSON marshals the given Binary as JSON into a byte slice
func (r Binary) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType"`
	}{
		OtherBinary:  OtherBinary(r),
		ResourceType: "Binary",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Binary
func (r *Binary) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBinary)(r)); err != nil {
		return err
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r Binary) GetResourceType() ResourceType {
	return ResourceTypeBinary
}
