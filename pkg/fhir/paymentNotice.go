package fhir

import "encoding/json"

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
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
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
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
