package fhir

import "encoding/json"

// Encounter is documented here http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                    `bson:"status" json:"status"`
	StatusHistory     []EncounterStatusHistory  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class             *Coding                   `bson:"class,omitempty" json:"class,omitempty"`
	ClassHistory      []EncounterClassHistory   `bson:"classHistory,omitempty" json:"classHistory,omitempty"`
	Type              []CodeableConcept         `bson:"type,omitempty" json:"type,omitempty"`
	Priority          *CodeableConcept          `bson:"priority,omitempty" json:"priority,omitempty"`
	EpisodeOfCare     []Reference               `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferral  []Reference               `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant       []EncounterParticipant    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment       *Reference                `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period            *Period                   `bson:"period,omitempty" json:"period,omitempty"`
	Length            *Duration                 `bson:"length,omitempty" json:"length,omitempty"`
	Reason            []CodeableConcept         `bson:"reason,omitempty" json:"reason,omitempty"`
	Diagnosis         []EncounterDiagnosis      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account           []Reference               `bson:"account,omitempty" json:"account,omitempty"`
	Hospitalization   *EncounterHospitalization `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location          []EncounterLocation       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider   *Reference                `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf            *Reference                `bson:"partOf,omitempty" json:"partOf,omitempty"`
}
type EncounterStatusHistory struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string      `bson:"status" json:"status"`
	Period            Period      `bson:"period" json:"period"`
}
type EncounterClassHistory struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Class             Coding      `bson:"class" json:"class"`
	Period            Period      `bson:"period" json:"period"`
}
type EncounterParticipant struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterDiagnosis struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Rank              *int             `bson:"rank,omitempty" json:"rank,omitempty"`
}
type EncounterHospitalization struct {
	Id                     *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}
type EncounterLocation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Location          Reference   `bson:"location" json:"location"`
	Status            *string     `bson:"status,omitempty" json:"status,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherEncounter Encounter

// MarshalJSON marshals the given Encounter as JSON into a byte slice
func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounter
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
}

// UnmarshalEncounter unmarshals a Encounter.
func UnmarshalEncounter(b []byte) (Encounter, error) {
	var encounter Encounter
	if err := json.Unmarshal(b, &encounter); err != nil {
		return encounter, err
	}
	return encounter, nil
}
