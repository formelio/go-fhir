package fhir

import "encoding/json"

// EpisodeOfCare is documented here http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
	Id                   *string                      `bson:"id" json:"id"`
	Meta                 *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules        *string                      `bson:"implicitRules" json:"implicitRules"`
	Language             *string                      `bson:"language" json:"language"`
	Text                 *Narrative                   `bson:"text" json:"text"`
	RawContained         []json.RawMessage            `bson:"contained" json:"contained"`
	Contained            []IResource                  `bson:"-" json:"-"`
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

// OtherEpisodeOfCare is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEpisodeOfCare EpisodeOfCare

// MarshalJSON marshals the given EpisodeOfCare as JSON into a byte slice
func (r EpisodeOfCare) MarshalJSON() ([]byte, error) {
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
		OtherEpisodeOfCare
		ResourceType string `json:"resourceType"`
	}{
		OtherEpisodeOfCare: OtherEpisodeOfCare(r),
		ResourceType:       "EpisodeOfCare",
	})
}

// UnmarshalJSON unmarshals the given byte slice into EpisodeOfCare
func (r *EpisodeOfCare) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEpisodeOfCare)(r)); err != nil {
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
func (r EpisodeOfCare) GetResourceType() ResourceType {
	return ResourceTypeEpisodeOfCare
}
