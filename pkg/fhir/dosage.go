package fhir

// Dosage is documented here http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	Id                       *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Sequence                 *int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Text                     *string           `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	PatientInstruction       *string           `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	Timing                   *Timing           `bson:"timing,omitempty" json:"timing,omitempty"`
	Site                     *CodeableConcept  `bson:"site,omitempty" json:"site,omitempty"`
	Route                    *CodeableConcept  `bson:"route,omitempty" json:"route,omitempty"`
	Method                   *CodeableConcept  `bson:"method,omitempty" json:"method,omitempty"`
	MaxDosePerPeriod         *Ratio            `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity         `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity         `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
}
