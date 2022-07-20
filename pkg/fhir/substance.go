package fhir

import "encoding/json"

// Substance is documented here http://hl7.org/fhir/StructureDefinition/Substance
type Substance struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Status            *FHIRSubstanceStatus  `bson:"status" json:"status"`
	Category          []CodeableConcept     `bson:"category" json:"category"`
	Code              CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Description       *string               `bson:"description" json:"description"`
	Instance          []SubstanceInstance   `bson:"instance" json:"instance"`
	Ingredient        []SubstanceIngredient `bson:"ingredient" json:"ingredient"`
}
type SubstanceInstance struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier `bson:"identifier" json:"identifier"`
	Expiry            *string     `bson:"expiry" json:"expiry"`
	Quantity          *Quantity   `bson:"quantity" json:"quantity"`
}
type SubstanceIngredient struct {
	Id                       *string          `bson:"id" json:"id"`
	Extension                []Extension      `bson:"extension" json:"extension"`
	ModifierExtension        []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Quantity                 *Ratio           `bson:"quantity" json:"quantity"`
	SubstanceCodeableConcept *CodeableConcept `bson:"substanceCodeableConcept,omitempty" json:"substanceCodeableConcept,omitempty"`
	SubstanceReference       *Reference       `bson:"substanceReference,omitempty" json:"substanceReference,omitempty"`
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

// UnmarshalSubstance unmarshalls a Substance.
func UnmarshalSubstance(b []byte) (Substance, error) {
	var substance Substance
	if err := json.Unmarshal(b, &substance); err != nil {
		return substance, err
	}
	return substance, nil
}
