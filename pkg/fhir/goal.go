package fhir

import "encoding/json"

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	RawContained         []json.RawMessage `bson:"contained" json:"contained"`
	Contained            []IResource       `bson:"-" json:"-"`
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

// OtherGoal is a helper type to use the default implementations of Marshall and Unmarshal
type OtherGoal Goal

// MarshalJSON marshals the given Goal as JSON into a byte slice
func (r Goal) MarshalJSON() ([]byte, error) {
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
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Goal
func (r *Goal) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherGoal)(r)); err != nil {
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
func (r Goal) GetResourceType() ResourceType {
	return ResourceTypeGoal
}
