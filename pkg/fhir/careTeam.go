package fhir

import "encoding/json"

// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *CareTeamStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Name                 *string               `bson:"name,omitempty" json:"name,omitempty"`
	Subject              *Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	Context              *Reference            `bson:"context,omitempty" json:"context,omitempty"`
	Period               *Period               `bson:"period,omitempty" json:"period,omitempty"`
	Participant          []CareTeamParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	ReasonCode           []CodeableConcept     `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference      []Reference           `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ManagingOrganization []Reference           `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Note                 []Annotation          `bson:"note,omitempty" json:"note,omitempty"`
}
type CareTeamParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member            *Reference       `bson:"member,omitempty" json:"member,omitempty"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	Period            *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

// OtherCareTeam is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCareTeam CareTeam

// MarshalJSON marshals the given CareTeam as JSON into a byte slice
func (r CareTeam) MarshalJSON() ([]byte, error) {
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
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// UnmarshalJSON unmarshals the given byte slice into CareTeam
func (r *CareTeam) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCareTeam)(r)); err != nil {
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
func (r CareTeam) GetResourceType() ResourceType {
	return ResourceTypeCareTeam
}
