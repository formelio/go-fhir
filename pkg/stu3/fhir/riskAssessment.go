package fhir

import (
	"bytes"
	"encoding/json"
)

// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                    *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn               *Reference                  `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent                *Reference                  `bson:"parent,omitempty" json:"parent,omitempty"`
	Status                ObservationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Method                *CodeableConcept            `bson:"method,omitempty" json:"method,omitempty"`
	Code                  *CodeableConcept            `bson:"code,omitempty" json:"code,omitempty"`
	Subject               *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context               *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime    *string                     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                     `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Condition             *Reference                  `bson:"condition,omitempty" json:"condition,omitempty"`
	Performer             *Reference                  `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCodeableConcept *CodeableConcept            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Basis                 []*Reference                `bson:"basis,omitempty" json:"basis,omitempty"`
	Prediction            []*RiskAssessmentPrediction `bson:"prediction,omitempty" json:"prediction,omitempty"`
	Mitigation            *string                     `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
	Comment               *string                     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type RiskAssessmentPrediction struct {
	Id                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Outcome            CodeableConcept  `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ProbabilityDecimal *float64         `bson:"probabilityDecimal,omitempty" json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `bson:"probabilityRange,omitempty" json:"probabilityRange,omitempty"`
	QualitativeRisk    *CodeableConcept `bson:"qualitativeRisk,omitempty" json:"qualitativeRisk,omitempty"`
	RelativeRisk       *float64         `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty"`
	WhenPeriod         *Period          `bson:"whenPeriod,omitempty" json:"whenPeriod,omitempty"`
	WhenRange          *Range           `bson:"whenRange,omitempty" json:"whenRange,omitempty"`
	Rationale          *string          `bson:"rationale,omitempty" json:"rationale,omitempty"`
}

// OtherRiskAssessment is a helper type to use the default implementations of Marshall and Unmarshal
type OtherRiskAssessment RiskAssessment

// MarshalJSON marshals the given RiskAssessment as JSON into a byte slice
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherRiskAssessment
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into RiskAssessment
func (r *RiskAssessment) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherRiskAssessment)(r)); err != nil {
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
func (r RiskAssessment) GetResourceType() ResourceType {
	return ResourceTypeRiskAssessment
}
