package fhir

import "encoding/json"

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	RawContained      []json.RawMessage  `bson:"contained" json:"contained"`
	Contained         []IResource        `bson:"-" json:"-"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier       `bson:"identifier" json:"identifier"`
	Status            *AccountStatus     `bson:"status" json:"status"`
	Type              *CodeableConcept   `bson:"type" json:"type"`
	Name              *string            `bson:"name" json:"name"`
	Subject           *Reference         `bson:"subject" json:"subject"`
	Period            *Period            `bson:"period" json:"period"`
	Active            *Period            `bson:"active" json:"active"`
	Balance           *Money             `bson:"balance" json:"balance"`
	Coverage          []AccountCoverage  `bson:"coverage" json:"coverage"`
	Owner             *Reference         `bson:"owner" json:"owner"`
	Description       *string            `bson:"description" json:"description"`
	Guarantor         []AccountGuarantor `bson:"guarantor" json:"guarantor"`
}
type AccountCoverage struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Coverage          Reference   `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Priority          *int        `bson:"priority" json:"priority"`
}
type AccountGuarantor struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Party             Reference   `bson:"party,omitempty" json:"party,omitempty"`
	OnHold            *bool       `bson:"onHold" json:"onHold"`
	Period            *Period     `bson:"period" json:"period"`
}

// OtherAccount is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAccount Account

// MarshalJSON marshals the given Account as JSON into a byte slice
func (r Account) MarshalJSON() ([]byte, error) {
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
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Account
func (r *Account) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAccount)(r)); err != nil {
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
func (r Account) GetResourceType() ResourceType {
	return ResourceTypeAccount
}
