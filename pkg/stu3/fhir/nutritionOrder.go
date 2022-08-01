package fhir

import (
	"bytes"
	"encoding/json"
)

// NutritionOrder is documented here http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrder struct {
	Id                     *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage             `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource                   `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []*Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []*Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []*Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                 *NutritionOrderStatus         `bson:"status,omitempty" json:"status,omitempty"`
	Patient                Reference                     `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter              *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               string                        `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Orderer                *Reference                    `bson:"orderer,omitempty" json:"orderer,omitempty"`
	AllergyIntolerance     []*Reference                  `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []*CodeableConcept            `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []*CodeableConcept            `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []*NutritionOrderSupplement   `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
}
type NutritionOrderOralDiet struct {
	Id                   *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []*Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 []*CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []*Timing                         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []*NutritionOrderOralDietNutrient `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []*NutritionOrderOralDietTexture  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []*CodeableConcept                `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          *string                           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderOralDietNutrient struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount            *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}
type NutritionOrderOralDietTexture struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType          *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}
type NutritionOrderSupplement struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName       *string          `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule          []*Timing        `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction       *string          `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderEnteralFormula struct {
	Id                        *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []*Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []*Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	BaseFormulaType           *CodeableConcept                              `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    *string                                       `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                              `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       *string                                       `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                     `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                              `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []*NutritionOrderEnteralFormulaAdministration `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                     `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction *string                                       `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}
type NutritionOrderEnteralFormulaAdministration struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Schedule          *Timing      `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity          *Quantity    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateQuantity      *Quantity    `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
	RateRatio         *Ratio       `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
}

// OtherNutritionOrder is a helper type to use the default implementations of Marshall and Unmarshal
type OtherNutritionOrder NutritionOrder

// MarshalJSON marshals the given NutritionOrder as JSON into a byte slice
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
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
		OtherNutritionOrder
	}{
		OtherNutritionOrder: OtherNutritionOrder(r),
		ResourceType:        "NutritionOrder",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into NutritionOrder
func (r *NutritionOrder) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherNutritionOrder)(r)); err != nil {
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
func (r NutritionOrder) GetResourceType() ResourceType {
	return ResourceTypeNutritionOrder
}
