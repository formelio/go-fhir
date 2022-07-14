package fhir

import "encoding/json"

// ProcessResponse is documented here http://hl7.org/fhir/StructureDefinition/ProcessResponse
type ProcessResponse struct {
	Id                   *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string                      `bson:"status,omitempty" json:"status,omitempty"`
	Created              *string                      `bson:"created,omitempty" json:"created,omitempty"`
	Organization         *Reference                   `bson:"organization,omitempty" json:"organization,omitempty"`
	Request              *Reference                   `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              *CodeableConcept             `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                      `bson:"disposition,omitempty" json:"disposition,omitempty"`
	RequestProvider      *Reference                   `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization  *Reference                   `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Form                 *CodeableConcept             `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ProcessResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	Error                []CodeableConcept            `bson:"error,omitempty" json:"error,omitempty"`
	CommunicationRequest []Reference                  `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
}
type ProcessResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
}
type OtherProcessResponse ProcessResponse

// MarshalJSON marshals the given ProcessResponse as JSON into a byte slice
func (r ProcessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcessResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherProcessResponse: OtherProcessResponse(r),
		ResourceType:         "ProcessResponse",
	})
}

// UnmarshalProcessResponse unmarshals a ProcessResponse.
func UnmarshalProcessResponse(b []byte) (ProcessResponse, error) {
	var processResponse ProcessResponse
	if err := json.Unmarshal(b, &processResponse); err != nil {
		return processResponse, err
	}
	return processResponse, nil
}
