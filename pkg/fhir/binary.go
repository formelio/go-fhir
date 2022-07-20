package fhir

import "encoding/json"

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
type Binary struct {
	Id              *string    `bson:"id,omitempty" json:"id,omitempty"`
	Meta            *Meta      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules   *string    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language        *string    `bson:"language,omitempty" json:"language,omitempty"`
	ContentType     string     `bson:"contentType" json:"contentType"`
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
	Content         string     `bson:"content" json:"content"`
}
type OtherBinary Binary

// MarshalJSON marshals the given Binary as JSON into a byte slice
func (r Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType"`
	}{
		OtherBinary:  OtherBinary(r),
		ResourceType: "Binary",
	})
}

// UnmarshalBinary unmarshals a Binary.
func UnmarshalBinary(b []byte) (Binary, error) {
	var binary Binary
	if err := json.Unmarshal(b, &binary); err != nil {
		return binary, err
	}
	return binary, nil
}
