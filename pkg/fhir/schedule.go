package fhir

import "encoding/json"

// Schedule is documented here http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Active            *bool             `bson:"active" json:"active"`
	ServiceCategory   *CodeableConcept  `bson:"serviceCategory" json:"serviceCategory"`
	ServiceType       []CodeableConcept `bson:"serviceType" json:"serviceType"`
	Specialty         []CodeableConcept `bson:"specialty" json:"specialty"`
	Actor             []Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	PlanningHorizon   *Period           `bson:"planningHorizon" json:"planningHorizon"`
	Comment           *string           `bson:"comment" json:"comment"`
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

// UnmarshalSchedule unmarshalls a Schedule.
func UnmarshalSchedule(b []byte) (Schedule, error) {
	var schedule Schedule
	if err := json.Unmarshal(b, &schedule); err != nil {
		return schedule, err
	}
	return schedule, nil
}
