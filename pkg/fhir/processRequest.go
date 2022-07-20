package fhir

import "encoding/json"

// ProcessRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcessRequest
type ProcessRequest struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string              `bson:"status,omitempty" json:"status,omitempty"`
	Action            *string              `bson:"action,omitempty" json:"action,omitempty"`
	Target            *Reference           `bson:"target,omitempty" json:"target,omitempty"`
	Created           *string              `bson:"created,omitempty" json:"created,omitempty"`
	Provider          *Reference           `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization      *Reference           `bson:"organization,omitempty" json:"organization,omitempty"`
	Request           *Reference           `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference           `bson:"response,omitempty" json:"response,omitempty"`
	Nullify           *bool                `bson:"nullify,omitempty" json:"nullify,omitempty"`
	Reference         *string              `bson:"reference,omitempty" json:"reference,omitempty"`
	Item              []ProcessRequestItem `bson:"item,omitempty" json:"item,omitempty"`
	Include           []string             `bson:"include,omitempty" json:"include,omitempty"`
	Exclude           []string             `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period            *Period              `bson:"period,omitempty" json:"period,omitempty"`
}
type ProcessRequestItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int         `bson:"sequenceLinkId" json:"sequenceLinkId"`
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

// UnmarshalProcessRequest unmarshals a ProcessRequest.
func UnmarshalProcessRequest(b []byte) (ProcessRequest, error) {
	var processRequest ProcessRequest
	if err := json.Unmarshal(b, &processRequest); err != nil {
		return processRequest, err
	}
	return processRequest, nil
}
