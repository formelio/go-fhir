package fhir

import "encoding/json"

// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	RawContained      []json.RawMessage  `bson:"contained" json:"contained"`
	Contained         []IResource        `bson:"-" json:"-"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Target            []Reference        `bson:"target,omitempty" json:"target,omitempty"`
	Period            *Period            `bson:"period" json:"period"`
	Recorded          string             `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Policy            []string           `bson:"policy" json:"policy"`
	Location          *Reference         `bson:"location" json:"location"`
	Reason            []Coding           `bson:"reason" json:"reason"`
	Activity          *Coding            `bson:"activity" json:"activity"`
	Agent             []ProvenanceAgent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity            []ProvenanceEntity `bson:"entity" json:"entity"`
	Signature         []Signature        `bson:"signature" json:"signature"`
}
type ProvenanceAgent struct {
	Id                  *string           `bson:"id" json:"id"`
	Extension           []Extension       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Role                []CodeableConcept `bson:"role" json:"role"`
	WhoUri              *string           `bson:"whoUri,omitempty" json:"whoUri,omitempty"`
	WhoReference        *Reference        `bson:"whoReference,omitempty" json:"whoReference,omitempty"`
	OnBehalfOfUri       *string           `bson:"onBehalfOfUri,omitempty" json:"onBehalfOfUri,omitempty"`
	OnBehalfOfReference *Reference        `bson:"onBehalfOfReference,omitempty" json:"onBehalfOfReference,omitempty"`
	RelatedAgentType    *CodeableConcept  `bson:"relatedAgentType" json:"relatedAgentType"`
}
type ProvenanceEntity struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Role              ProvenanceEntityRole `bson:"role,omitempty" json:"role,omitempty"`
	WhatUri           *string              `bson:"whatUri,omitempty" json:"whatUri,omitempty"`
	WhatReference     *Reference           `bson:"whatReference,omitempty" json:"whatReference,omitempty"`
	WhatIdentifier    *Identifier          `bson:"whatIdentifier,omitempty" json:"whatIdentifier,omitempty"`
	Agent             []ProvenanceAgent    `bson:"agent" json:"agent"`
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
