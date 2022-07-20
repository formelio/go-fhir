package fhir

import "encoding/json"

// AuditEvent is documented here http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding             `bson:"type" json:"type"`
	Subtype           []Coding           `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action            *string            `bson:"action,omitempty" json:"action,omitempty"`
	Recorded          string             `bson:"recorded" json:"recorded"`
	Outcome           *string            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc       *string            `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent    []CodeableConcept  `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent             []AuditEventAgent  `bson:"agent" json:"agent"`
	Source            AuditEventSource   `bson:"source" json:"source"`
	Entity            []AuditEventEntity `bson:"entity,omitempty" json:"entity,omitempty"`
}
type AuditEventAgent struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept       `bson:"role,omitempty" json:"role,omitempty"`
	UserId            *Identifier             `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId             *string                 `bson:"altId,omitempty" json:"altId,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Requestor         bool                    `bson:"requestor" json:"requestor"`
	Location          *Reference              `bson:"location,omitempty" json:"location,omitempty"`
	Policy            []string                `bson:"policy,omitempty" json:"policy,omitempty"`
	Media             *Coding                 `bson:"media,omitempty" json:"media,omitempty"`
	Network           *AuditEventAgentNetwork `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse      []CodeableConcept       `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}
type AuditEventAgentNetwork struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Address           *string     `bson:"address,omitempty" json:"address,omitempty"`
	Type              *string     `bson:"type,omitempty" json:"type,omitempty"`
}
type AuditEventSource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Site              *string     `bson:"site,omitempty" json:"site,omitempty"`
	Identifier        Identifier  `bson:"identifier" json:"identifier"`
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
	Type              string      `bson:"type" json:"type"`
	Value             string      `bson:"value" json:"value"`
}
type OtherAuditEvent AuditEvent

// MarshalJSON marshals the given AuditEvent as JSON into a byte slice
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAuditEvent: OtherAuditEvent(r),
		ResourceType:    "AuditEvent",
	})
}

// UnmarshalAuditEvent unmarshals a AuditEvent.
func UnmarshalAuditEvent(b []byte) (AuditEvent, error) {
	var auditEvent AuditEvent
	if err := json.Unmarshal(b, &auditEvent); err != nil {
		return auditEvent, err
	}
	return auditEvent, nil
}
