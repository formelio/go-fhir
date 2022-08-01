package fhir

import (
	"bytes"
	"encoding/json"
)

// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                    *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource               `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []*Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                *SupplyRequestStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Category              *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	Priority              *RequestPriority          `bson:"priority,omitempty" json:"priority,omitempty"`
	OrderedItem           *SupplyRequestOrderedItem `bson:"orderedItem,omitempty" json:"orderedItem,omitempty"`
	OccurrenceDateTime    *string                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                   `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                   `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn            *string                   `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester             *SupplyRequestRequester   `bson:"requester,omitempty" json:"requester,omitempty"`
	Supplier              []*Reference              `bson:"supplier,omitempty" json:"supplier,omitempty"`
	ReasonCodeableConcept *CodeableConcept          `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	DeliverFrom           *Reference                `bson:"deliverFrom,omitempty" json:"deliverFrom,omitempty"`
	DeliverTo             *Reference                `bson:"deliverTo,omitempty" json:"deliverTo,omitempty"`
}
type SupplyRequestOrderedItem struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity            Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
}
type SupplyRequestRequester struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference    `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference   `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherSupplyRequest
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
	return buffer.Bytes(), err
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
