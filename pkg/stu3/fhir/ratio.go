package fhir

// Ratio is documented here http://hl7.org/fhir/StructureDefinition/Ratio
type Ratio struct {
	Id          *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Numerator   *Quantity    `bson:"numerator,omitempty" json:"numerator,omitempty"`
	Denominator *Quantity    `bson:"denominator,omitempty" json:"denominator,omitempty"`
}
