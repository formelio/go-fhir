package fhir

import (
	"bytes"
	"encoding/json"
)

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                   *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               GoalStatus         `bson:"status,omitempty" json:"status,omitempty"`
	Category             []*CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Priority             *CodeableConcept   `bson:"priority,omitempty" json:"priority,omitempty"`
	Description          CodeableConcept    `bson:"description,omitempty" json:"description,omitempty"`
	Subject              *Reference         `bson:"subject,omitempty" json:"subject,omitempty"`
	StartDate            *string            `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept   `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	Target               *GoalTarget        `bson:"target,omitempty" json:"target,omitempty"`
	StatusDate           *string            `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason         *string            `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	ExpressedBy          *Reference         `bson:"expressedBy,omitempty" json:"expressedBy,omitempty"`
	Addresses            []*Reference       `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Note                 []*Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	OutcomeCode          []*CodeableConcept `bson:"outcomeCode,omitempty" json:"outcomeCode,omitempty"`
	OutcomeReference     []*Reference       `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
}
type GoalTarget struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherGoal
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
	return buffer.Bytes(), err
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
