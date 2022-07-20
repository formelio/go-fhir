package fhir

import "encoding/json"

// ProcessRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcessRequest
type ProcessRequest struct {
	Id                *string              `bson:"id" json:"id"`
	Meta              *Meta                `bson:"meta" json:"meta"`
	ImplicitRules     *string              `bson:"implicitRules" json:"implicitRules"`
	Language          *string              `bson:"language" json:"language"`
	Text              *Narrative           `bson:"text" json:"text"`
	Contained         []json.RawMessage    `bson:"contained" json:"contained"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier         `bson:"identifier" json:"identifier"`
	Status            *string              `bson:"status" json:"status"`
	Action            *ActionList          `bson:"action" json:"action"`
	Target            *Reference           `bson:"target" json:"target"`
	Created           *string              `bson:"created" json:"created"`
	Provider          *Reference           `bson:"provider" json:"provider"`
	Organization      *Reference           `bson:"organization" json:"organization"`
	Request           *Reference           `bson:"request" json:"request"`
	Response          *Reference           `bson:"response" json:"response"`
	Nullify           *bool                `bson:"nullify" json:"nullify"`
	Reference         *string              `bson:"reference" json:"reference"`
	Item              []ProcessRequestItem `bson:"item" json:"item"`
	Include           []string             `bson:"include" json:"include"`
	Exclude           []string             `bson:"exclude" json:"exclude"`
	Period            *Period              `bson:"period" json:"period"`
}
type ProcessRequestItem struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    int         `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
}
type OtherProcessRequest ProcessRequest

// MarshalJSON marshals the given ProcessRequest as JSON into a byte slice
func (r ProcessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcessRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherProcessRequest: OtherProcessRequest(r),
		ResourceType:        "ProcessRequest",
	})
}

// UnmarshalProcessRequest unmarshalls a ProcessRequest.
func UnmarshalProcessRequest(b []byte) (ProcessRequest, error) {
	var processRequest ProcessRequest
	if err := json.Unmarshal(b, &processRequest); err != nil {
		return processRequest, err
	}
	return processRequest, nil
}
