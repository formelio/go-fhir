package fhir

import "encoding/json"

// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                    *string                    `bson:"id" json:"id"`
	Meta                  *Meta                      `bson:"meta" json:"meta"`
	ImplicitRules         *string                    `bson:"implicitRules" json:"implicitRules"`
	Language              *string                    `bson:"language" json:"language"`
	Text                  *Narrative                 `bson:"text" json:"text"`
	RawContained          []json.RawMessage          `bson:"contained" json:"contained"`
	Contained             []IResource                `bson:"-" json:"-"`
	Extension             []Extension                `bson:"extension" json:"extension"`
	ModifierExtension     []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            *Identifier                `bson:"identifier" json:"identifier"`
	BasedOn               *Reference                 `bson:"basedOn" json:"basedOn"`
	Parent                *Reference                 `bson:"parent" json:"parent"`
	Status                ObservationStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Method                *CodeableConcept           `bson:"method" json:"method"`
	Code                  *CodeableConcept           `bson:"code" json:"code"`
	Subject               *Reference                 `bson:"subject" json:"subject"`
	Context               *Reference                 `bson:"context" json:"context"`
	OccurrenceDateTime    *string                    `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                    `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Condition             *Reference                 `bson:"condition" json:"condition"`
	Performer             *Reference                 `bson:"performer" json:"performer"`
	ReasonCodeableConcept *CodeableConcept           `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                 `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Basis                 []Reference                `bson:"basis" json:"basis"`
	Prediction            []RiskAssessmentPrediction `bson:"prediction" json:"prediction"`
	Mitigation            *string                    `bson:"mitigation" json:"mitigation"`
	Comment               *string                    `bson:"comment" json:"comment"`
}
type RiskAssessmentPrediction struct {
	Id                 *string          `bson:"id" json:"id"`
	Extension          []Extension      `bson:"extension" json:"extension"`
	ModifierExtension  []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Outcome            CodeableConcept  `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ProbabilityDecimal *float64         `bson:"probabilityDecimal,omitempty" json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `bson:"probabilityRange,omitempty" json:"probabilityRange,omitempty"`
	QualitativeRisk    *CodeableConcept `bson:"qualitativeRisk" json:"qualitativeRisk"`
	RelativeRisk       *float64         `bson:"relativeRisk" json:"relativeRisk"`
	WhenPeriod         *Period          `bson:"whenPeriod,omitempty" json:"whenPeriod,omitempty"`
	WhenRange          *Range           `bson:"whenRange,omitempty" json:"whenRange,omitempty"`
	Rationale          *string          `bson:"rationale" json:"rationale"`
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
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
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
