package fhir

import "encoding/json"

// EnrollmentRequest is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Status            *string           `bson:"status" json:"status"`
	Created           *string           `bson:"created" json:"created"`
	Insurer           *Reference        `bson:"insurer" json:"insurer"`
	Provider          *Reference        `bson:"provider" json:"provider"`
	Organization      *Reference        `bson:"organization" json:"organization"`
	Subject           *Reference        `bson:"subject" json:"subject"`
	Coverage          *Reference        `bson:"coverage" json:"coverage"`
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

// UnmarshalEnrollmentRequest unmarshalls a EnrollmentRequest.
func UnmarshalEnrollmentRequest(b []byte) (EnrollmentRequest, error) {
	var enrollmentRequest EnrollmentRequest
	if err := json.Unmarshal(b, &enrollmentRequest); err != nil {
		return enrollmentRequest, err
	}
	return enrollmentRequest, nil
}
