package fhir

import "encoding/json"

// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           *Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent            *Reference                 `bson:"parent,omitempty" json:"parent,omitempty"`
	Status            string                     `bson:"status" json:"status"`
	Method            *CodeableConcept           `bson:"method,omitempty" json:"method,omitempty"`
	Code              *CodeableConcept           `bson:"code,omitempty" json:"code,omitempty"`
	Condition         *Reference                 `bson:"condition,omitempty" json:"condition,omitempty"`
	Basis             []Reference                `bson:"basis,omitempty" json:"basis,omitempty"`
	Prediction        []RiskAssessmentPrediction `bson:"prediction,omitempty" json:"prediction,omitempty"`
	Mitigation        *string                    `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
	Comment           *string                    `bson:"comment,omitempty" json:"comment,omitempty"`
}
type RiskAssessmentPrediction struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Outcome           CodeableConcept  `bson:"outcome" json:"outcome"`
	QualitativeRisk   *CodeableConcept `bson:"qualitativeRisk,omitempty" json:"qualitativeRisk,omitempty"`
	RelativeRisk      *string          `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty"`
	Rationale         *string          `bson:"rationale,omitempty" json:"rationale,omitempty"`
}
type OtherRiskAssessment RiskAssessment

// MarshalJSON marshals the given RiskAssessment as JSON into a byte slice
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
}

// UnmarshalRiskAssessment unmarshals a RiskAssessment.
func UnmarshalRiskAssessment(b []byte) (RiskAssessment, error) {
	var riskAssessment RiskAssessment
	if err := json.Unmarshal(b, &riskAssessment); err != nil {
		return riskAssessment, err
	}
	return riskAssessment, nil
}
