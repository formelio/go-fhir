package fhir

import "encoding/json"

// EpisodeOfCare is documented here http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
	Id                   *string                      `bson:"id" json:"id"`
	Meta                 *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules        *string                      `bson:"implicitRules" json:"implicitRules"`
	Language             *string                      `bson:"language" json:"language"`
	Text                 *Narrative                   `bson:"text" json:"text"`
	Contained            []json.RawMessage            `bson:"contained" json:"contained"`
	Extension            []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier                 `bson:"identifier" json:"identifier"`
	Status               EpisodeOfCareStatus          `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory        []EpisodeOfCareStatusHistory `bson:"statusHistory" json:"statusHistory"`
	Type                 []CodeableConcept            `bson:"type" json:"type"`
	Diagnosis            []EpisodeOfCareDiagnosis     `bson:"diagnosis" json:"diagnosis"`
	Patient              Reference                    `bson:"patient,omitempty" json:"patient,omitempty"`
	ManagingOrganization *Reference                   `bson:"managingOrganization" json:"managingOrganization"`
	Period               *Period                      `bson:"period" json:"period"`
	ReferralRequest      []Reference                  `bson:"referralRequest" json:"referralRequest"`
	CareManager          *Reference                   `bson:"careManager" json:"careManager"`
	Team                 []Reference                  `bson:"team" json:"team"`
	Account              []Reference                  `bson:"account" json:"account"`
}
type EpisodeOfCareStatusHistory struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Status            EpisodeOfCareStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            Period              `bson:"period,omitempty" json:"period,omitempty"`
}
type EpisodeOfCareDiagnosis struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Condition         Reference        `bson:"condition,omitempty" json:"condition,omitempty"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Rank              *int             `bson:"rank" json:"rank"`
}
type OtherEpisodeOfCare EpisodeOfCare

// MarshalJSON marshals the given EpisodeOfCare as JSON into a byte slice
func (r EpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEpisodeOfCare
		ResourceType string `json:"resourceType"`
	}{
		OtherEpisodeOfCare: OtherEpisodeOfCare(r),
		ResourceType:       "EpisodeOfCare",
	})
}

// UnmarshalEpisodeOfCare unmarshalls a EpisodeOfCare.
func UnmarshalEpisodeOfCare(b []byte) (EpisodeOfCare, error) {
	var episodeOfCare EpisodeOfCare
	if err := json.Unmarshal(b, &episodeOfCare); err != nil {
		return episodeOfCare, err
	}
	return episodeOfCare, nil
}
