package fhir

import "encoding/json"

// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
type Consent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string            `bson:"status" json:"status"`
	Category          []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Patient           Reference         `bson:"patient" json:"patient"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	DateTime          *string           `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Actor             []ConsentActor    `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Organization      []Reference       `bson:"organization,omitempty" json:"organization,omitempty"`
	Policy            []ConsentPolicy   `bson:"policy,omitempty" json:"policy,omitempty"`
	PolicyRule        *string           `bson:"policyRule,omitempty" json:"policyRule,omitempty"`
	SecurityLabel     []Coding          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []Coding          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	DataPeriod        *Period           `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []ConsentData     `bson:"data,omitempty" json:"data,omitempty"`
	Except            []ConsentExcept   `bson:"except,omitempty" json:"except,omitempty"`
}
type ConsentActor struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role" json:"role"`
}
type ConsentPolicy struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *string     `bson:"authority,omitempty" json:"authority,omitempty"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ConsentData struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           string      `bson:"meaning" json:"meaning"`
	Reference         Reference   `bson:"reference" json:"reference"`
}
type ConsentExcept struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string               `bson:"type" json:"type"`
	Period            *Period              `bson:"period,omitempty" json:"period,omitempty"`
	Actor             []ConsentExceptActor `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []CodeableConcept    `bson:"action,omitempty" json:"action,omitempty"`
	SecurityLabel     []Coding             `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []Coding             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Class             []Coding             `bson:"class,omitempty" json:"class,omitempty"`
	Code              []Coding             `bson:"code,omitempty" json:"code,omitempty"`
	DataPeriod        *Period              `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []ConsentExceptData  `bson:"data,omitempty" json:"data,omitempty"`
}
type ConsentExceptActor struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role" json:"role"`
}
type ConsentExceptData struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           string      `bson:"meaning" json:"meaning"`
	Reference         Reference   `bson:"reference" json:"reference"`
}
type OtherConsent Consent

// MarshalJSON marshals the given Consent as JSON into a byte slice
func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConsent
		ResourceType string `json:"resourceType"`
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
}

// UnmarshalConsent unmarshals a Consent.
func UnmarshalConsent(b []byte) (Consent, error) {
	var consent Consent
	if err := json.Unmarshal(b, &consent); err != nil {
		return consent, err
	}
	return consent, nil
}
