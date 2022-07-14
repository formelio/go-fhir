package fhir

// UsageContext is documented here http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Code      Coding      `bson:"code" json:"code"`
}
