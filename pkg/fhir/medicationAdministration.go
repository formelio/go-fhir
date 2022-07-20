package fhir

import "encoding/json"

// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                    *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                              `bson:"status" json:"status"`
	Category              *CodeableConcept                    `bson:"category,omitempty" json:"category,omitempty"`
	SupportingInformation []Reference                         `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer             []MedicationAdministrationPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	NotGiven              *bool                               `bson:"notGiven,omitempty" json:"notGiven,omitempty"`
	ReasonNotGiven        []CodeableConcept                   `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
	ReasonCode            []CodeableConcept                   `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Prescription          *Reference                          `bson:"prescription,omitempty" json:"prescription,omitempty"`
	Device                []Reference                         `bson:"device,omitempty" json:"device,omitempty"`
	Note                  []Annotation                        `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                *MedicationAdministrationDosage     `bson:"dosage,omitempty" json:"dosage,omitempty"`
	EventHistory          []Reference                         `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationAdministrationPerformer struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type MedicationAdministrationDosage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Site              *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	Route             *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Dose              *Quantity        `bson:"dose,omitempty" json:"dose,omitempty"`
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

// UnmarshalMedicationAdministration unmarshals a MedicationAdministration.
func UnmarshalMedicationAdministration(b []byte) (MedicationAdministration, error) {
	var medicationAdministration MedicationAdministration
	if err := json.Unmarshal(b, &medicationAdministration); err != nil {
		return medicationAdministration, err
	}
	return medicationAdministration, nil
}
