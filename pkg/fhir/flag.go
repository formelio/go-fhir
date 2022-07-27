package fhir

import "encoding/json"

// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
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
	Status            FlagStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept  `bson:"category" json:"category"`
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           Reference         `bson:"subject,omitempty" json:"subject,omitempty"`
	Period            *Period           `bson:"period" json:"period"`
	Encounter         *Reference        `bson:"encounter" json:"encounter"`
	Author            *Reference        `bson:"author" json:"author"`
}

// OtherFlag is a helper type to use the default implementations of Marshall and Unmarshal
type OtherFlag Flag

// MarshalJSON marshals the given Flag as JSON into a byte slice
func (r Flag) MarshalJSON() ([]byte, error) {
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
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Flag
func (r *Flag) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherFlag)(r)); err != nil {
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
func (r Flag) GetResourceType() ResourceType {
	return ResourceTypeFlag
}
