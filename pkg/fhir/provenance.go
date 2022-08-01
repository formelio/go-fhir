package fhir

import "encoding/json"

// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            []Reference        `bson:"target,omitempty" json:"target,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Recorded          string             `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Policy            []string           `bson:"policy,omitempty" json:"policy,omitempty"`
	Location          *Reference         `bson:"location,omitempty" json:"location,omitempty"`
	Reason            []Coding           `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity          *Coding            `bson:"activity,omitempty" json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity            []ProvenanceEntity `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature         []Signature        `bson:"signature,omitempty" json:"signature,omitempty"`
}
type ProvenanceAgent struct {
	Id                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role                []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	WhoUri              *string           `bson:"whoUri,omitempty" json:"whoUri,omitempty"`
	WhoReference        *Reference        `bson:"whoReference,omitempty" json:"whoReference,omitempty"`
	OnBehalfOfUri       *string           `bson:"onBehalfOfUri,omitempty" json:"onBehalfOfUri,omitempty"`
	OnBehalfOfReference *Reference        `bson:"onBehalfOfReference,omitempty" json:"onBehalfOfReference,omitempty"`
	RelatedAgentType    *CodeableConcept  `bson:"relatedAgentType,omitempty" json:"relatedAgentType,omitempty"`
}
type ProvenanceEntity struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              ProvenanceEntityRole `bson:"role,omitempty" json:"role,omitempty"`
	WhatUri           *string              `bson:"whatUri,omitempty" json:"whatUri,omitempty"`
	WhatReference     *Reference           `bson:"whatReference,omitempty" json:"whatReference,omitempty"`
	WhatIdentifier    *Identifier          `bson:"whatIdentifier,omitempty" json:"whatIdentifier,omitempty"`
	Agent             []ProvenanceAgent    `bson:"agent,omitempty" json:"agent,omitempty"`
}

// OtherProvenance is a helper type to use the default implementations of Marshall and Unmarshal
type OtherProvenance Provenance

// MarshalJSON marshals the given Provenance as JSON into a byte slice
func (r Provenance) MarshalJSON() ([]byte, error) {
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
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Provenance
func (r *Provenance) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherProvenance)(r)); err != nil {
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
func (r Provenance) GetResourceType() ResourceType {
	return ResourceTypeProvenance
}
