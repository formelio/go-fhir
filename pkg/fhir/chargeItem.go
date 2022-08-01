package fhir

import "encoding/json"

// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             *Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition             []string                `bson:"definition,omitempty" json:"definition,omitempty"`
	Status                 ChargeItemStatus        `bson:"status,omitempty" json:"status,omitempty"`
	PartOf                 []Reference             `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Code                   CodeableConcept         `bson:"code,omitempty" json:"code,omitempty"`
	Subject                Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                *Reference              `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime     *string                 `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period                 `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing                 `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Participant            []ChargeItemParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	PerformingOrganization *Reference              `bson:"performingOrganization,omitempty" json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference              `bson:"requestingOrganization,omitempty" json:"requestingOrganization,omitempty"`
	Quantity               *Quantity               `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Bodysite               []CodeableConcept       `bson:"bodysite,omitempty" json:"bodysite,omitempty"`
	FactorOverride         *float64                `bson:"factorOverride,omitempty" json:"factorOverride,omitempty"`
	PriceOverride          *Money                  `bson:"priceOverride,omitempty" json:"priceOverride,omitempty"`
	OverrideReason         *string                 `bson:"overrideReason,omitempty" json:"overrideReason,omitempty"`
	Enterer                *Reference              `bson:"enterer,omitempty" json:"enterer,omitempty"`
	EnteredDate            *string                 `bson:"enteredDate,omitempty" json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept       `bson:"reason,omitempty" json:"reason,omitempty"`
	Service                []Reference             `bson:"service,omitempty" json:"service,omitempty"`
	Account                []Reference             `bson:"account,omitempty" json:"account,omitempty"`
	Note                   []Annotation            `bson:"note,omitempty" json:"note,omitempty"`
	SupportingInformation  []Reference             `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}
type ChargeItemParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
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
