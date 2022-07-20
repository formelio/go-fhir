package fhir

import "encoding/json"

// Schedule is documented here http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool             `bson:"active,omitempty" json:"active,omitempty"`
	ServiceCategory   *CodeableConcept  `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	PlanningHorizon   *Period           `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty"`
	Comment           *string           `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherSchedule Schedule

// MarshalJSON marshals the given Schedule as JSON into a byte slice
func (r Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSchedule
		ResourceType string `json:"resourceType"`
	}{
		OtherSchedule: OtherSchedule(r),
		ResourceType:  "Schedule",
	})
}

// UnmarshalSchedule unmarshals a Schedule.
func UnmarshalSchedule(b []byte) (Schedule, error) {
	var schedule Schedule
	if err := json.Unmarshal(b, &schedule); err != nil {
		return schedule, err
	}
	return schedule, nil
}
