package fhir

import "encoding/json"

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string          `bson:"status,omitempty" json:"status,omitempty"`
	Request           *Reference       `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference       `bson:"response,omitempty" json:"response,omitempty"`
	StatusDate        *string          `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	Created           *string          `bson:"created,omitempty" json:"created,omitempty"`
	Target            *Reference       `bson:"target,omitempty" json:"target,omitempty"`
	Provider          *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization      *Reference       `bson:"organization,omitempty" json:"organization,omitempty"`
	PaymentStatus     *CodeableConcept `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty"`
}
type OtherPaymentNotice PaymentNotice

// MarshalJSON marshals the given PaymentNotice as JSON into a byte slice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
}

// UnmarshalPaymentNotice unmarshals a PaymentNotice.
func UnmarshalPaymentNotice(b []byte) (PaymentNotice, error) {
	var paymentNotice PaymentNotice
	if err := json.Unmarshal(b, &paymentNotice); err != nil {
		return paymentNotice, err
	}
	return paymentNotice, nil
}
