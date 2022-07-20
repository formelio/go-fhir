package fhir

// Dosage is documented here http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	Id                       *string           `bson:"id" json:"id"`
	Extension                []Extension       `bson:"extension" json:"extension"`
	Sequence                 *int              `bson:"sequence" json:"sequence"`
	Text                     *string           `bson:"text" json:"text"`
	AdditionalInstruction    []CodeableConcept `bson:"additionalInstruction" json:"additionalInstruction"`
	PatientInstruction       *string           `bson:"patientInstruction" json:"patientInstruction"`
	Timing                   *Timing           `bson:"timing" json:"timing"`
	AsNeededBoolean          *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept  `bson:"site" json:"site"`
	Route                    *CodeableConcept  `bson:"route" json:"route"`
	Method                   *CodeableConcept  `bson:"method" json:"method"`
	DoseRange                *Range            `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseQuantity             *Quantity         `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	MaxDosePerPeriod         *Ratio            `bson:"maxDosePerPeriod" json:"maxDosePerPeriod"`
	MaxDosePerAdministration *Quantity         `bson:"maxDosePerAdministration" json:"maxDosePerAdministration"`
	MaxDosePerLifetime       *Quantity         `bson:"maxDosePerLifetime" json:"maxDosePerLifetime"`
	RateRatio                *Ratio            `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange                *Range            `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	RateQuantity             *Quantity         `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
}
