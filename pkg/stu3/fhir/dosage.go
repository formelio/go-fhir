package fhir

// Dosage is documented here http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	Id                       *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Sequence                 *int               `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Text                     *string            `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstruction    []*CodeableConcept `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	PatientInstruction       *string            `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	Timing                   *Timing            `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean          *bool              `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept   `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept   `bson:"site,omitempty" json:"site,omitempty"`
	Route                    *CodeableConcept   `bson:"route,omitempty" json:"route,omitempty"`
	Method                   *CodeableConcept   `bson:"method,omitempty" json:"method,omitempty"`
	DoseRange                *Range             `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseQuantity             *Quantity          `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	MaxDosePerPeriod         *Ratio             `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity          `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity          `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
	RateRatio                *Ratio             `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange                *Range             `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	RateQuantity             *Quantity          `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
}
