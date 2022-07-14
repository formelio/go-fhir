package fhir

// Range is documented here http://hl7.org/fhir/StructureDefinition/Range
type Range struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Low       *Quantity   `bson:"low,omitempty" json:"low,omitempty"`
	High      *Quantity   `bson:"high,omitempty" json:"high,omitempty"`
}
