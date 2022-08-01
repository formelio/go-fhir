package fhir

// Timing is documented here http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	Event     []string         `bson:"event,omitempty" json:"event,omitempty"`
	Repeat    *TimingRepeat    `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Code      *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type TimingRepeat struct {
	Id             *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension      []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	BoundsDuration *Duration    `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty"`
	BoundsRange    *Range       `bson:"boundsRange,omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod   *Period      `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty"`
	Count          *int         `bson:"count,omitempty" json:"count,omitempty"`
	CountMax       *int         `bson:"countMax,omitempty" json:"countMax,omitempty"`
	Duration       *float64     `bson:"duration,omitempty" json:"duration,omitempty"`
	DurationMax    *float64     `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	DurationUnit   *string      `bson:"durationUnit,omitempty" json:"durationUnit,omitempty"`
	Frequency      *int         `bson:"frequency,omitempty" json:"frequency,omitempty"`
	FrequencyMax   *int         `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	Period         *float64     `bson:"period,omitempty" json:"period,omitempty"`
	PeriodMax      *float64     `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	PeriodUnit     *string      `bson:"periodUnit,omitempty" json:"periodUnit,omitempty"`
	DayOfWeek      []DaysOfWeek `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	TimeOfDay      []string     `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty"`
	When           []string     `bson:"when,omitempty" json:"when,omitempty"`
	Offset         *int         `bson:"offset,omitempty" json:"offset,omitempty"`
}
