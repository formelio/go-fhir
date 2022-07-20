package fhir

// UsageContext is documented here http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
	Id                   *string          `bson:"id" json:"id"`
	Extension            []Extension      `bson:"extension" json:"extension"`
	Code                 Coding           `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
}
