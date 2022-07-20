package fhir

import "encoding/json"

// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
	Id                   *string               `bson:"id" json:"id"`
	Meta                 *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules        *string               `bson:"implicitRules" json:"implicitRules"`
	Language             *string               `bson:"language" json:"language"`
	Text                 *Narrative            `bson:"text" json:"text"`
	Contained            []json.RawMessage     `bson:"contained" json:"contained"`
	Extension            []Extension           `bson:"extension" json:"extension"`
	ModifierExtension    []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier          `bson:"identifier" json:"identifier"`
	Status               *CareTeamStatus       `bson:"status" json:"status"`
	Category             []CodeableConcept     `bson:"category" json:"category"`
	Name                 *string               `bson:"name" json:"name"`
	Subject              *Reference            `bson:"subject" json:"subject"`
	Context              *Reference            `bson:"context" json:"context"`
	Period               *Period               `bson:"period" json:"period"`
	Participant          []CareTeamParticipant `bson:"participant" json:"participant"`
	ReasonCode           []CodeableConcept     `bson:"reasonCode" json:"reasonCode"`
	ReasonReference      []Reference           `bson:"reasonReference" json:"reasonReference"`
	ManagingOrganization []Reference           `bson:"managingOrganization" json:"managingOrganization"`
	Note                 []Annotation          `bson:"note" json:"note"`
}
type CareTeamParticipant struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Member            *Reference       `bson:"member" json:"member"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf" json:"onBehalfOf"`
	Period            *Period          `bson:"period" json:"period"`
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

// UnmarshalCareTeam unmarshalls a CareTeam.
func UnmarshalCareTeam(b []byte) (CareTeam, error) {
	var careTeam CareTeam
	if err := json.Unmarshal(b, &careTeam); err != nil {
		return careTeam, err
	}
	return careTeam, nil
}
