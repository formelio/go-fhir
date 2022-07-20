package fhir

import "encoding/json"

// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                        *string                             `bson:"id" json:"id"`
	Meta                      *Meta                               `bson:"meta" json:"meta"`
	ImplicitRules             *string                             `bson:"implicitRules" json:"implicitRules"`
	Language                  *string                             `bson:"language" json:"language"`
	Text                      *Narrative                          `bson:"text" json:"text"`
	Contained                 []json.RawMessage                   `bson:"contained" json:"contained"`
	Extension                 []Extension                         `bson:"extension" json:"extension"`
	ModifierExtension         []Extension                         `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                []Identifier                        `bson:"identifier" json:"identifier"`
	Definition                []Reference                         `bson:"definition" json:"definition"`
	PartOf                    []Reference                         `bson:"partOf" json:"partOf"`
	Status                    MedicationAdministrationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept                    `bson:"category" json:"category"`
	MedicationCodeableConcept *CodeableConcept                    `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                          `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                          `bson:"context" json:"context"`
	SupportingInformation     []Reference                         `bson:"supportingInformation" json:"supportingInformation"`
	EffectiveDateTime         *string                             `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Performer                 []MedicationAdministrationPerformer `bson:"performer" json:"performer"`
	NotGiven                  *bool                               `bson:"notGiven" json:"notGiven"`
	ReasonNotGiven            []CodeableConcept                   `bson:"reasonNotGiven" json:"reasonNotGiven"`
	ReasonCode                []CodeableConcept                   `bson:"reasonCode" json:"reasonCode"`
	ReasonReference           []Reference                         `bson:"reasonReference" json:"reasonReference"`
	Prescription              *Reference                          `bson:"prescription" json:"prescription"`
	Device                    []Reference                         `bson:"device" json:"device"`
	Note                      []Annotation                        `bson:"note" json:"note"`
	Dosage                    *MedicationAdministrationDosage     `bson:"dosage" json:"dosage"`
	EventHistory              []Reference                         `bson:"eventHistory" json:"eventHistory"`
}
type MedicationAdministrationPerformer struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Actor             Reference   `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type MedicationAdministrationDosage struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Text              *string          `bson:"text" json:"text"`
	Site              *CodeableConcept `bson:"site" json:"site"`
	Route             *CodeableConcept `bson:"route" json:"route"`
	Method            *CodeableConcept `bson:"method" json:"method"`
	Dose              *Quantity        `bson:"dose" json:"dose"`
	RateRatio         *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateQuantity      *Quantity        `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
}
type OtherMedicationAdministration MedicationAdministration

// MarshalJSON marshals the given MedicationAdministration as JSON into a byte slice
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
}

// UnmarshalMedicationAdministration unmarshalls a MedicationAdministration.
func UnmarshalMedicationAdministration(b []byte) (MedicationAdministration, error) {
	var medicationAdministration MedicationAdministration
	if err := json.Unmarshal(b, &medicationAdministration); err != nil {
		return medicationAdministration, err
	}
	return medicationAdministration, nil
}
