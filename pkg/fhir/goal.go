package fhir

import "encoding/json"

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	Contained            []json.RawMessage `bson:"contained" json:"contained"`
	Extension            []Extension       `bson:"extension" json:"extension"`
	ModifierExtension    []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier      `bson:"identifier" json:"identifier"`
	Status               GoalStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept `bson:"category" json:"category"`
	Priority             *CodeableConcept  `bson:"priority" json:"priority"`
	Description          CodeableConcept   `bson:"description,omitempty" json:"description,omitempty"`
	Subject              *Reference        `bson:"subject" json:"subject"`
	StartDate            *string           `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept  `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	Target               *GoalTarget       `bson:"target" json:"target"`
	StatusDate           *string           `bson:"statusDate" json:"statusDate"`
	StatusReason         *string           `bson:"statusReason" json:"statusReason"`
	ExpressedBy          *Reference        `bson:"expressedBy" json:"expressedBy"`
	Addresses            []Reference       `bson:"addresses" json:"addresses"`
	Note                 []Annotation      `bson:"note" json:"note"`
	OutcomeCode          []CodeableConcept `bson:"outcomeCode" json:"outcomeCode"`
	OutcomeReference     []Reference       `bson:"outcomeReference" json:"outcomeReference"`
}
type GoalTarget struct {
	Id                    *string          `bson:"id" json:"id"`
	Extension             []Extension      `bson:"extension" json:"extension"`
	ModifierExtension     []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Measure               *CodeableConcept `bson:"measure" json:"measure"`
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty"`
	DueDate               *string          `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	DueDuration           *Duration        `bson:"dueDuration,omitempty" json:"dueDuration,omitempty"`
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

// UnmarshalGoal unmarshalls a Goal.
func UnmarshalGoal(b []byte) (Goal, error) {
	var goal Goal
	if err := json.Unmarshal(b, &goal); err != nil {
		return goal, err
	}
	return goal, nil
}
