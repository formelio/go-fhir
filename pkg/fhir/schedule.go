package fhir

import "encoding/json"

// Schedule is documented here http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
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

// OtherSchedule is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSchedule Schedule

// MarshalJSON marshals the given Schedule as JSON into a byte slice
func (r Schedule) MarshalJSON() ([]byte, error) {
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
		OtherSchedule
		ResourceType string `json:"resourceType"`
	}{
		OtherSchedule: OtherSchedule(r),
		ResourceType:  "Schedule",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Schedule
func (r *Schedule) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSchedule)(r)); err != nil {
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
func (r Schedule) GetResourceType() ResourceType {
	return ResourceTypeSchedule
}
