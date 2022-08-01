package fhir

import (
	"bytes"
	"encoding/json"
)

// PaymentReconciliation is documented here http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
	Id                  *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage                   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource                         `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []*Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []*Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string                             `bson:"status,omitempty" json:"status,omitempty"`
	Period              *Period                             `bson:"period,omitempty" json:"period,omitempty"`
	Created             *string                             `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference                          `bson:"organization,omitempty" json:"organization,omitempty"`
	Request             *Reference                          `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *CodeableConcept                    `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         *string                             `bson:"disposition,omitempty" json:"disposition,omitempty"`
	RequestProvider     *Reference                          `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                          `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Detail              []*PaymentReconciliationDetail      `bson:"detail,omitempty" json:"detail,omitempty"`
	Form                *CodeableConcept                    `bson:"form,omitempty" json:"form,omitempty"`
	Total               *Money                              `bson:"total,omitempty" json:"total,omitempty"`
	ProcessNote         []*PaymentReconciliationProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
}
type PaymentReconciliationDetail struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Request           *Reference      `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference      `bson:"response,omitempty" json:"response,omitempty"`
	Submitter         *Reference      `bson:"submitter,omitempty" json:"submitter,omitempty"`
	Payee             *Reference      `bson:"payee,omitempty" json:"payee,omitempty"`
	Date              *string         `bson:"date,omitempty" json:"date,omitempty"`
	Amount            *Money          `bson:"amount,omitempty" json:"amount,omitempty"`
}
type PaymentReconciliationProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
}

// OtherPaymentReconciliation is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPaymentReconciliation PaymentReconciliation

// MarshalJSON marshals the given PaymentReconciliation as JSON into a byte slice
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
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
		OtherPaymentReconciliation
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into PaymentReconciliation
func (r *PaymentReconciliation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPaymentReconciliation)(r)); err != nil {
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
func (r PaymentReconciliation) GetResourceType() ResourceType {
	return ResourceTypePaymentReconciliation
}
