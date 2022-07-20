package fhir

import "encoding/json"

// EnrollmentResponse is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string          `bson:"status,omitempty" json:"status,omitempty"`
	Request             *Reference       `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         *string          `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Created             *string          `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference       `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider     *Reference       `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference       `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
}
type OtherEnrollmentResponse EnrollmentResponse

// MarshalJSON marshals the given EnrollmentResponse as JSON into a byte slice
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
}

// UnmarshalEnrollmentResponse unmarshals a EnrollmentResponse.
func UnmarshalEnrollmentResponse(b []byte) (EnrollmentResponse, error) {
	var enrollmentResponse EnrollmentResponse
	if err := json.Unmarshal(b, &enrollmentResponse); err != nil {
		return enrollmentResponse, err
	}
	return enrollmentResponse, nil
}
