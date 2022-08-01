package fhir

import (
	"bytes"
	"encoding/json"
)

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string           `bson:"status,omitempty" json:"status,omitempty"`
	Request           *Reference        `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference        `bson:"response,omitempty" json:"response,omitempty"`
	StatusDate        *string           `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	Created           *string           `bson:"created,omitempty" json:"created,omitempty"`
	Target            *Reference        `bson:"target,omitempty" json:"target,omitempty"`
	Provider          *Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization      *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	PaymentStatus     *CodeableConcept  `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty"`
}

// OtherPaymentNotice is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPaymentNotice PaymentNotice

// MarshalJSON marshals the given PaymentNotice as JSON into a byte slice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
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
		OtherPaymentNotice
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into PaymentNotice
func (r *PaymentNotice) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPaymentNotice)(r)); err != nil {
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
func (r PaymentNotice) GetResourceType() ResourceType {
	return ResourceTypePaymentNotice
}
