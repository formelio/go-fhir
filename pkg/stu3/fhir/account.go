package fhir

import (
	"bytes"
	"encoding/json"
)

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource         `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *AccountStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept    `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string             `bson:"name,omitempty" json:"name,omitempty"`
	Subject           *Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Period            *Period             `bson:"period,omitempty" json:"period,omitempty"`
	Active            *Period             `bson:"active,omitempty" json:"active,omitempty"`
	Balance           *Money              `bson:"balance,omitempty" json:"balance,omitempty"`
	Coverage          []*AccountCoverage  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Owner             *Reference          `bson:"owner,omitempty" json:"owner,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	Guarantor         []*AccountGuarantor `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
}
type AccountCoverage struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          Reference    `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Priority          *int         `bson:"priority,omitempty" json:"priority,omitempty"`
}
type AccountGuarantor struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Party             Reference    `bson:"party,omitempty" json:"party,omitempty"`
	OnHold            *bool        `bson:"onHold,omitempty" json:"onHold,omitempty"`
	Period            *Period      `bson:"period,omitempty" json:"period,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherAccount
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
	return buffer.Bytes(), err
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
