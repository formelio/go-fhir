package fhir

import "encoding/json"

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
type Account struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	Contained         []json.RawMessage  `bson:"contained" json:"contained"`
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

// UnmarshalAccount unmarshalls a Account.
func UnmarshalAccount(b []byte) (Account, error) {
	var account Account
	if err := json.Unmarshal(b, &account); err != nil {
		return account, err
	}
	return account, nil
}
