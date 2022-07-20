package fhir

import "encoding/json"

// Encounter is documented here http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
	Id                *string                   `bson:"id" json:"id"`
	Meta              *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules     *string                   `bson:"implicitRules" json:"implicitRules"`
	Language          *string                   `bson:"language" json:"language"`
	Text              *Narrative                `bson:"text" json:"text"`
	Contained         []json.RawMessage         `bson:"contained" json:"contained"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier              `bson:"identifier" json:"identifier"`
	Status            EncounterStatus           `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory     []EncounterStatusHistory  `bson:"statusHistory" json:"statusHistory"`
	Class             *Coding                   `bson:"class" json:"class"`
	ClassHistory      []EncounterClassHistory   `bson:"classHistory" json:"classHistory"`
	Type              []CodeableConcept         `bson:"type" json:"type"`
	Priority          *CodeableConcept          `bson:"priority" json:"priority"`
	Subject           *Reference                `bson:"subject" json:"subject"`
	EpisodeOfCare     []Reference               `bson:"episodeOfCare" json:"episodeOfCare"`
	IncomingReferral  []Reference               `bson:"incomingReferral" json:"incomingReferral"`
	Participant       []EncounterParticipant    `bson:"participant" json:"participant"`
	Appointment       *Reference                `bson:"appointment" json:"appointment"`
	Period            *Period                   `bson:"period" json:"period"`
	Length            *Duration                 `bson:"length" json:"length"`
	Reason            []CodeableConcept         `bson:"reason" json:"reason"`
	Diagnosis         []EncounterDiagnosis      `bson:"diagnosis" json:"diagnosis"`
	Account           []Reference               `bson:"account" json:"account"`
	Hospitalization   *EncounterHospitalization `bson:"hospitalization" json:"hospitalization"`
	Location          []EncounterLocation       `bson:"location" json:"location"`
	ServiceProvider   *Reference                `bson:"serviceProvider" json:"serviceProvider"`
	PartOf            *Reference                `bson:"partOf" json:"partOf"`
}
type EncounterStatusHistory struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Status            EncounterStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            Period          `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterClassHistory struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Class             Coding      `bson:"class,omitempty" json:"class,omitempty"`
	Period            Period      `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterParticipant struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Type              []CodeableConcept `bson:"type" json:"type"`
	Period            *Period           `bson:"period" json:"period"`
	Individual        *Reference        `bson:"individual" json:"individual"`
}
type EncounterDiagnosis struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Condition         Reference        `bson:"condition,omitempty" json:"condition,omitempty"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Rank              *int             `bson:"rank" json:"rank"`
}
type EncounterHospitalization struct {
	Id                     *string           `bson:"id" json:"id"`
	Extension              []Extension       `bson:"extension" json:"extension"`
	ModifierExtension      []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier" json:"preAdmissionIdentifier"`
	Origin                 *Reference        `bson:"origin" json:"origin"`
	AdmitSource            *CodeableConcept  `bson:"admitSource" json:"admitSource"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission" json:"reAdmission"`
	DietPreference         []CodeableConcept `bson:"dietPreference" json:"dietPreference"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy" json:"specialCourtesy"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement" json:"specialArrangement"`
	Destination            *Reference        `bson:"destination" json:"destination"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition" json:"dischargeDisposition"`
}
type EncounterLocation struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Location          Reference                `bson:"location,omitempty" json:"location,omitempty"`
	Status            *EncounterLocationStatus `bson:"status" json:"status"`
	Period            *Period                  `bson:"period" json:"period"`
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

// UnmarshalEncounter unmarshalls a Encounter.
func UnmarshalEncounter(b []byte) (Encounter, error) {
	var encounter Encounter
	if err := json.Unmarshal(b, &encounter); err != nil {
		return encounter, err
	}
	return encounter, nil
}
