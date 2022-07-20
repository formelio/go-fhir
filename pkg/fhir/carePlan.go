package fhir

import "encoding/json"

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	Contained         []json.RawMessage  `bson:"contained" json:"contained"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier       `bson:"identifier" json:"identifier"`
	Definition        []Reference        `bson:"definition" json:"definition"`
	BasedOn           []Reference        `bson:"basedOn" json:"basedOn"`
	Replaces          []Reference        `bson:"replaces" json:"replaces"`
	PartOf            []Reference        `bson:"partOf" json:"partOf"`
	Status            CarePlanStatus     `bson:"status,omitempty" json:"status,omitempty"`
	Intent            CarePlanIntent     `bson:"intent,omitempty" json:"intent,omitempty"`
	Category          []CodeableConcept  `bson:"category" json:"category"`
	Title             *string            `bson:"title" json:"title"`
	Description       *string            `bson:"description" json:"description"`
	Subject           Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Context           *Reference         `bson:"context" json:"context"`
	Period            *Period            `bson:"period" json:"period"`
	Author            []Reference        `bson:"author" json:"author"`
	CareTeam          []Reference        `bson:"careTeam" json:"careTeam"`
	Addresses         []Reference        `bson:"addresses" json:"addresses"`
	SupportingInfo    []Reference        `bson:"supportingInfo" json:"supportingInfo"`
	Goal              []Reference        `bson:"goal" json:"goal"`
	Activity          []CarePlanActivity `bson:"activity" json:"activity"`
	Note              []Annotation       `bson:"note" json:"note"`
}
type CarePlanActivity struct {
	Id                     *string                 `bson:"id" json:"id"`
	Extension              []Extension             `bson:"extension" json:"extension"`
	ModifierExtension      []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	OutcomeCodeableConcept []CodeableConcept       `bson:"outcomeCodeableConcept" json:"outcomeCodeableConcept"`
	OutcomeReference       []Reference             `bson:"outcomeReference" json:"outcomeReference"`
	Progress               []Annotation            `bson:"progress" json:"progress"`
	Reference              *Reference              `bson:"reference" json:"reference"`
	Detail                 *CarePlanActivityDetail `bson:"detail" json:"detail"`
}
type CarePlanActivityDetail struct {
	Id                     *string                `bson:"id" json:"id"`
	Extension              []Extension            `bson:"extension" json:"extension"`
	ModifierExtension      []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Category               *CodeableConcept       `bson:"category" json:"category"`
	Definition             *Reference             `bson:"definition" json:"definition"`
	Code                   *CodeableConcept       `bson:"code" json:"code"`
	ReasonCode             []CodeableConcept      `bson:"reasonCode" json:"reasonCode"`
	ReasonReference        []Reference            `bson:"reasonReference" json:"reasonReference"`
	Goal                   []Reference            `bson:"goal" json:"goal"`
	Status                 CarePlanActivityStatus `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason           *string                `bson:"statusReason" json:"statusReason"`
	Prohibited             *bool                  `bson:"prohibited" json:"prohibited"`
	ScheduledTiming        *Timing                `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period                `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        *string                `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference             `bson:"location" json:"location"`
	Performer              []Reference            `bson:"performer" json:"performer"`
	ProductCodeableConcept *CodeableConcept       `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference             `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity              `bson:"dailyAmount" json:"dailyAmount"`
	Quantity               *Quantity              `bson:"quantity" json:"quantity"`
	Description            *string                `bson:"description" json:"description"`
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

// UnmarshalCarePlan unmarshalls a CarePlan.
func UnmarshalCarePlan(b []byte) (CarePlan, error) {
	var carePlan CarePlan
	if err := json.Unmarshal(b, &carePlan); err != nil {
		return carePlan, err
	}
	return carePlan, nil
}
