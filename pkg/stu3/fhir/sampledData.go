package fhir

// SampledData is documented here http://hl7.org/fhir/StructureDefinition/SampledData
type SampledData struct {
	Id         *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Origin     Quantity    `bson:"origin,omitempty" json:"origin,omitempty"`
	Period     float64     `bson:"period,omitempty" json:"period,omitempty"`
	Factor     *float64    `bson:"factor,omitempty" json:"factor,omitempty"`
	LowerLimit *float64    `bson:"lowerLimit,omitempty" json:"lowerLimit,omitempty"`
	UpperLimit *float64    `bson:"upperLimit,omitempty" json:"upperLimit,omitempty"`
	Dimensions int         `bson:"dimensions,omitempty" json:"dimensions,omitempty"`
	Data       string      `bson:"data,omitempty" json:"data,omitempty"`
}
