package fhir

import "encoding/json"

// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             *Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition             []string                `bson:"definition,omitempty" json:"definition,omitempty"`
	Status                 string                  `bson:"status" json:"status"`
	PartOf                 []Reference             `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Code                   CodeableConcept         `bson:"code" json:"code"`
	Participant            []ChargeItemParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	PerformingOrganization *Reference              `bson:"performingOrganization,omitempty" json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference              `bson:"requestingOrganization,omitempty" json:"requestingOrganization,omitempty"`
	Quantity               *Quantity               `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Bodysite               []CodeableConcept       `bson:"bodysite,omitempty" json:"bodysite,omitempty"`
	FactorOverride         *string                 `bson:"factorOverride,omitempty" json:"factorOverride,omitempty"`
	PriceOverride          *Money                  `bson:"priceOverride,omitempty" json:"priceOverride,omitempty"`
	OverrideReason         *string                 `bson:"overrideReason,omitempty" json:"overrideReason,omitempty"`
	EnteredDate            *string                 `bson:"enteredDate,omitempty" json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept       `bson:"reason,omitempty" json:"reason,omitempty"`
	Account                []Reference             `bson:"account,omitempty" json:"account,omitempty"`
	Note                   []Annotation            `bson:"note,omitempty" json:"note,omitempty"`
	SupportingInformation  []Reference             `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}
type ChargeItemParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type OtherChargeItem ChargeItem

// MarshalJSON marshals the given ChargeItem as JSON into a byte slice
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

// UnmarshalChargeItem unmarshals a ChargeItem.
func UnmarshalChargeItem(b []byte) (ChargeItem, error) {
	var chargeItem ChargeItem
	if err := json.Unmarshal(b, &chargeItem); err != nil {
		return chargeItem, err
	}
	return chargeItem, nil
}
