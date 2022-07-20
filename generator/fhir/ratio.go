package fhir

// Ratio is documented here http://hl7.org/fhir/StructureDefinition/Ratio
type Ratio struct {
	Id          *string     `bson:"id" json:"id"`
	Extension   []Extension `bson:"extension" json:"extension"`
	Numerator   *Quantity   `bson:"numerator" json:"numerator"`
	Denominator *Quantity   `bson:"denominator" json:"denominator"`
}
