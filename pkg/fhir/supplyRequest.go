package fhir

import "encoding/json"

// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                    *string                   `bson:"id" json:"id"`
	Meta                  *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules         *string                   `bson:"implicitRules" json:"implicitRules"`
	Language              *string                   `bson:"language" json:"language"`
	Text                  *Narrative                `bson:"text" json:"text"`
	RawContained          []json.RawMessage         `bson:"contained" json:"contained"`
	Contained             []IResource               `bson:"-" json:"-"`
	Extension             []Extension               `bson:"extension" json:"extension"`
	ModifierExtension     []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            *Identifier               `bson:"identifier" json:"identifier"`
	Status                *SupplyRequestStatus      `bson:"status" json:"status"`
	Category              *CodeableConcept          `bson:"category" json:"category"`
	Priority              *RequestPriority          `bson:"priority" json:"priority"`
	OrderedItem           *SupplyRequestOrderedItem `bson:"orderedItem" json:"orderedItem"`
	OccurrenceDateTime    *string                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                   `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                   `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn            *string                   `bson:"authoredOn" json:"authoredOn"`
	Requester             *SupplyRequestRequester   `bson:"requester" json:"requester"`
	Supplier              []Reference               `bson:"supplier" json:"supplier"`
	ReasonCodeableConcept *CodeableConcept          `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	DeliverFrom           *Reference                `bson:"deliverFrom" json:"deliverFrom"`
	DeliverTo             *Reference                `bson:"deliverTo" json:"deliverTo"`
}
type SupplyRequestOrderedItem struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Quantity            Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
}
type SupplyRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}

// OtherSupplyRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSupplyRequest SupplyRequest

// MarshalJSON marshals the given SupplyRequest as JSON into a byte slice
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
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
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into SupplyRequest
func (r *SupplyRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSupplyRequest)(r)); err != nil {
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
func (r SupplyRequest) GetResourceType() ResourceType {
	return ResourceTypeSupplyRequest
}
