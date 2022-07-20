package fhir

// Address is documented here http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
	Id         *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Use        *string     `bson:"use,omitempty" json:"use,omitempty"`
	Type       *string     `bson:"type,omitempty" json:"type,omitempty"`
	Text       *string     `bson:"text,omitempty" json:"text,omitempty"`
	Line       []string    `bson:"line,omitempty" json:"line,omitempty"`
	City       *string     `bson:"city,omitempty" json:"city,omitempty"`
	District   *string     `bson:"district,omitempty" json:"district,omitempty"`
	State      *string     `bson:"state,omitempty" json:"state,omitempty"`
	PostalCode *string     `bson:"postalCode,omitempty" json:"postalCode,omitempty"`
	Country    *string     `bson:"country,omitempty" json:"country,omitempty"`
	Period     *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
