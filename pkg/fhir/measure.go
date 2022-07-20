package fhir

import "encoding/json"

// Measure is documented here http://hl7.org/fhir/StructureDefinition/Measure
type Measure struct {
	Id                              *string                   `bson:"id" json:"id"`
	Meta                            *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules                   *string                   `bson:"implicitRules" json:"implicitRules"`
	Language                        *string                   `bson:"language" json:"language"`
	Text                            *Narrative                `bson:"text" json:"text"`
	Contained                       []json.RawMessage         `bson:"contained" json:"contained"`
	Extension                       []Extension               `bson:"extension" json:"extension"`
	ModifierExtension               []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Url                             *string                   `bson:"url" json:"url"`
	Identifier                      []Identifier              `bson:"identifier" json:"identifier"`
	Version                         *string                   `bson:"version" json:"version"`
	Name                            *string                   `bson:"name" json:"name"`
	Title                           *string                   `bson:"title" json:"title"`
	Status                          PublicationStatus         `bson:"status,omitempty" json:"status,omitempty"`
	Experimental                    *bool                     `bson:"experimental" json:"experimental"`
	Date                            *string                   `bson:"date" json:"date"`
	Publisher                       *string                   `bson:"publisher" json:"publisher"`
	Description                     *string                   `bson:"description" json:"description"`
	Purpose                         *string                   `bson:"purpose" json:"purpose"`
	Usage                           *string                   `bson:"usage" json:"usage"`
	ApprovalDate                    *string                   `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate                  *string                   `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod                 *Period                   `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext                      []UsageContext            `bson:"useContext" json:"useContext"`
	Jurisdiction                    []CodeableConcept         `bson:"jurisdiction" json:"jurisdiction"`
	Topic                           []CodeableConcept         `bson:"topic" json:"topic"`
	Contributor                     []Contributor             `bson:"contributor" json:"contributor"`
	Contact                         []ContactDetail           `bson:"contact" json:"contact"`
	Copyright                       *string                   `bson:"copyright" json:"copyright"`
	RelatedArtifact                 []RelatedArtifact         `bson:"relatedArtifact" json:"relatedArtifact"`
	Library                         []Reference               `bson:"library" json:"library"`
	Disclaimer                      *string                   `bson:"disclaimer" json:"disclaimer"`
	Scoring                         *CodeableConcept          `bson:"scoring" json:"scoring"`
	CompositeScoring                *CodeableConcept          `bson:"compositeScoring" json:"compositeScoring"`
	Type                            []CodeableConcept         `bson:"type" json:"type"`
	RiskAdjustment                  *string                   `bson:"riskAdjustment" json:"riskAdjustment"`
	RateAggregation                 *string                   `bson:"rateAggregation" json:"rateAggregation"`
	Rationale                       *string                   `bson:"rationale" json:"rationale"`
	ClinicalRecommendationStatement *string                   `bson:"clinicalRecommendationStatement" json:"clinicalRecommendationStatement"`
	ImprovementNotation             *string                   `bson:"improvementNotation" json:"improvementNotation"`
	Definition                      []string                  `bson:"definition" json:"definition"`
	Guidance                        *string                   `bson:"guidance" json:"guidance"`
	Set                             *string                   `bson:"set" json:"set"`
	Group                           []MeasureGroup            `bson:"group" json:"group"`
	SupplementalData                []MeasureSupplementalData `bson:"supplementalData" json:"supplementalData"`
}
type MeasureGroup struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string                  `bson:"name" json:"name"`
	Description       *string                  `bson:"description" json:"description"`
	Population        []MeasureGroupPopulation `bson:"population" json:"population"`
	Stratifier        []MeasureGroupStratifier `bson:"stratifier" json:"stratifier"`
}
type MeasureGroupPopulation struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier      `bson:"identifier" json:"identifier"`
	Code              *CodeableConcept `bson:"code" json:"code"`
	Name              *string          `bson:"name" json:"name"`
	Description       *string          `bson:"description" json:"description"`
	Criteria          string           `bson:"criteria,omitempty" json:"criteria,omitempty"`
}
type MeasureGroupStratifier struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier `bson:"identifier" json:"identifier"`
	Criteria          *string     `bson:"criteria" json:"criteria"`
	Path              *string     `bson:"path" json:"path"`
}
type MeasureSupplementalData struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier       `bson:"identifier" json:"identifier"`
	Usage             []CodeableConcept `bson:"usage" json:"usage"`
	Criteria          *string           `bson:"criteria" json:"criteria"`
	Path              *string           `bson:"path" json:"path"`
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

// UnmarshalMeasure unmarshalls a Measure.
func UnmarshalMeasure(b []byte) (Measure, error) {
	var measure Measure
	if err := json.Unmarshal(b, &measure); err != nil {
		return measure, err
	}
	return measure, nil
}
