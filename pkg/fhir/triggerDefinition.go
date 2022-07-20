package fhir

// TriggerDefinition is documented here http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id                   *string          `bson:"id" json:"id"`
	Extension            []Extension      `bson:"extension" json:"extension"`
	Type                 TriggerType      `bson:"type,omitempty" json:"type,omitempty"`
	EventName            *string          `bson:"eventName" json:"eventName"`
	EventTimingTiming    *Timing          `bson:"eventTimingTiming,omitempty" json:"eventTimingTiming,omitempty"`
	EventTimingReference *Reference       `bson:"eventTimingReference,omitempty" json:"eventTimingReference,omitempty"`
	EventTimingDate      *string          `bson:"eventTimingDate,omitempty" json:"eventTimingDate,omitempty"`
	EventTimingDateTime  *string          `bson:"eventTimingDateTime,omitempty" json:"eventTimingDateTime,omitempty"`
	EventData            *DataRequirement `bson:"eventData" json:"eventData"`
}
