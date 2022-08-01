package fhir

// TriggerDefinition is documented here http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Type                 TriggerType      `bson:"type,omitempty" json:"type,omitempty"`
	EventName            *string          `bson:"eventName,omitempty" json:"eventName,omitempty"`
	EventTimingTiming    *Timing          `bson:"eventTimingTiming,omitempty" json:"eventTimingTiming,omitempty"`
	EventTimingReference *Reference       `bson:"eventTimingReference,omitempty" json:"eventTimingReference,omitempty"`
	EventTimingDate      *string          `bson:"eventTimingDate,omitempty" json:"eventTimingDate,omitempty"`
	EventTimingDateTime  *string          `bson:"eventTimingDateTime,omitempty" json:"eventTimingDateTime,omitempty"`
	EventData            *DataRequirement `bson:"eventData,omitempty" json:"eventData,omitempty"`
}
