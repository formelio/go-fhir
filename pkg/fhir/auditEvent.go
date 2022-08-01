package fhir

import "encoding/json"

// AuditEvent is documented here http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding             `bson:"type,omitempty" json:"type,omitempty"`
	Subtype           []Coding           `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action            *AuditEventAction  `bson:"action,omitempty" json:"action,omitempty"`
	Recorded          string             `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome           *AuditEventOutcome `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc       *string            `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent    []CodeableConcept  `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent             []AuditEventAgent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source            AuditEventSource   `bson:"source,omitempty" json:"source,omitempty"`
	Entity            []AuditEventEntity `bson:"entity,omitempty" json:"entity,omitempty"`
}
type AuditEventAgent struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept       `bson:"role,omitempty" json:"role,omitempty"`
	Reference         *Reference              `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId            *Identifier             `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId             *string                 `bson:"altId,omitempty" json:"altId,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Requestor         bool                    `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location          *Reference              `bson:"location,omitempty" json:"location,omitempty"`
	Policy            []string                `bson:"policy,omitempty" json:"policy,omitempty"`
	Media             *Coding                 `bson:"media,omitempty" json:"media,omitempty"`
	Network           *AuditEventAgentNetwork `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse      []CodeableConcept       `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}
type AuditEventAgentNetwork struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Address           *string                     `bson:"address,omitempty" json:"address,omitempty"`
	Type              *AuditEventAgentNetworkType `bson:"type,omitempty" json:"type,omitempty"`
}
type AuditEventSource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Site              *string     `bson:"site,omitempty" json:"site,omitempty"`
	Identifier        Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}
type AuditEventEntity struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference         *Reference               `bson:"reference,omitempty" json:"reference,omitempty"`
	Type              *Coding                  `bson:"type,omitempty" json:"type,omitempty"`
	Role              *Coding                  `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle         *Coding                  `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel     []Coding                 `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name              *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                  `bson:"description,omitempty" json:"description,omitempty"`
	Query             *string                  `bson:"query,omitempty" json:"query,omitempty"`
	Detail            []AuditEventEntityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type AuditEventEntityDetail struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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
