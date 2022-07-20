package fhir

import "encoding/json"

// PaymentReconciliation is documented here http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
	Id                  *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string                            `bson:"status,omitempty" json:"status,omitempty"`
	Period              *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Created             *string                            `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference                         `bson:"organization,omitempty" json:"organization,omitempty"`
	Request             *Reference                         `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *CodeableConcept                   `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         *string                            `bson:"disposition,omitempty" json:"disposition,omitempty"`
	RequestProvider     *Reference                         `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                         `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Detail              []PaymentReconciliationDetail      `bson:"detail,omitempty" json:"detail,omitempty"`
	Form                *CodeableConcept                   `bson:"form,omitempty" json:"form,omitempty"`
	Total               *Money                             `bson:"total,omitempty" json:"total,omitempty"`
	ProcessNote         []PaymentReconciliationProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
}
type PaymentReconciliationDetail struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Request           *Reference      `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference      `bson:"response,omitempty" json:"response,omitempty"`
	Submitter         *Reference      `bson:"submitter,omitempty" json:"submitter,omitempty"`
	Payee             *Reference      `bson:"payee,omitempty" json:"payee,omitempty"`
	Date              *string         `bson:"date,omitempty" json:"date,omitempty"`
	Amount            *Money          `bson:"amount,omitempty" json:"amount,omitempty"`
}
type PaymentReconciliationProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
}
type OtherPaymentReconciliation PaymentReconciliation

// MarshalJSON marshals the given PaymentReconciliation as JSON into a byte slice
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
}

// UnmarshalPaymentReconciliation unmarshals a PaymentReconciliation.
func UnmarshalPaymentReconciliation(b []byte) (PaymentReconciliation, error) {
	var paymentReconciliation PaymentReconciliation
	if err := json.Unmarshal(b, &paymentReconciliation); err != nil {
		return paymentReconciliation, err
	}
	return paymentReconciliation, nil
}
