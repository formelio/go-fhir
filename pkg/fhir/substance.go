package fhir

import "encoding/json"

// Substance is documented here http://hl7.org/fhir/StructureDefinition/Substance
type Substance struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string               `bson:"status,omitempty" json:"status,omitempty"`
	Category          []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept       `bson:"code" json:"code"`
	Description       *string               `bson:"description,omitempty" json:"description,omitempty"`
	Instance          []SubstanceInstance   `bson:"instance,omitempty" json:"instance,omitempty"`
	Ingredient        []SubstanceIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
}
type SubstanceInstance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Expiry            *string     `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
type SubstanceIngredient struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          *Ratio      `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
type OtherSubstance Substance

// MarshalJSON marshals the given Substance as JSON into a byte slice
func (r Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstance
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstance: OtherSubstance(r),
		ResourceType:   "Substance",
	})
}

// UnmarshalSubstance unmarshals a Substance.
func UnmarshalSubstance(b []byte) (Substance, error) {
	var substance Substance
	if err := json.Unmarshal(b, &substance); err != nil {
		return substance, err
	}
	return substance, nil
}
