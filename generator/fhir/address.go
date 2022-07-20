package fhir

// Address is documented here http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
	Id         *string      `bson:"id" json:"id"`
	Extension  []Extension  `bson:"extension" json:"extension"`
	Use        *AddressUse  `bson:"use" json:"use"`
	Type       *AddressType `bson:"type" json:"type"`
	Text       *string      `bson:"text" json:"text"`
	Line       []string     `bson:"line" json:"line"`
	City       *string      `bson:"city" json:"city"`
	District   *string      `bson:"district" json:"district"`
	State      *string      `bson:"state" json:"state"`
	PostalCode *string      `bson:"postalCode" json:"postalCode"`
	Country    *string      `bson:"country" json:"country"`
	Period     *Period      `bson:"period" json:"period"`
}
