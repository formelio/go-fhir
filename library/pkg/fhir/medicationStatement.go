package fhir

import "encoding/json"

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string            `bson:"status" json:"status"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	DateAsserted      *string           `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	DerivedFrom       []Reference       `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Taken             string            `bson:"taken" json:"taken"`
	ReasonNotTaken    []CodeableConcept `bson:"reasonNotTaken,omitempty" json:"reasonNotTaken,omitempty"`
	ReasonCode        []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	Dosage            []Dosage          `bson:"dosage,omitempty" json:"dosage,omitempty"`
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

// UnmarshalMedicationStatement unmarshals a MedicationStatement.
func UnmarshalMedicationStatement(b []byte) (MedicationStatement, error) {
	var medicationStatement MedicationStatement
	if err := json.Unmarshal(b, &medicationStatement); err != nil {
		return medicationStatement, err
	}
	return medicationStatement, nil
}
