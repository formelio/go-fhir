package fhir

import (
	"bytes"
	"encoding/json"
)

// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string               `bson:"status,omitempty" json:"status,omitempty"`
	Issued            *string               `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period               `bson:"applies,omitempty" json:"applies,omitempty"`
	Subject           []*Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Topic             []*Reference          `bson:"topic,omitempty" json:"topic,omitempty"`
	Authority         []*Reference          `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain            []*Reference          `bson:"domain,omitempty" json:"domain,omitempty"`
	Type              *CodeableConcept      `bson:"type,omitempty" json:"type,omitempty"`
	SubType           []*CodeableConcept    `bson:"subType,omitempty" json:"subType,omitempty"`
	Action            []*CodeableConcept    `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []*CodeableConcept    `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	DecisionType      *CodeableConcept      `bson:"decisionType,omitempty" json:"decisionType,omitempty"`
	ContentDerivative *CodeableConcept      `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	SecurityLabel     []*Coding             `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent             []*ContractAgent      `bson:"agent,omitempty" json:"agent,omitempty"`
	Signer            []*ContractSigner     `bson:"signer,omitempty" json:"signer,omitempty"`
	ValuedItem        []*ContractValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Term              []*ContractTerm       `bson:"term,omitempty" json:"term,omitempty"`
	BindingAttachment *Attachment           `bson:"bindingAttachment,omitempty" json:"bindingAttachment,omitempty"`
	BindingReference  *Reference            `bson:"bindingReference,omitempty" json:"bindingReference,omitempty"`
	Friendly          []*ContractFriendly   `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal             []*ContractLegal      `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule              []*ContractRule       `bson:"rule,omitempty" json:"rule,omitempty"`
}
type ContractAgent struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Actor             Reference          `bson:"actor,omitempty" json:"actor,omitempty"`
	Role              []*CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractSigner struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding       `bson:"type,omitempty" json:"type,omitempty"`
	Party             Reference    `bson:"party,omitempty" json:"party,omitempty"`
	Signature         []Signature  `bson:"signature,omitempty" json:"signature,omitempty"`
}
type ContractValuedItem struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *string          `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Money           `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Money           `bson:"net,omitempty" json:"net,omitempty"`
}
type ContractTerm struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued            *string                   `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                   `bson:"applies,omitempty" json:"applies,omitempty"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	SubType           *CodeableConcept          `bson:"subType,omitempty" json:"subType,omitempty"`
	Topic             []*Reference              `bson:"topic,omitempty" json:"topic,omitempty"`
	Action            []*CodeableConcept        `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []*CodeableConcept        `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	SecurityLabel     []*Coding                 `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent             []*ContractTermAgent      `bson:"agent,omitempty" json:"agent,omitempty"`
	Text              *string                   `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem        []*ContractTermValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group             []*ContractTerm           `bson:"group,omitempty" json:"group,omitempty"`
}
type ContractTermAgent struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Actor             Reference          `bson:"actor,omitempty" json:"actor,omitempty"`
	Role              []*CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractTermValuedItem struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *string          `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Money           `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Money           `bson:"net,omitempty" json:"net,omitempty"`
}
type ContractFriendly struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment  `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference   `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type ContractLegal struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment  `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference   `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type ContractRule struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment  `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference   `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

// OtherContract is a helper type to use the default implementations of Marshall and Unmarshal
type OtherContract Contract

// MarshalJSON marshals the given Contract as JSON into a byte slice
func (r Contract) MarshalJSON() ([]byte, error) {
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
		OtherContract
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Contract
func (r *Contract) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherContract)(r)); err != nil {
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
func (r Contract) GetResourceType() ResourceType {
	return ResourceTypeContract
}
