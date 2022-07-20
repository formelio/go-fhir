package fhir

import "encoding/json"

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference        `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces          []Reference        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf            []Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            string             `bson:"status" json:"status"`
	Intent            string             `bson:"intent" json:"intent"`
	Category          []CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Title             *string            `bson:"title,omitempty" json:"title,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	CareTeam          []Reference        `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Addresses         []Reference        `bson:"addresses,omitempty" json:"addresses,omitempty"`
	SupportingInfo    []Reference        `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Goal              []Reference        `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity          []CarePlanActivity `bson:"activity,omitempty" json:"activity,omitempty"`
	Note              []Annotation       `bson:"note,omitempty" json:"note,omitempty"`
}
type CarePlanActivity struct {
	Id                     *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OutcomeCodeableConcept []CodeableConcept       `bson:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation            `bson:"progress,omitempty" json:"progress,omitempty"`
	Detail                 *CarePlanActivityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type CarePlanActivityDetail struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode        []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal              []Reference       `bson:"goal,omitempty" json:"goal,omitempty"`
	Status            string            `bson:"status" json:"status"`
	StatusReason      *string           `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Prohibited        *bool             `bson:"prohibited,omitempty" json:"prohibited,omitempty"`
	Location          *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	DailyAmount       *Quantity         `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity          *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
}
type OtherCarePlan CarePlan

// MarshalJSON marshals the given CarePlan as JSON into a byte slice
func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCarePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherCarePlan: OtherCarePlan(r),
		ResourceType:  "CarePlan",
	})
}

// UnmarshalCarePlan unmarshals a CarePlan.
func UnmarshalCarePlan(b []byte) (CarePlan, error) {
	var carePlan CarePlan
	if err := json.Unmarshal(b, &carePlan); err != nil {
		return carePlan, err
	}
	return carePlan, nil
}
