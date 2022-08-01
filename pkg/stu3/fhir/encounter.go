package fhir

import (
	"bytes"
	"encoding/json"
)

// Encounter is documented here http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource               `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            EncounterStatus           `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory     []*EncounterStatusHistory `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class             *Coding                   `bson:"class,omitempty" json:"class,omitempty"`
	ClassHistory      []*EncounterClassHistory  `bson:"classHistory,omitempty" json:"classHistory,omitempty"`
	Type              []*CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Priority          *CodeableConcept          `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject           *Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	EpisodeOfCare     []*Reference              `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferral  []*Reference              `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant       []*EncounterParticipant   `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment       *Reference                `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period            *Period                   `bson:"period,omitempty" json:"period,omitempty"`
	Length            *Duration                 `bson:"length,omitempty" json:"length,omitempty"`
	Reason            []*CodeableConcept        `bson:"reason,omitempty" json:"reason,omitempty"`
	Diagnosis         []*EncounterDiagnosis     `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account           []*Reference              `bson:"account,omitempty" json:"account,omitempty"`
	Hospitalization   *EncounterHospitalization `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location          []*EncounterLocation      `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider   *Reference                `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf            *Reference                `bson:"partOf,omitempty" json:"partOf,omitempty"`
}
type EncounterStatusHistory struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            EncounterStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            Period          `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterClassHistory struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Class             Coding       `bson:"class,omitempty" json:"class,omitempty"`
	Period            Period       `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterParticipant struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []*CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Individual        *Reference         `bson:"individual,omitempty" json:"individual,omitempty"`
}
type EncounterDiagnosis struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Condition         Reference        `bson:"condition,omitempty" json:"condition,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Rank              *int             `bson:"rank,omitempty" json:"rank,omitempty"`
}
type EncounterHospitalization struct {
	Id                     *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier        `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference         `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept   `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept   `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []*CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []*CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []*CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference         `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept   `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}
type EncounterLocation struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Location          Reference                `bson:"location,omitempty" json:"location,omitempty"`
	Status            *EncounterLocationStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            *Period                  `bson:"period,omitempty" json:"period,omitempty"`
}

// OtherEncounter is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEncounter Encounter

// MarshalJSON marshals the given Encounter as JSON into a byte slice
func (r Encounter) MarshalJSON() ([]byte, error) {
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
		OtherEncounter
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Encounter
func (r *Encounter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEncounter)(r)); err != nil {
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
func (r Encounter) GetResourceType() ResourceType {
	return ResourceTypeEncounter
}
