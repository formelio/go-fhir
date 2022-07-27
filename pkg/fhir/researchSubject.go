package fhir

import "encoding/json"

// ResearchSubject is documented here http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier           `bson:"identifier" json:"identifier"`
	Status            ResearchSubjectStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            *Period               `bson:"period" json:"period"`
	Study             Reference             `bson:"study,omitempty" json:"study,omitempty"`
	Individual        Reference             `bson:"individual,omitempty" json:"individual,omitempty"`
	AssignedArm       *string               `bson:"assignedArm" json:"assignedArm"`
	ActualArm         *string               `bson:"actualArm" json:"actualArm"`
	Consent           *Reference            `bson:"consent" json:"consent"`
}

// OtherResearchSubject is a helper type to use the default implementations of Marshall and Unmarshal
type OtherResearchSubject ResearchSubject

// MarshalJSON marshals the given ResearchSubject as JSON into a byte slice
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
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
		OtherResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchSubject: OtherResearchSubject(r),
		ResourceType:         "ResearchSubject",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ResearchSubject
func (r *ResearchSubject) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherResearchSubject)(r)); err != nil {
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
func (r ResearchSubject) GetResourceType() ResourceType {
	return ResourceTypeResearchSubject
}
