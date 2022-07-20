package fhir

import "encoding/json"

// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	Contained         []json.RawMessage  `bson:"contained" json:"contained"`
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
type OtherProvenance Provenance

// MarshalJSON marshals the given Provenance as JSON into a byte slice
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}

// UnmarshalProvenance unmarshalls a Provenance.
func UnmarshalProvenance(b []byte) (Provenance, error) {
	var provenance Provenance
	if err := json.Unmarshal(b, &provenance); err != nil {
		return provenance, err
	}
	return provenance, nil
}
