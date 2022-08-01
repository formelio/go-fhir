package fhir

import (
	"bytes"
	"encoding/json"
)

// EpisodeOfCare is documented here http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
	Id                   *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage             `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource                   `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               EpisodeOfCareStatus           `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory        []*EpisodeOfCareStatusHistory `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Type                 []*CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Diagnosis            []*EpisodeOfCareDiagnosis     `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Patient              Reference                     `bson:"patient,omitempty" json:"patient,omitempty"`
	ManagingOrganization *Reference                    `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Period               *Period                       `bson:"period,omitempty" json:"period,omitempty"`
	ReferralRequest      []*Reference                  `bson:"referralRequest,omitempty" json:"referralRequest,omitempty"`
	CareManager          *Reference                    `bson:"careManager,omitempty" json:"careManager,omitempty"`
	Team                 []*Reference                  `bson:"team,omitempty" json:"team,omitempty"`
	Account              []*Reference                  `bson:"account,omitempty" json:"account,omitempty"`
}
type EpisodeOfCareStatusHistory struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            EpisodeOfCareStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            Period              `bson:"period,omitempty" json:"period,omitempty"`
}
type EpisodeOfCareDiagnosis struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Condition         Reference        `bson:"condition,omitempty" json:"condition,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Rank              *int             `bson:"rank,omitempty" json:"rank,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherEpisodeOfCare
	}{
		OtherEpisodeOfCare: OtherEpisodeOfCare(r),
		ResourceType:       "EpisodeOfCare",
	})
	return buffer.Bytes(), err
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
