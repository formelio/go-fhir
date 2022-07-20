package fhir

import "encoding/json"

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                              `bson:"id" json:"id"`
	Meta               *Meta                                `bson:"meta" json:"meta"`
	ImplicitRules      *string                              `bson:"implicitRules" json:"implicitRules"`
	Language           *string                              `bson:"language" json:"language"`
	Text               *Narrative                           `bson:"text" json:"text"`
	Contained          []json.RawMessage                    `bson:"contained" json:"contained"`
	Extension          []Extension                          `bson:"extension" json:"extension"`
	ModifierExtension  []Extension                          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier                         `bson:"identifier" json:"identifier"`
	ClinicalStatus     *AllergyIntoleranceClinicalStatus    `bson:"clinicalStatus" json:"clinicalStatus"`
	VerificationStatus AllergyIntoleranceVerificationStatus `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               *AllergyIntoleranceType              `bson:"type" json:"type"`
	Category           []AllergyIntoleranceCategory         `bson:"category" json:"category"`
	Criticality        *AllergyIntoleranceCriticality       `bson:"criticality" json:"criticality"`
	Code               *CodeableConcept                     `bson:"code" json:"code"`
	Patient            Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	OnsetDateTime      *string                              `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                                 `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                              `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                               `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        *string                              `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AssertedDate       *string                              `bson:"assertedDate" json:"assertedDate"`
	Recorder           *Reference                           `bson:"recorder" json:"recorder"`
	Asserter           *Reference                           `bson:"asserter" json:"asserter"`
	LastOccurrence     *string                              `bson:"lastOccurrence" json:"lastOccurrence"`
	Note               []Annotation                         `bson:"note" json:"note"`
	Reaction           []AllergyIntoleranceReaction         `bson:"reaction" json:"reaction"`
}
type AllergyIntoleranceReaction struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Substance         *CodeableConcept            `bson:"substance" json:"substance"`
	Manifestation     []CodeableConcept           `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description       *string                     `bson:"description" json:"description"`
	Onset             *string                     `bson:"onset" json:"onset"`
	Severity          *AllergyIntoleranceSeverity `bson:"severity" json:"severity"`
	ExposureRoute     *CodeableConcept            `bson:"exposureRoute" json:"exposureRoute"`
	Note              []Annotation                `bson:"note" json:"note"`
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

// UnmarshalAllergyIntolerance unmarshalls a AllergyIntolerance.
func UnmarshalAllergyIntolerance(b []byte) (AllergyIntolerance, error) {
	var allergyIntolerance AllergyIntolerance
	if err := json.Unmarshal(b, &allergyIntolerance); err != nil {
		return allergyIntolerance, err
	}
	return allergyIntolerance, nil
}
