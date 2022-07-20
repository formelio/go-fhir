package fhir

// SampledData is documented here http://hl7.org/fhir/StructureDefinition/SampledData
type SampledData struct {
	Id         *string     `bson:"id" json:"id"`
	Extension  []Extension `bson:"extension" json:"extension"`
	Origin     Quantity    `bson:"origin,omitempty" json:"origin,omitempty"`
	Period     float64     `bson:"period,omitempty" json:"period,omitempty"`
	Factor     *float64    `bson:"factor" json:"factor"`
	LowerLimit *float64    `bson:"lowerLimit" json:"lowerLimit"`
	UpperLimit *float64    `bson:"upperLimit" json:"upperLimit"`
	Dimensions int         `bson:"dimensions,omitempty" json:"dimensions,omitempty"`
	Data       string      `bson:"data,omitempty" json:"data,omitempty"`
}
