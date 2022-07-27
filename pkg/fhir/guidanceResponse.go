package fhir

import "encoding/json"

// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                    *string                `bson:"id" json:"id"`
	Meta                  *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules         *string                `bson:"implicitRules" json:"implicitRules"`
	Language              *string                `bson:"language" json:"language"`
	Text                  *Narrative             `bson:"text" json:"text"`
	RawContained          []json.RawMessage      `bson:"contained" json:"contained"`
	Contained             []IResource            `bson:"-" json:"-"`
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

// OtherGuidanceResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherGuidanceResponse GuidanceResponse

// MarshalJSON marshals the given GuidanceResponse as JSON into a byte slice
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// UnmarshalJSON unmarshals the given byte slice into GuidanceResponse
func (r *GuidanceResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherGuidanceResponse)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r GuidanceResponse) GetResourceType() ResourceType {
	return ResourceTypeGuidanceResponse
}
