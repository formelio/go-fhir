package fhir

import "encoding/json"

// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string               `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Name                 *string               `bson:"name,omitempty" json:"name,omitempty"`
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
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	Period            *Period          `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherCareTeam CareTeam

// MarshalJSON marshals the given CareTeam as JSON into a byte slice
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// UnmarshalCareTeam unmarshals a CareTeam.
func UnmarshalCareTeam(b []byte) (CareTeam, error) {
	var careTeam CareTeam
	if err := json.Unmarshal(b, &careTeam); err != nil {
		return careTeam, err
	}
	return careTeam, nil
}
