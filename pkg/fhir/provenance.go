package fhir

import "encoding/json"

// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            []Reference        `bson:"target" json:"target"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Recorded          string             `bson:"recorded" json:"recorded"`
	Policy            []string           `bson:"policy,omitempty" json:"policy,omitempty"`
	Location          *Reference         `bson:"location,omitempty" json:"location,omitempty"`
	Reason            []Coding           `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity          *Coding            `bson:"activity,omitempty" json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `bson:"agent" json:"agent"`
	Entity            []ProvenanceEntity `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature         []Signature        `bson:"signature,omitempty" json:"signature,omitempty"`
}
type ProvenanceAgent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	RelatedAgentType  *CodeableConcept  `bson:"relatedAgentType,omitempty" json:"relatedAgentType,omitempty"`
}
type ProvenanceEntity struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              string            `bson:"role" json:"role"`
	Agent             []ProvenanceAgent `bson:"agent,omitempty" json:"agent,omitempty"`
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

// UnmarshalProvenance unmarshals a Provenance.
func UnmarshalProvenance(b []byte) (Provenance, error) {
	var provenance Provenance
	if err := json.Unmarshal(b, &provenance); err != nil {
		return provenance, err
	}
	return provenance, nil
}
