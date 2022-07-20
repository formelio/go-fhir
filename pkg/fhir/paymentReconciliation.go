package fhir

import "encoding/json"

// PaymentReconciliation is documented here http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
	Id                  *string                            `bson:"id" json:"id"`
	Meta                *Meta                              `bson:"meta" json:"meta"`
	ImplicitRules       *string                            `bson:"implicitRules" json:"implicitRules"`
	Language            *string                            `bson:"language" json:"language"`
	Text                *Narrative                         `bson:"text" json:"text"`
	Contained           []json.RawMessage                  `bson:"contained" json:"contained"`
	Extension           []Extension                        `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                        `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier                       `bson:"identifier" json:"identifier"`
	Status              *string                            `bson:"status" json:"status"`
	Period              *Period                            `bson:"period" json:"period"`
	Created             *string                            `bson:"created" json:"created"`
	Organization        *Reference                         `bson:"organization" json:"organization"`
	Request             *Reference                         `bson:"request" json:"request"`
	Outcome             *CodeableConcept                   `bson:"outcome" json:"outcome"`
	Disposition         *string                            `bson:"disposition" json:"disposition"`
	RequestProvider     *Reference                         `bson:"requestProvider" json:"requestProvider"`
	RequestOrganization *Reference                         `bson:"requestOrganization" json:"requestOrganization"`
	Detail              []PaymentReconciliationDetail      `bson:"detail" json:"detail"`
	Form                *CodeableConcept                   `bson:"form" json:"form"`
	Total               *Money                             `bson:"total" json:"total"`
	ProcessNote         []PaymentReconciliationProcessNote `bson:"processNote" json:"processNote"`
}
type PaymentReconciliationDetail struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Type              CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Request           *Reference      `bson:"request" json:"request"`
	Response          *Reference      `bson:"response" json:"response"`
	Submitter         *Reference      `bson:"submitter" json:"submitter"`
	Payee             *Reference      `bson:"payee" json:"payee"`
	Date              *string         `bson:"date" json:"date"`
	Amount            *Money          `bson:"amount" json:"amount"`
}
type PaymentReconciliationProcessNote struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	Text              *string          `bson:"text" json:"text"`
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

// UnmarshalPaymentReconciliation unmarshalls a PaymentReconciliation.
func UnmarshalPaymentReconciliation(b []byte) (PaymentReconciliation, error) {
	var paymentReconciliation PaymentReconciliation
	if err := json.Unmarshal(b, &paymentReconciliation); err != nil {
		return paymentReconciliation, err
	}
	return paymentReconciliation, nil
}
