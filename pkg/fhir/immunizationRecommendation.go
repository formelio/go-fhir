package fhir

import "encoding/json"

// ImmunizationRecommendation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
	Id                *string                                    `bson:"id" json:"id"`
	Meta              *Meta                                      `bson:"meta" json:"meta"`
	ImplicitRules     *string                                    `bson:"implicitRules" json:"implicitRules"`
	Language          *string                                    `bson:"language" json:"language"`
	Text              *Narrative                                 `bson:"text" json:"text"`
	RawContained      []json.RawMessage                          `bson:"contained" json:"contained"`
	Contained         []IResource                                `bson:"-" json:"-"`
	Extension         []Extension                                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier                               `bson:"identifier" json:"identifier"`
	Patient           Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Recommendation    []ImmunizationRecommendationRecommendation `bson:"recommendation,omitempty" json:"recommendation,omitempty"`
}
type ImmunizationRecommendationRecommendation struct {
	Id                           *string                                                 `bson:"id" json:"id"`
	Extension                    []Extension                                             `bson:"extension" json:"extension"`
	ModifierExtension            []Extension                                             `bson:"modifierExtension" json:"modifierExtension"`
	Date                         string                                                  `bson:"date,omitempty" json:"date,omitempty"`
	VaccineCode                  *CodeableConcept                                        `bson:"vaccineCode" json:"vaccineCode"`
	TargetDisease                *CodeableConcept                                        `bson:"targetDisease" json:"targetDisease"`
	DoseNumber                   *int                                                    `bson:"doseNumber" json:"doseNumber"`
	ForecastStatus               CodeableConcept                                         `bson:"forecastStatus,omitempty" json:"forecastStatus,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion" json:"dateCriterion"`
	Protocol                     *ImmunizationRecommendationRecommendationProtocol       `bson:"protocol" json:"protocol"`
	SupportingImmunization       []Reference                                             `bson:"supportingImmunization" json:"supportingImmunization"`
	SupportingPatientInformation []Reference                                             `bson:"supportingPatientInformation" json:"supportingPatientInformation"`
}
type ImmunizationRecommendationRecommendationDateCriterion struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Value             string          `bson:"value,omitempty" json:"value,omitempty"`
}
type ImmunizationRecommendationRecommendationProtocol struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	DoseSequence      *int        `bson:"doseSequence" json:"doseSequence"`
	Description       *string     `bson:"description" json:"description"`
	Authority         *Reference  `bson:"authority" json:"authority"`
	Series            *string     `bson:"series" json:"series"`
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
