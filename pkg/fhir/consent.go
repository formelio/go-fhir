package fhir

import "encoding/json"

// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
type Consent struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier       `bson:"identifier" json:"identifier"`
	Status            ConsentState      `bson:"status,omitempty" json:"status,omitempty"`
	Category          []CodeableConcept `bson:"category" json:"category"`
	Patient           Reference         `bson:"patient,omitempty" json:"patient,omitempty"`
	Period            *Period           `bson:"period" json:"period"`
	DateTime          *string           `bson:"dateTime" json:"dateTime"`
	ConsentingParty   []Reference       `bson:"consentingParty" json:"consentingParty"`
	Actor             []ConsentActor    `bson:"actor" json:"actor"`
	Action            []CodeableConcept `bson:"action" json:"action"`
	Organization      []Reference       `bson:"organization" json:"organization"`
	SourceAttachment  *Attachment       `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	SourceIdentifier  *Identifier       `bson:"sourceIdentifier,omitempty" json:"sourceIdentifier,omitempty"`
	SourceReference   *Reference        `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	Policy            []ConsentPolicy   `bson:"policy" json:"policy"`
	PolicyRule        *string           `bson:"policyRule" json:"policyRule"`
	SecurityLabel     []Coding          `bson:"securityLabel" json:"securityLabel"`
	Purpose           []Coding          `bson:"purpose" json:"purpose"`
	DataPeriod        *Period           `bson:"dataPeriod" json:"dataPeriod"`
	Data              []ConsentData     `bson:"data" json:"data"`
	Except            []ConsentExcept   `bson:"except" json:"except"`
}
type ConsentActor struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Role              CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference         Reference       `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentPolicy struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Authority         *string     `bson:"authority" json:"authority"`
	Uri               *string     `bson:"uri" json:"uri"`
}
type ConsentData struct {
	Id                *string            `bson:"id" json:"id"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Meaning           ConsentDataMeaning `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Reference         Reference          `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentExcept struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Type              ConsentExceptType    `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period              `bson:"period" json:"period"`
	Actor             []ConsentExceptActor `bson:"actor" json:"actor"`
	Action            []CodeableConcept    `bson:"action" json:"action"`
	SecurityLabel     []Coding             `bson:"securityLabel" json:"securityLabel"`
	Purpose           []Coding             `bson:"purpose" json:"purpose"`
	Class             []Coding             `bson:"class" json:"class"`
	Code              []Coding             `bson:"code" json:"code"`
	DataPeriod        *Period              `bson:"dataPeriod" json:"dataPeriod"`
	Data              []ConsentExceptData  `bson:"data" json:"data"`
}
type ConsentExceptActor struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Role              CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference         Reference       `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ConsentExceptData struct {
	Id                *string            `bson:"id" json:"id"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Meaning           ConsentDataMeaning `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Reference         Reference          `bson:"reference,omitempty" json:"reference,omitempty"`
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

// UnmarshalConsent unmarshalls a Consent.
func UnmarshalConsent(b []byte) (Consent, error) {
	var consent Consent
	if err := json.Unmarshal(b, &consent); err != nil {
		return consent, err
	}
	return consent, nil
}
