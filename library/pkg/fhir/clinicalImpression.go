package fhir

import "encoding/json"

// ClinicalImpression is documented here http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                     *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules            *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                 *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                     *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension                []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier               []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                   string                            `bson:"status" json:"status"`
	Code                     *CodeableConcept                  `bson:"code,omitempty" json:"code,omitempty"`
	Description              *string                           `bson:"description,omitempty" json:"description,omitempty"`
	Date                     *string                           `bson:"date,omitempty" json:"date,omitempty"`
	Assessor                 *Reference                        `bson:"assessor,omitempty" json:"assessor,omitempty"`
	Previous                 *Reference                        `bson:"previous,omitempty" json:"previous,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `bson:"investigation,omitempty" json:"investigation,omitempty"`
	Protocol                 []string                          `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                  *string                           `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `bson:"finding,omitempty" json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	Note                     []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
}
type ClinicalImpressionInvestigation struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
}
type ClinicalImpressionFinding struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Basis             *string     `bson:"basis,omitempty" json:"basis,omitempty"`
}
type OtherClinicalImpression ClinicalImpression

// MarshalJSON marshals the given ClinicalImpression as JSON into a byte slice
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
}

// UnmarshalClinicalImpression unmarshals a ClinicalImpression.
func UnmarshalClinicalImpression(b []byte) (ClinicalImpression, error) {
	var clinicalImpression ClinicalImpression
	if err := json.Unmarshal(b, &clinicalImpression); err != nil {
		return clinicalImpression, err
	}
	return clinicalImpression, nil
}
