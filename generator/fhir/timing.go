package fhir

// Timing is documented here http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	Id        *string          `bson:"id" json:"id"`
	Extension []Extension      `bson:"extension" json:"extension"`
	Event     []string         `bson:"event" json:"event"`
	Repeat    *TimingRepeat    `bson:"repeat" json:"repeat"`
	Code      *CodeableConcept `bson:"code" json:"code"`
}
type TimingRepeat struct {
	Id             *string      `bson:"id" json:"id"`
	Extension      []Extension  `bson:"extension" json:"extension"`
	BoundsDuration *Duration    `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty"`
	BoundsRange    *Range       `bson:"boundsRange,omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod   *Period      `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty"`
	Count          *int         `bson:"count" json:"count"`
	CountMax       *int         `bson:"countMax" json:"countMax"`
	Duration       *float64     `bson:"duration" json:"duration"`
	DurationMax    *float64     `bson:"durationMax" json:"durationMax"`
	DurationUnit   *string      `bson:"durationUnit" json:"durationUnit"`
	Frequency      *int         `bson:"frequency" json:"frequency"`
	FrequencyMax   *int         `bson:"frequencyMax" json:"frequencyMax"`
	Period         *float64     `bson:"period" json:"period"`
	PeriodMax      *float64     `bson:"periodMax" json:"periodMax"`
	PeriodUnit     *string      `bson:"periodUnit" json:"periodUnit"`
	DayOfWeek      []DaysOfWeek `bson:"dayOfWeek" json:"dayOfWeek"`
	TimeOfDay      []string     `bson:"timeOfDay" json:"timeOfDay"`
	When           []string     `bson:"when" json:"when"`
	Offset         *int         `bson:"offset" json:"offset"`
}
