package fhir

import "encoding/json"

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                        *string                   `bson:"id" json:"id"`
	Meta                      *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules             *string                   `bson:"implicitRules" json:"implicitRules"`
	Language                  *string                   `bson:"language" json:"language"`
	Text                      *Narrative                `bson:"text" json:"text"`
	Contained                 []json.RawMessage         `bson:"contained" json:"contained"`
	Extension                 []Extension               `bson:"extension" json:"extension"`
	ModifierExtension         []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                []Identifier              `bson:"identifier" json:"identifier"`
	BasedOn                   []Reference               `bson:"basedOn" json:"basedOn"`
	PartOf                    []Reference               `bson:"partOf" json:"partOf"`
	Context                   *Reference                `bson:"context" json:"context"`
	Status                    MedicationStatementStatus `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept          `bson:"category" json:"category"`
	MedicationCodeableConcept *CodeableConcept          `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	EffectiveDateTime         *string                   `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	DateAsserted              *string                   `bson:"dateAsserted" json:"dateAsserted"`
	InformationSource         *Reference                `bson:"informationSource" json:"informationSource"`
	Subject                   Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	DerivedFrom               []Reference               `bson:"derivedFrom" json:"derivedFrom"`
	Taken                     MedicationStatementTaken  `bson:"taken,omitempty" json:"taken,omitempty"`
	ReasonNotTaken            []CodeableConcept         `bson:"reasonNotTaken" json:"reasonNotTaken"`
	ReasonCode                []CodeableConcept         `bson:"reasonCode" json:"reasonCode"`
	ReasonReference           []Reference               `bson:"reasonReference" json:"reasonReference"`
	Note                      []Annotation              `bson:"note" json:"note"`
	Dosage                    []Dosage                  `bson:"dosage" json:"dosage"`
}
type OtherMedicationStatement MedicationStatement

// MarshalJSON marshals the given MedicationStatement as JSON into a byte slice
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
}

// UnmarshalMedicationStatement unmarshalls a MedicationStatement.
func UnmarshalMedicationStatement(b []byte) (MedicationStatement, error) {
	var medicationStatement MedicationStatement
	if err := json.Unmarshal(b, &medicationStatement); err != nil {
		return medicationStatement, err
	}
	return medicationStatement, nil
}
