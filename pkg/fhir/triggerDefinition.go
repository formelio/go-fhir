package fhir

// TriggerDefinition is documented here http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      string           `bson:"type" json:"type"`
	EventName *string          `bson:"eventName,omitempty" json:"eventName,omitempty"`
	EventData *DataRequirement `bson:"eventData,omitempty" json:"eventData,omitempty"`
}
