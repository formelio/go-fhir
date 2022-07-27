package fhir

import "encoding/json"

// Substance is documented here http://hl7.org/fhir/StructureDefinition/Substance
type Substance struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
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

// OtherSubstance is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSubstance Substance

// MarshalJSON marshals the given Substance as JSON into a byte slice
func (r Substance) MarshalJSON() ([]byte, error) {
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
		OtherSubstance
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstance: OtherSubstance(r),
		ResourceType:   "Substance",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Substance
func (r *Substance) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSubstance)(r)); err != nil {
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
func (r Substance) GetResourceType() ResourceType {
	return ResourceTypeSubstance
}
