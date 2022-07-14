package fhir

import "encoding/json"

// Condition is documented here http://hl7.org/fhir/StructureDefinition/Condition
type Condition struct {
	Id                 *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *string             `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus *string             `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Category           []CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Severity           *CodeableConcept    `bson:"severity,omitempty" json:"severity,omitempty"`
	Code               *CodeableConcept    `bson:"code,omitempty" json:"code,omitempty"`
	BodySite           []CodeableConcept   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	AssertedDate       *string             `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Stage              *ConditionStage     `bson:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []ConditionEvidence `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Note               []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
}
type ConditionStage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
}
type ConditionEvidence struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type OtherCondition Condition

// MarshalJSON marshals the given Condition as JSON into a byte slice
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

// UnmarshalCondition unmarshals a Condition.
func UnmarshalCondition(b []byte) (Condition, error) {
	var condition Condition
	if err := json.Unmarshal(b, &condition); err != nil {
		return condition, err
	}
	return condition, nil
}
