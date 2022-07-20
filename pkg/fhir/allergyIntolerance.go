package fhir

import "encoding/json"

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *string                      `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus string                       `bson:"verificationStatus" json:"verificationStatus"`
	Type               *string                      `bson:"type,omitempty" json:"type,omitempty"`
	Category           []string                     `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        *string                      `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Patient            Reference                    `bson:"patient" json:"patient"`
	AssertedDate       *string                      `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	LastOccurrence     *string                      `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                 `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction `bson:"reaction,omitempty" json:"reaction,omitempty"`
}
type AllergyIntoleranceReaction struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Manifestation     []CodeableConcept `bson:"manifestation" json:"manifestation"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	Onset             *string           `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity          *string           `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherAllergyIntolerance AllergyIntolerance

// MarshalJSON marshals the given AllergyIntolerance as JSON into a byte slice
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

// UnmarshalAllergyIntolerance unmarshals a AllergyIntolerance.
func UnmarshalAllergyIntolerance(b []byte) (AllergyIntolerance, error) {
	var allergyIntolerance AllergyIntolerance
	if err := json.Unmarshal(b, &allergyIntolerance); err != nil {
		return allergyIntolerance, err
	}
	return allergyIntolerance, nil
}
