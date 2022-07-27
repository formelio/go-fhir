package fhir

import "encoding/json"

// AuditEvent is documented here http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	RawContained      []json.RawMessage  `bson:"contained" json:"contained"`
	Contained         []IResource        `bson:"-" json:"-"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Type              Coding             `bson:"type,omitempty" json:"type,omitempty"`
	Subtype           []Coding           `bson:"subtype" json:"subtype"`
	Action            *AuditEventAction  `bson:"action" json:"action"`
	Recorded          string             `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome           *AuditEventOutcome `bson:"outcome" json:"outcome"`
	OutcomeDesc       *string            `bson:"outcomeDesc" json:"outcomeDesc"`
	PurposeOfEvent    []CodeableConcept  `bson:"purposeOfEvent" json:"purposeOfEvent"`
	Agent             []AuditEventAgent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source            AuditEventSource   `bson:"source,omitempty" json:"source,omitempty"`
	Entity            []AuditEventEntity `bson:"entity" json:"entity"`
}
type AuditEventAgent struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Role              []CodeableConcept       `bson:"role" json:"role"`
	Reference         *Reference              `bson:"reference" json:"reference"`
	UserId            *Identifier             `bson:"userId" json:"userId"`
	AltId             *string                 `bson:"altId" json:"altId"`
	Name              *string                 `bson:"name" json:"name"`
	Requestor         bool                    `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location          *Reference              `bson:"location" json:"location"`
	Policy            []string                `bson:"policy" json:"policy"`
	Media             *Coding                 `bson:"media" json:"media"`
	Network           *AuditEventAgentNetwork `bson:"network" json:"network"`
	PurposeOfUse      []CodeableConcept       `bson:"purposeOfUse" json:"purposeOfUse"`
}
type AuditEventAgentNetwork struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Address           *string                     `bson:"address" json:"address"`
	Type              *AuditEventAgentNetworkType `bson:"type" json:"type"`
}
type AuditEventSource struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Site              *string     `bson:"site" json:"site"`
	Identifier        Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              []Coding    `bson:"type" json:"type"`
}
type AuditEventEntity struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier              `bson:"identifier" json:"identifier"`
	Reference         *Reference               `bson:"reference" json:"reference"`
	Type              *Coding                  `bson:"type" json:"type"`
	Role              *Coding                  `bson:"role" json:"role"`
	Lifecycle         *Coding                  `bson:"lifecycle" json:"lifecycle"`
	SecurityLabel     []Coding                 `bson:"securityLabel" json:"securityLabel"`
	Name              *string                  `bson:"name" json:"name"`
	Description       *string                  `bson:"description" json:"description"`
	Query             *string                  `bson:"query" json:"query"`
	Detail            []AuditEventEntityDetail `bson:"detail" json:"detail"`
}
type AuditEventEntityDetail struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Type              string      `bson:"type,omitempty" json:"type,omitempty"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}

// OtherAuditEvent is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAuditEvent AuditEvent

// MarshalJSON marshals the given AuditEvent as JSON into a byte slice
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAuditEvent: OtherAuditEvent(r),
		ResourceType:    "AuditEvent",
	})
}

// UnmarshalJSON unmarshals the given byte slice into AuditEvent
func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAuditEvent)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r AuditEvent) GetResourceType() ResourceType {
	return ResourceTypeAuditEvent
}
