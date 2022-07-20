package fhir

import "encoding/json"

// EnrollmentResponse is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	Id                  *string           `bson:"id" json:"id"`
	Meta                *Meta             `bson:"meta" json:"meta"`
	ImplicitRules       *string           `bson:"implicitRules" json:"implicitRules"`
	Language            *string           `bson:"language" json:"language"`
	Text                *Narrative        `bson:"text" json:"text"`
	Contained           []json.RawMessage `bson:"contained" json:"contained"`
	Extension           []Extension       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier      `bson:"identifier" json:"identifier"`
	Status              *string           `bson:"status" json:"status"`
	Request             *Reference        `bson:"request" json:"request"`
	Outcome             *CodeableConcept  `bson:"outcome" json:"outcome"`
	Disposition         *string           `bson:"disposition" json:"disposition"`
	Created             *string           `bson:"created" json:"created"`
	Organization        *Reference        `bson:"organization" json:"organization"`
	RequestProvider     *Reference        `bson:"requestProvider" json:"requestProvider"`
	RequestOrganization *Reference        `bson:"requestOrganization" json:"requestOrganization"`
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

// UnmarshalEnrollmentResponse unmarshalls a EnrollmentResponse.
func UnmarshalEnrollmentResponse(b []byte) (EnrollmentResponse, error) {
	var enrollmentResponse EnrollmentResponse
	if err := json.Unmarshal(b, &enrollmentResponse); err != nil {
		return enrollmentResponse, err
	}
	return enrollmentResponse, nil
}
