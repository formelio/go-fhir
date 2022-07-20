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

// UnmarshalBinary unmarshalls a Binary.
func UnmarshalBinary(b []byte) (Binary, error) {
	var binary Binary
	if err := json.Unmarshal(b, &binary); err != nil {
		return binary, err
	}
	return binary, nil
}
