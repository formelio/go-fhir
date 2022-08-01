package fhir

import "encoding/json"

// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                    *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RequestId             *string                `bson:"requestId,omitempty" json:"requestId,omitempty"`
	Identifier            *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Module                Reference              `bson:"module,omitempty" json:"module,omitempty"`
	Status                GuidanceResponseStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject               *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	Context               *Reference             `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime    *string                `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	Performer             *Reference             `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCodeableConcept *CodeableConcept       `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference             `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	EvaluationMessage     []Reference            `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty"`
	OutputParameters      *Reference             `bson:"outputParameters,omitempty" json:"outputParameters,omitempty"`
	Result                *Reference             `bson:"result,omitempty" json:"result,omitempty"`
	DataRequirement       []DataRequirement      `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
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
