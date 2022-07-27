package fhir

import "encoding/json"

// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string                 `bson:"id" json:"id"`
	Meta                   *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules          *string                 `bson:"implicitRules" json:"implicitRules"`
	Language               *string                 `bson:"language" json:"language"`
	Text                   *Narrative              `bson:"text" json:"text"`
	RawContained           []json.RawMessage       `bson:"contained" json:"contained"`
	Contained              []IResource             `bson:"-" json:"-"`
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

// OtherChargeItem is a helper type to use the default implementations of Marshall and Unmarshal
type OtherChargeItem ChargeItem

// MarshalJSON marshals the given ChargeItem as JSON into a byte slice
func (r ChargeItem) MarshalJSON() ([]byte, error) {
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
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ChargeItem
func (r *ChargeItem) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherChargeItem)(r)); err != nil {
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
func (r ChargeItem) GetResourceType() ResourceType {
	return ResourceTypeChargeItem
}
