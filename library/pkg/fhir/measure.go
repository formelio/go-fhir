package fhir

import "encoding/json"

// Measure is documented here http://hl7.org/fhir/StructureDefinition/Measure
type Measure struct {
	Id                              *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                            *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules                   *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                        *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                            *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension                       []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension               []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                             *string                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier                      []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                         *string                   `bson:"version,omitempty" json:"version,omitempty"`
	Name                            *string                   `bson:"name,omitempty" json:"name,omitempty"`
	Title                           *string                   `bson:"title,omitempty" json:"title,omitempty"`
	Status                          string                    `bson:"status" json:"status"`
	Experimental                    *bool                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                            *string                   `bson:"date,omitempty" json:"date,omitempty"`
	Publisher                       *string                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description                     *string                   `bson:"description,omitempty" json:"description,omitempty"`
	Purpose                         *string                   `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage                           *string                   `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate                    *string                   `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate                  *string                   `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod                 *Period                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext                      []UsageContext            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept         `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic                           []CodeableConcept         `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor                     []Contributor             `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact                         []ContactDetail           `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright                       *string                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact                 []RelatedArtifact         `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library                         []Reference               `bson:"library,omitempty" json:"library,omitempty"`
	Disclaimer                      *string                   `bson:"disclaimer,omitempty" json:"disclaimer,omitempty"`
	Scoring                         *CodeableConcept          `bson:"scoring,omitempty" json:"scoring,omitempty"`
	CompositeScoring                *CodeableConcept          `bson:"compositeScoring,omitempty" json:"compositeScoring,omitempty"`
	Type                            []CodeableConcept         `bson:"type,omitempty" json:"type,omitempty"`
	RiskAdjustment                  *string                   `bson:"riskAdjustment,omitempty" json:"riskAdjustment,omitempty"`
	RateAggregation                 *string                   `bson:"rateAggregation,omitempty" json:"rateAggregation,omitempty"`
	Rationale                       *string                   `bson:"rationale,omitempty" json:"rationale,omitempty"`
	ClinicalRecommendationStatement *string                   `bson:"clinicalRecommendationStatement,omitempty" json:"clinicalRecommendationStatement,omitempty"`
	ImprovementNotation             *string                   `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	Definition                      []string                  `bson:"definition,omitempty" json:"definition,omitempty"`
	Guidance                        *string                   `bson:"guidance,omitempty" json:"guidance,omitempty"`
	Set                             *string                   `bson:"set,omitempty" json:"set,omitempty"`
	Group                           []MeasureGroup            `bson:"group,omitempty" json:"group,omitempty"`
	SupplementalData                []MeasureSupplementalData `bson:"supplementalData,omitempty" json:"supplementalData,omitempty"`
}
type MeasureGroup struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        Identifier               `bson:"identifier" json:"identifier"`
	Name              *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                  `bson:"description,omitempty" json:"description,omitempty"`
	Population        []MeasureGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	Stratifier        []MeasureGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}
type MeasureGroupPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Name              *string          `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Criteria          string           `bson:"criteria" json:"criteria"`
}
type MeasureGroupStratifier struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Criteria          *string     `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
}
type MeasureSupplementalData struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Usage             []CodeableConcept `bson:"usage,omitempty" json:"usage,omitempty"`
	Criteria          *string           `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Path              *string           `bson:"path,omitempty" json:"path,omitempty"`
}
type OtherMeasure Measure

// MarshalJSON marshals the given Measure as JSON into a byte slice
func (r Measure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasure
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasure: OtherMeasure(r),
		ResourceType: "Measure",
	})
}

// UnmarshalMeasure unmarshals a Measure.
func UnmarshalMeasure(b []byte) (Measure, error) {
	var measure Measure
	if err := json.Unmarshal(b, &measure); err != nil {
		return measure, err
	}
	return measure, nil
}
