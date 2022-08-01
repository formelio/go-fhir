package fhir

import "encoding/json"

// ImmunizationRecommendation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage                          `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                                `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient           Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Recommendation    []ImmunizationRecommendationRecommendation `bson:"recommendation,omitempty" json:"recommendation,omitempty"`
}
type ImmunizationRecommendationRecommendation struct {
	Id                           *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension                    []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date                         string                                                  `bson:"date,omitempty" json:"date,omitempty"`
	VaccineCode                  *CodeableConcept                                        `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	TargetDisease                *CodeableConcept                                        `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseNumber                   *int                                                    `bson:"doseNumber,omitempty" json:"doseNumber,omitempty"`
	ForecastStatus               CodeableConcept                                         `bson:"forecastStatus,omitempty" json:"forecastStatus,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	Protocol                     *ImmunizationRecommendationRecommendationProtocol       `bson:"protocol,omitempty" json:"protocol,omitempty"`
	SupportingImmunization       []Reference                                             `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                             `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty"`
}
type ImmunizationRecommendationRecommendationDateCriterion struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Value             string          `bson:"value,omitempty" json:"value,omitempty"`
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

// OtherImmunizationRecommendation is a helper type to use the default implementations of Marshall and Unmarshal
type OtherImmunizationRecommendation ImmunizationRecommendation

// MarshalJSON marshals the given ImmunizationRecommendation as JSON into a byte slice
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
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
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationRecommendation: OtherImmunizationRecommendation(r),
		ResourceType:                    "ImmunizationRecommendation",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ImmunizationRecommendation
func (r *ImmunizationRecommendation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherImmunizationRecommendation)(r)); err != nil {
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
func (r ImmunizationRecommendation) GetResourceType() ResourceType {
	return ResourceTypeImmunizationRecommendation
}
