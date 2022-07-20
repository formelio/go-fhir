package fhir

import "encoding/json"

// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                 *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RequestId          *string           `bson:"requestId,omitempty" json:"requestId,omitempty"`
	Identifier         *Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Module             Reference         `bson:"module" json:"module"`
	Status             string            `bson:"status" json:"status"`
	OccurrenceDateTime *string           `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	Performer          *Reference        `bson:"performer,omitempty" json:"performer,omitempty"`
	Note               []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	EvaluationMessage  []Reference       `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty"`
	OutputParameters   *Reference        `bson:"outputParameters,omitempty" json:"outputParameters,omitempty"`
	DataRequirement    []DataRequirement `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
}
type OtherGuidanceResponse GuidanceResponse

// MarshalJSON marshals the given GuidanceResponse as JSON into a byte slice
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// UnmarshalGuidanceResponse unmarshals a GuidanceResponse.
func UnmarshalGuidanceResponse(b []byte) (GuidanceResponse, error) {
	var guidanceResponse GuidanceResponse
	if err := json.Unmarshal(b, &guidanceResponse); err != nil {
		return guidanceResponse, err
	}
	return guidanceResponse, nil
}
