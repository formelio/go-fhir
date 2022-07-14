package fhir

import "encoding/json"

// EnrollmentRequest is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string      `bson:"status,omitempty" json:"status,omitempty"`
	Created           *string      `bson:"created,omitempty" json:"created,omitempty"`
	Insurer           *Reference   `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider          *Reference   `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization      *Reference   `bson:"organization,omitempty" json:"organization,omitempty"`
	Subject           *Reference   `bson:"subject,omitempty" json:"subject,omitempty"`
	Coverage          *Reference   `bson:"coverage,omitempty" json:"coverage,omitempty"`
}
type OtherEnrollmentRequest EnrollmentRequest

// MarshalJSON marshals the given EnrollmentRequest as JSON into a byte slice
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
}

// UnmarshalEnrollmentRequest unmarshals a EnrollmentRequest.
func UnmarshalEnrollmentRequest(b []byte) (EnrollmentRequest, error) {
	var enrollmentRequest EnrollmentRequest
	if err := json.Unmarshal(b, &enrollmentRequest); err != nil {
		return enrollmentRequest, err
	}
	return enrollmentRequest, nil
}
