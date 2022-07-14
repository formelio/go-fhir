package fhir

import "encoding/json"

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string            `bson:"status" json:"status"`
	Category          []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Priority          *CodeableConcept  `bson:"priority,omitempty" json:"priority,omitempty"`
	Description       CodeableConcept   `bson:"description" json:"description"`
	Target            *GoalTarget       `bson:"target,omitempty" json:"target,omitempty"`
	StatusDate        *string           `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason      *string           `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	OutcomeCode       []CodeableConcept `bson:"outcomeCode,omitempty" json:"outcomeCode,omitempty"`
	OutcomeReference  []Reference       `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
}
type GoalTarget struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure           *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
}
type OtherGoal Goal

// MarshalJSON marshals the given Goal as JSON into a byte slice
func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// UnmarshalGoal unmarshals a Goal.
func UnmarshalGoal(b []byte) (Goal, error) {
	var goal Goal
	if err := json.Unmarshal(b, &goal); err != nil {
		return goal, err
	}
	return goal, nil
}
