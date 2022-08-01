package fhir

import "encoding/json"

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition        []Reference        `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn           []Reference        `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces          []Reference        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf            []Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            CarePlanStatus     `bson:"status,omitempty" json:"status,omitempty"`
	Intent            CarePlanIntent     `bson:"intent,omitempty" json:"intent,omitempty"`
	Category          []CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Title             *string            `bson:"title,omitempty" json:"title,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	Subject           Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Context           *Reference         `bson:"context,omitempty" json:"context,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Author            []Reference        `bson:"author,omitempty" json:"author,omitempty"`
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
	Reference              *Reference              `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type CarePlanActivityDetail struct {
	Id                     *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category               *CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	Definition             *Reference             `bson:"definition,omitempty" json:"definition,omitempty"`
	Code                   *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference            `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                 CarePlanActivityStatus `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason           *string                `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Prohibited             *bool                  `bson:"prohibited,omitempty" json:"prohibited,omitempty"`
	ScheduledTiming        *Timing                `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period                `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        *string                `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference             `bson:"location,omitempty" json:"location,omitempty"`
	Performer              []Reference            `bson:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept       `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference             `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity              `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description            *string                `bson:"description,omitempty" json:"description,omitempty"`
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
