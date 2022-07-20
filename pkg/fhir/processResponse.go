package fhir

import "encoding/json"

// ProcessResponse is documented here http://hl7.org/fhir/StructureDefinition/ProcessResponse
type ProcessResponse struct {
	Id                   *string                      `bson:"id" json:"id"`
	Meta                 *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules        *string                      `bson:"implicitRules" json:"implicitRules"`
	Language             *string                      `bson:"language" json:"language"`
	Text                 *Narrative                   `bson:"text" json:"text"`
	Contained            []json.RawMessage            `bson:"contained" json:"contained"`
	Extension            []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier                 `bson:"identifier" json:"identifier"`
	Status               *string                      `bson:"status" json:"status"`
	Created              *string                      `bson:"created" json:"created"`
	Organization         *Reference                   `bson:"organization" json:"organization"`
	Request              *Reference                   `bson:"request" json:"request"`
	Outcome              *CodeableConcept             `bson:"outcome" json:"outcome"`
	Disposition          *string                      `bson:"disposition" json:"disposition"`
	RequestProvider      *Reference                   `bson:"requestProvider" json:"requestProvider"`
	RequestOrganization  *Reference                   `bson:"requestOrganization" json:"requestOrganization"`
	Form                 *CodeableConcept             `bson:"form" json:"form"`
	ProcessNote          []ProcessResponseProcessNote `bson:"processNote" json:"processNote"`
	Error                []CodeableConcept            `bson:"error" json:"error"`
	CommunicationRequest []Reference                  `bson:"communicationRequest" json:"communicationRequest"`
}
type ProcessResponseProcessNote struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	Text              *string          `bson:"text" json:"text"`
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

// UnmarshalProcessResponse unmarshalls a ProcessResponse.
func UnmarshalProcessResponse(b []byte) (ProcessResponse, error) {
	var processResponse ProcessResponse
	if err := json.Unmarshal(b, &processResponse); err != nil {
		return processResponse, err
	}
	return processResponse, nil
}
