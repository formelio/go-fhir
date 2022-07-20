package fhir

import "encoding/json"

// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string                 `bson:"id" json:"id"`
	Meta                   *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules          *string                 `bson:"implicitRules" json:"implicitRules"`
	Language               *string                 `bson:"language" json:"language"`
	Text                   *Narrative              `bson:"text" json:"text"`
	Contained              []json.RawMessage       `bson:"contained" json:"contained"`
	Extension              []Extension             `bson:"extension" json:"extension"`
	ModifierExtension      []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Identifier             *Identifier             `bson:"identifier" json:"identifier"`
	Definition             []string                `bson:"definition" json:"definition"`
	Status                 ChargeItemStatus        `bson:"status,omitempty" json:"status,omitempty"`
	PartOf                 []Reference             `bson:"partOf" json:"partOf"`
	Code                   CodeableConcept         `bson:"code,omitempty" json:"code,omitempty"`
	Subject                Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                *Reference              `bson:"context" json:"context"`
	OccurrenceDateTime     *string                 `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period                 `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing                 `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Participant            []ChargeItemParticipant `bson:"participant" json:"participant"`
	PerformingOrganization *Reference              `bson:"performingOrganization" json:"performingOrganization"`
	RequestingOrganization *Reference              `bson:"requestingOrganization" json:"requestingOrganization"`
	Quantity               *Quantity               `bson:"quantity" json:"quantity"`
	Bodysite               []CodeableConcept       `bson:"bodysite" json:"bodysite"`
	FactorOverride         *float64                `bson:"factorOverride" json:"factorOverride"`
	PriceOverride          *Money                  `bson:"priceOverride" json:"priceOverride"`
	OverrideReason         *string                 `bson:"overrideReason" json:"overrideReason"`
	Enterer                *Reference              `bson:"enterer" json:"enterer"`
	EnteredDate            *string                 `bson:"enteredDate" json:"enteredDate"`
	Reason                 []CodeableConcept       `bson:"reason" json:"reason"`
	Service                []Reference             `bson:"service" json:"service"`
	Account                []Reference             `bson:"account" json:"account"`
	Note                   []Annotation            `bson:"note" json:"note"`
	SupportingInformation  []Reference             `bson:"supportingInformation" json:"supportingInformation"`
}
type ChargeItemParticipant struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
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

// UnmarshalChargeItem unmarshalls a ChargeItem.
func UnmarshalChargeItem(b []byte) (ChargeItem, error) {
	var chargeItem ChargeItem
	if err := json.Unmarshal(b, &chargeItem); err != nil {
		return chargeItem, err
	}
	return chargeItem, nil
}
