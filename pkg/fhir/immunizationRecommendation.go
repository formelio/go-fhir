package fhir

import "encoding/json"

// ImmunizationRecommendation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient           Reference                                  `bson:"patient" json:"patient"`
	Recommendation    []ImmunizationRecommendationRecommendation `bson:"recommendation" json:"recommendation"`
}
type ImmunizationRecommendationRecommendation struct {
	Id                     *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date                   string                                                  `bson:"date" json:"date"`
	VaccineCode            *CodeableConcept                                        `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	TargetDisease          *CodeableConcept                                        `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseNumber             *int                                                    `bson:"doseNumber,omitempty" json:"doseNumber,omitempty"`
	ForecastStatus         CodeableConcept                                         `bson:"forecastStatus" json:"forecastStatus"`
	DateCriterion          []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	Protocol               *ImmunizationRecommendationRecommendationProtocol       `bson:"protocol,omitempty" json:"protocol,omitempty"`
	SupportingImmunization []Reference                                             `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
}
type ImmunizationRecommendationRecommendationDateCriterion struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Value             string          `bson:"value" json:"value"`
}
type ImmunizationRecommendationRecommendationProtocol struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DoseSequence      *int        `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Authority         *Reference  `bson:"authority,omitempty" json:"authority,omitempty"`
	Series            *string     `bson:"series,omitempty" json:"series,omitempty"`
}
type OtherImmunizationRecommendation ImmunizationRecommendation

// MarshalJSON marshals the given ImmunizationRecommendation as JSON into a byte slice
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationRecommendation: OtherImmunizationRecommendation(r),
		ResourceType:                    "ImmunizationRecommendation",
	})
}

// UnmarshalImmunizationRecommendation unmarshals a ImmunizationRecommendation.
func UnmarshalImmunizationRecommendation(b []byte) (ImmunizationRecommendation, error) {
	var immunizationRecommendation ImmunizationRecommendation
	if err := json.Unmarshal(b, &immunizationRecommendation); err != nil {
		return immunizationRecommendation, err
	}
	return immunizationRecommendation, nil
}
