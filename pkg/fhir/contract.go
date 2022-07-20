package fhir

import "encoding/json"

// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string              `bson:"status,omitempty" json:"status,omitempty"`
	Issued            *string              `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period              `bson:"applies,omitempty" json:"applies,omitempty"`
	Subject           []Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Topic             []Reference          `bson:"topic,omitempty" json:"topic,omitempty"`
	Authority         []Reference          `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain            []Reference          `bson:"domain,omitempty" json:"domain,omitempty"`
	Type              *CodeableConcept     `bson:"type,omitempty" json:"type,omitempty"`
	SubType           []CodeableConcept    `bson:"subType,omitempty" json:"subType,omitempty"`
	Action            []CodeableConcept    `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []CodeableConcept    `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	DecisionType      *CodeableConcept     `bson:"decisionType,omitempty" json:"decisionType,omitempty"`
	ContentDerivative *CodeableConcept     `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	SecurityLabel     []Coding             `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent             []ContractAgent      `bson:"agent,omitempty" json:"agent,omitempty"`
	Signer            []ContractSigner     `bson:"signer,omitempty" json:"signer,omitempty"`
	ValuedItem        []ContractValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Term              []ContractTerm       `bson:"term,omitempty" json:"term,omitempty"`
	Friendly          []ContractFriendly   `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal             []ContractLegal      `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule              []ContractRule       `bson:"rule,omitempty" json:"rule,omitempty"`
}
type ContractAgent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractSigner struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding      `bson:"type" json:"type"`
	Signature         []Signature `bson:"signature" json:"signature"`
}
type ContractValuedItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime     *string     `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money      `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string     `bson:"factor,omitempty" json:"factor,omitempty"`
	Points            *string     `bson:"points,omitempty" json:"points,omitempty"`
	Net               *Money      `bson:"net,omitempty" json:"net,omitempty"`
}
type ContractTerm struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued            *string                  `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                  `bson:"applies,omitempty" json:"applies,omitempty"`
	Type              *CodeableConcept         `bson:"type,omitempty" json:"type,omitempty"`
	SubType           *CodeableConcept         `bson:"subType,omitempty" json:"subType,omitempty"`
	Topic             []Reference              `bson:"topic,omitempty" json:"topic,omitempty"`
	Action            []CodeableConcept        `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []CodeableConcept        `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	SecurityLabel     []Coding                 `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent             []ContractTermAgent      `bson:"agent,omitempty" json:"agent,omitempty"`
	Text              *string                  `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem        []ContractTermValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group             []ContractTerm           `bson:"group,omitempty" json:"group,omitempty"`
}
type ContractTermAgent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractTermValuedItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime     *string     `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money      `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string     `bson:"factor,omitempty" json:"factor,omitempty"`
	Points            *string     `bson:"points,omitempty" json:"points,omitempty"`
	Net               *Money      `bson:"net,omitempty" json:"net,omitempty"`
}
type ContractFriendly struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type ContractLegal struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type ContractRule struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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

// UnmarshalContract unmarshals a Contract.
func UnmarshalContract(b []byte) (Contract, error) {
	var contract Contract
	if err := json.Unmarshal(b, &contract); err != nil {
		return contract, err
	}
	return contract, nil
}
