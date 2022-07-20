package fhir

import "encoding/json"

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string            `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
	Period            *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Active            *Period            `bson:"active,omitempty" json:"active,omitempty"`
	Balance           *Money             `bson:"balance,omitempty" json:"balance,omitempty"`
	Coverage          []AccountCoverage  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Owner             *Reference         `bson:"owner,omitempty" json:"owner,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	Guarantor         []AccountGuarantor `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
}
type AccountCoverage struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          Reference   `bson:"coverage" json:"coverage"`
	Priority          *int        `bson:"priority,omitempty" json:"priority,omitempty"`
}
type AccountGuarantor struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnHold            *bool       `bson:"onHold,omitempty" json:"onHold,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherAccount Account

// MarshalJSON marshals the given Account as JSON into a byte slice
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}

// UnmarshalAccount unmarshals a Account.
func UnmarshalAccount(b []byte) (Account, error) {
	var account Account
	if err := json.Unmarshal(b, &account); err != nil {
		return account, err
	}
	return account, nil
}
