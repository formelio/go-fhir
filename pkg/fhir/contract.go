package fhir

import "encoding/json"

// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
	Id                *string              `bson:"id" json:"id"`
	Meta              *Meta                `bson:"meta" json:"meta"`
	ImplicitRules     *string              `bson:"implicitRules" json:"implicitRules"`
	Language          *string              `bson:"language" json:"language"`
	Text              *Narrative           `bson:"text" json:"text"`
	Contained         []json.RawMessage    `bson:"contained" json:"contained"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier          `bson:"identifier" json:"identifier"`
	Status            *string              `bson:"status" json:"status"`
	Issued            *string              `bson:"issued" json:"issued"`
	Applies           *Period              `bson:"applies" json:"applies"`
	Subject           []Reference          `bson:"subject" json:"subject"`
	Topic             []Reference          `bson:"topic" json:"topic"`
	Authority         []Reference          `bson:"authority" json:"authority"`
	Domain            []Reference          `bson:"domain" json:"domain"`
	Type              *CodeableConcept     `bson:"type" json:"type"`
	SubType           []CodeableConcept    `bson:"subType" json:"subType"`
	Action            []CodeableConcept    `bson:"action" json:"action"`
	ActionReason      []CodeableConcept    `bson:"actionReason" json:"actionReason"`
	DecisionType      *CodeableConcept     `bson:"decisionType" json:"decisionType"`
	ContentDerivative *CodeableConcept     `bson:"contentDerivative" json:"contentDerivative"`
	SecurityLabel     []Coding             `bson:"securityLabel" json:"securityLabel"`
	Agent             []ContractAgent      `bson:"agent" json:"agent"`
	Signer            []ContractSigner     `bson:"signer" json:"signer"`
	ValuedItem        []ContractValuedItem `bson:"valuedItem" json:"valuedItem"`
	Term              []ContractTerm       `bson:"term" json:"term"`
	BindingAttachment *Attachment          `bson:"bindingAttachment,omitempty" json:"bindingAttachment,omitempty"`
	BindingReference  *Reference           `bson:"bindingReference,omitempty" json:"bindingReference,omitempty"`
	Friendly          []ContractFriendly   `bson:"friendly" json:"friendly"`
	Legal             []ContractLegal      `bson:"legal" json:"legal"`
	Rule              []ContractRule       `bson:"rule" json:"rule"`
}
type ContractAgent struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Actor             Reference         `bson:"actor,omitempty" json:"actor,omitempty"`
	Role              []CodeableConcept `bson:"role" json:"role"`
}
type ContractSigner struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Type              Coding      `bson:"type,omitempty" json:"type,omitempty"`
	Party             Reference   `bson:"party,omitempty" json:"party,omitempty"`
	Signature         []Signature `bson:"signature,omitempty" json:"signature,omitempty"`
}
type ContractValuedItem struct {
	Id                    *string          `bson:"id" json:"id"`
	Extension             []Extension      `bson:"extension" json:"extension"`
	ModifierExtension     []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier" json:"identifier"`
	EffectiveTime         *string          `bson:"effectiveTime" json:"effectiveTime"`
	Quantity              *Quantity        `bson:"quantity" json:"quantity"`
	UnitPrice             *Money           `bson:"unitPrice" json:"unitPrice"`
	Factor                *float64         `bson:"factor" json:"factor"`
	Points                *float64         `bson:"points" json:"points"`
	Net                   *Money           `bson:"net" json:"net"`
}
type ContractTerm struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier              `bson:"identifier" json:"identifier"`
	Issued            *string                  `bson:"issued" json:"issued"`
	Applies           *Period                  `bson:"applies" json:"applies"`
	Type              *CodeableConcept         `bson:"type" json:"type"`
	SubType           *CodeableConcept         `bson:"subType" json:"subType"`
	Topic             []Reference              `bson:"topic" json:"topic"`
	Action            []CodeableConcept        `bson:"action" json:"action"`
	ActionReason      []CodeableConcept        `bson:"actionReason" json:"actionReason"`
	SecurityLabel     []Coding                 `bson:"securityLabel" json:"securityLabel"`
	Agent             []ContractTermAgent      `bson:"agent" json:"agent"`
	Text              *string                  `bson:"text" json:"text"`
	ValuedItem        []ContractTermValuedItem `bson:"valuedItem" json:"valuedItem"`
	Group             []ContractTerm           `bson:"group" json:"group"`
}
type ContractTermAgent struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Actor             Reference         `bson:"actor,omitempty" json:"actor,omitempty"`
	Role              []CodeableConcept `bson:"role" json:"role"`
}
type ContractTermValuedItem struct {
	Id                    *string          `bson:"id" json:"id"`
	Extension             []Extension      `bson:"extension" json:"extension"`
	ModifierExtension     []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier" json:"identifier"`
	EffectiveTime         *string          `bson:"effectiveTime" json:"effectiveTime"`
	Quantity              *Quantity        `bson:"quantity" json:"quantity"`
	UnitPrice             *Money           `bson:"unitPrice" json:"unitPrice"`
	Factor                *float64         `bson:"factor" json:"factor"`
	Points                *float64         `bson:"points" json:"points"`
	Net                   *Money           `bson:"net" json:"net"`
}
type ContractFriendly struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type ContractLegal struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type ContractRule struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type OtherContract Contract

// MarshalJSON marshals the given Contract as JSON into a byte slice
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType"`
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
}

// UnmarshalContract unmarshalls a Contract.
func UnmarshalContract(b []byte) (Contract, error) {
	var contract Contract
	if err := json.Unmarshal(b, &contract); err != nil {
		return contract, err
	}
	return contract, nil
}
