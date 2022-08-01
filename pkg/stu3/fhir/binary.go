package fhir

import "encoding/json"

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	Id              *string    `bson:"id,omitempty" json:"id,omitempty"`
	Meta            *Meta      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules   *string    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language        *string    `bson:"language,omitempty" json:"language,omitempty"`
	ContentType     string     `bson:"contentType,omitempty" json:"contentType,omitempty"`
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
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
