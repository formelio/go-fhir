package fhir

import "encoding/json"

// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                    *string                `bson:"id" json:"id"`
	Meta                  *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules         *string                `bson:"implicitRules" json:"implicitRules"`
	Language              *string                `bson:"language" json:"language"`
	Text                  *Narrative             `bson:"text" json:"text"`
	Contained             []json.RawMessage      `bson:"contained" json:"contained"`
	Extension             []Extension            `bson:"extension" json:"extension"`
	ModifierExtension     []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	RequestId             *string                `bson:"requestId" json:"requestId"`
	Identifier            *Identifier            `bson:"identifier" json:"identifier"`
	Module                Reference              `bson:"module,omitempty" json:"module,omitempty"`
	Status                GuidanceResponseStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject               *Reference             `bson:"subject" json:"subject"`
	Context               *Reference             `bson:"context" json:"context"`
	OccurrenceDateTime    *string                `bson:"occurrenceDateTime" json:"occurrenceDateTime"`
	Performer             *Reference             `bson:"performer" json:"performer"`
	ReasonCodeableConcept *CodeableConcept       `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference             `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation           `bson:"note" json:"note"`
	EvaluationMessage     []Reference            `bson:"evaluationMessage" json:"evaluationMessage"`
	OutputParameters      *Reference             `bson:"outputParameters" json:"outputParameters"`
	Result                *Reference             `bson:"result" json:"result"`
	DataRequirement       []DataRequirement      `bson:"dataRequirement" json:"dataRequirement"`
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

// UnmarshalGuidanceResponse unmarshalls a GuidanceResponse.
func UnmarshalGuidanceResponse(b []byte) (GuidanceResponse, error) {
	var guidanceResponse GuidanceResponse
	if err := json.Unmarshal(b, &guidanceResponse); err != nil {
		return guidanceResponse, err
	}
	return guidanceResponse, nil
}
