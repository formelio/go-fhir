package fhir

import "encoding/json"

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	RawContained      []json.RawMessage  `bson:"contained" json:"contained"`
	Contained         []IResource        `bson:"-" json:"-"`
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

// OtherCarePlan is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCarePlan CarePlan

// MarshalJSON marshals the given CarePlan as JSON into a byte slice
func (r CarePlan) MarshalJSON() ([]byte, error) {
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
		OtherCarePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherCarePlan: OtherCarePlan(r),
		ResourceType:  "CarePlan",
	})
}

// UnmarshalJSON unmarshals the given byte slice into CarePlan
func (r *CarePlan) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCarePlan)(r)); err != nil {
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
func (r CarePlan) GetResourceType() ResourceType {
	return ResourceTypeCarePlan
}
