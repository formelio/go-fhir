package fhir

import (
	"bytes"
	"encoding/json"
)

// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
type Consent struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ConsentState       `bson:"status,omitempty" json:"status,omitempty"`
	Category          []*CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Patient           Reference          `bson:"patient,omitempty" json:"patient,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	DateTime          *string            `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	ConsentingParty   []*Reference       `bson:"consentingParty,omitempty" json:"consentingParty,omitempty"`
	Actor             []*ConsentActor    `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []*CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Organization      []*Reference       `bson:"organization,omitempty" json:"organization,omitempty"`
	SourceAttachment  *Attachment        `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	SourceIdentifier  *Identifier        `bson:"sourceIdentifier,omitempty" json:"sourceIdentifier,omitempty"`
	SourceReference   *Reference         `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	Policy            []*ConsentPolicy   `bson:"policy,omitempty" json:"policy,omitempty"`
	PolicyRule        *string            `bson:"policyRule,omitempty" json:"policyRule,omitempty"`
	SecurityLabel     []*Coding          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []*Coding          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	DataPeriod        *Period            `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []*ConsentData     `bson:"data,omitempty" json:"data,omitempty"`
	Except            []*ConsentExcept   `bson:"except,omitempty" json:"except,omitempty"`
}
type ConsentActor struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference         Reference       `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentPolicy struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *string      `bson:"authority,omitempty" json:"authority,omitempty"`
	Uri               *string      `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ConsentData struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaning `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Reference         Reference          `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentExcept struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ConsentExceptType     `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
	Actor             []*ConsentExceptActor `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []*CodeableConcept    `bson:"action,omitempty" json:"action,omitempty"`
	SecurityLabel     []*Coding             `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []*Coding             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Class             []*Coding             `bson:"class,omitempty" json:"class,omitempty"`
	Code              []*Coding             `bson:"code,omitempty" json:"code,omitempty"`
	DataPeriod        *Period               `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []*ConsentExceptData  `bson:"data,omitempty" json:"data,omitempty"`
}
type ConsentExceptActor struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference         Reference       `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentExceptData struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaning `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Reference         Reference          `bson:"reference,omitempty" json:"reference,omitempty"`
}

// OtherConsent is a helper type to use the default implementations of Marshall and Unmarshal
type OtherConsent Consent

// MarshalJSON marshals the given Consent as JSON into a byte slice
func (r Consent) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherConsent
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Consent
func (r *Consent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherConsent)(r)); err != nil {
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
func (r Consent) GetResourceType() ResourceType {
	return ResourceTypeConsent
}
