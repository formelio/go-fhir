package fhir

import "encoding/json"

// NutritionOrder is documented here http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrder struct {
	Id                     *string                       `bson:"id" json:"id"`
	Meta                   *Meta                         `bson:"meta" json:"meta"`
	ImplicitRules          *string                       `bson:"implicitRules" json:"implicitRules"`
	Language               *string                       `bson:"language" json:"language"`
	Text                   *Narrative                    `bson:"text" json:"text"`
	RawContained           []json.RawMessage             `bson:"contained" json:"contained"`
	Contained              []IResource                   `bson:"-" json:"-"`
	Extension              []Extension                   `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                   `bson:"modifierExtension" json:"modifierExtension"`
	Identifier             []Identifier                  `bson:"identifier" json:"identifier"`
	Status                 *NutritionOrderStatus         `bson:"status" json:"status"`
	Patient                Reference                     `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter              *Reference                    `bson:"encounter" json:"encounter"`
	DateTime               string                        `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Orderer                *Reference                    `bson:"orderer" json:"orderer"`
	AllergyIntolerance     []Reference                   `bson:"allergyIntolerance" json:"allergyIntolerance"`
	FoodPreferenceModifier []CodeableConcept             `bson:"foodPreferenceModifier" json:"foodPreferenceModifier"`
	ExcludeFoodModifier    []CodeableConcept             `bson:"excludeFoodModifier" json:"excludeFoodModifier"`
	OralDiet               *NutritionOrderOralDiet       `bson:"oralDiet" json:"oralDiet"`
	Supplement             []NutritionOrderSupplement    `bson:"supplement" json:"supplement"`
	EnteralFormula         *NutritionOrderEnteralFormula `bson:"enteralFormula" json:"enteralFormula"`
}
type NutritionOrderOralDiet struct {
	Id                   *string                          `bson:"id" json:"id"`
	Extension            []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Type                 []CodeableConcept                `bson:"type" json:"type"`
	Schedule             []Timing                         `bson:"schedule" json:"schedule"`
	Nutrient             []NutritionOrderOralDietNutrient `bson:"nutrient" json:"nutrient"`
	Texture              []NutritionOrderOralDietTexture  `bson:"texture" json:"texture"`
	FluidConsistencyType []CodeableConcept                `bson:"fluidConsistencyType" json:"fluidConsistencyType"`
	Instruction          *string                          `bson:"instruction" json:"instruction"`
}
type NutritionOrderOralDietNutrient struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Modifier          *CodeableConcept `bson:"modifier" json:"modifier"`
	Amount            *Quantity        `bson:"amount" json:"amount"`
}
type NutritionOrderOralDietTexture struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Modifier          *CodeableConcept `bson:"modifier" json:"modifier"`
	FoodType          *CodeableConcept `bson:"foodType" json:"foodType"`
}
type NutritionOrderSupplement struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	ProductName       *string          `bson:"productName" json:"productName"`
	Schedule          []Timing         `bson:"schedule" json:"schedule"`
	Quantity          *Quantity        `bson:"quantity" json:"quantity"`
	Instruction       *string          `bson:"instruction" json:"instruction"`
}
type NutritionOrderEnteralFormula struct {
	Id                        *string                                      `bson:"id" json:"id"`
	Extension                 []Extension                                  `bson:"extension" json:"extension"`
	ModifierExtension         []Extension                                  `bson:"modifierExtension" json:"modifierExtension"`
	BaseFormulaType           *CodeableConcept                             `bson:"baseFormulaType" json:"baseFormulaType"`
	BaseFormulaProductName    *string                                      `bson:"baseFormulaProductName" json:"baseFormulaProductName"`
	AdditiveType              *CodeableConcept                             `bson:"additiveType" json:"additiveType"`
	AdditiveProductName       *string                                      `bson:"additiveProductName" json:"additiveProductName"`
	CaloricDensity            *Quantity                                    `bson:"caloricDensity" json:"caloricDensity"`
	RouteofAdministration     *CodeableConcept                             `bson:"routeofAdministration" json:"routeofAdministration"`
	Administration            []NutritionOrderEnteralFormulaAdministration `bson:"administration" json:"administration"`
	MaxVolumeToDeliver        *Quantity                                    `bson:"maxVolumeToDeliver" json:"maxVolumeToDeliver"`
	AdministrationInstruction *string                                      `bson:"administrationInstruction" json:"administrationInstruction"`
}
type NutritionOrderEnteralFormulaAdministration struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Schedule          *Timing     `bson:"schedule" json:"schedule"`
	Quantity          *Quantity   `bson:"quantity" json:"quantity"`
	RateQuantity      *Quantity   `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
	RateRatio         *Ratio      `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
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
	return json.Marshal(struct {
		OtherNutritionOrder
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionOrder: OtherNutritionOrder(r),
		ResourceType:        "NutritionOrder",
	})
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
