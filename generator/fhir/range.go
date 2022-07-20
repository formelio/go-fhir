package fhir

// Range is documented here http://hl7.org/fhir/StructureDefinition/Range
type Range struct {
	Id        *string     `bson:"id" json:"id"`
	Extension []Extension `bson:"extension" json:"extension"`
	Low       *Quantity   `bson:"low" json:"low"`
	High      *Quantity   `bson:"high" json:"high"`
}
