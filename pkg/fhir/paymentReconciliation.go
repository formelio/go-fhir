package fhir

import "encoding/json"

// PaymentReconciliation is documented here http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
	Id                  *string                            `bson:"id" json:"id"`
	Meta                *Meta                              `bson:"meta" json:"meta"`
	ImplicitRules       *string                            `bson:"implicitRules" json:"implicitRules"`
	Language            *string                            `bson:"language" json:"language"`
	Text                *Narrative                         `bson:"text" json:"text"`
	RawContained        []json.RawMessage                  `bson:"contained" json:"contained"`
	Contained           []IResource                        `bson:"-" json:"-"`
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
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
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
