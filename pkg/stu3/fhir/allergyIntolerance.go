package fhir

import "encoding/json"

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                              `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage                    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource                          `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *AllergyIntoleranceClinicalStatus    `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus AllergyIntoleranceVerificationStatus `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               *AllergyIntoleranceType              `bson:"type,omitempty" json:"type,omitempty"`
	Category           []AllergyIntoleranceCategory         `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        *AllergyIntoleranceCriticality       `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Patient            Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	OnsetDateTime      *string                              `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                                 `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                              `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                               `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        *string                              `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AssertedDate       *string                              `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Recorder           *Reference                           `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference                           `bson:"asserter,omitempty" json:"asserter,omitempty"`
	LastOccurrence     *string                              `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                         `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction         `bson:"reaction,omitempty" json:"reaction,omitempty"`
}
type AllergyIntoleranceReaction struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept            `bson:"substance,omitempty" json:"substance,omitempty"`
	Manifestation     []CodeableConcept           `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Onset             *string                     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity          *AllergyIntoleranceSeverity `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept            `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note              []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}

// OtherAllergyIntolerance is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAllergyIntolerance AllergyIntolerance

// MarshalJSON marshals the given AllergyIntolerance as JSON into a byte slice
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
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
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

// UnmarshalJSON unmarshals the given byte slice into AllergyIntolerance
func (r *AllergyIntolerance) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAllergyIntolerance)(r)); err != nil {
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
func (r AllergyIntolerance) GetResourceType() ResourceType {
	return ResourceTypeAllergyIntolerance
}
