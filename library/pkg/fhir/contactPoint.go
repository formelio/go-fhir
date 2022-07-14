package fhir

// ContactPoint is documented here http://hl7.org/fhir/StructureDefinition/ContactPoint
type ContactPoint struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	System    *string     `bson:"system,omitempty" json:"system,omitempty"`
	Value     *string     `bson:"value,omitempty" json:"value,omitempty"`
	Use       *string     `bson:"use,omitempty" json:"use,omitempty"`
	Rank      *int        `bson:"rank,omitempty" json:"rank,omitempty"`
	Period    *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
