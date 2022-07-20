package fhir

import "encoding/json"

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Status            *string           `bson:"status" json:"status"`
	Request           *Reference        `bson:"request" json:"request"`
	Response          *Reference        `bson:"response" json:"response"`
	StatusDate        *string           `bson:"statusDate" json:"statusDate"`
	Created           *string           `bson:"created" json:"created"`
	Target            *Reference        `bson:"target" json:"target"`
	Provider          *Reference        `bson:"provider" json:"provider"`
	Organization      *Reference        `bson:"organization" json:"organization"`
	PaymentStatus     *CodeableConcept  `bson:"paymentStatus" json:"paymentStatus"`
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

// UnmarshalPaymentNotice unmarshalls a PaymentNotice.
func UnmarshalPaymentNotice(b []byte) (PaymentNotice, error) {
	var paymentNotice PaymentNotice
	if err := json.Unmarshal(b, &paymentNotice); err != nil {
		return paymentNotice, err
	}
	return paymentNotice, nil
}
