package fhir

import (
	"bytes"
	"encoding/json"
)

// FamilyMemberHistory is documented here http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition        []*Reference                    `bson:"definition,omitempty" json:"definition,omitempty"`
	Status            FamilyHistoryStatus             `bson:"status,omitempty" json:"status,omitempty"`
	NotDone           *bool                           `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason     *CodeableConcept                `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Patient           Reference                       `bson:"patient,omitempty" json:"patient,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Name              *string                         `bson:"name,omitempty" json:"name,omitempty"`
	Relationship      CodeableConcept                 `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Gender            *AdministrativeGender           `bson:"gender,omitempty" json:"gender,omitempty"`
	BornPeriod        *Period                         `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate          *string                         `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString        *string                         `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge            *Age                            `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange          *Range                          `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString         *string                         `bson:"ageString,omitempty" json:"ageString,omitempty"`
	EstimatedAge      *bool                           `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	DeceasedBoolean   *bool                           `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge       *Age                            `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange     *Range                          `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate      *string                         `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString    *string                         `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	ReasonCode        []*CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []*Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note              []*Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	Condition         []*FamilyMemberHistoryCondition `bson:"condition,omitempty" json:"condition,omitempty"`
}
type FamilyMemberHistoryCondition struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Outcome           *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OnsetAge          *Age             `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange        *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetPeriod       *Period          `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString       *string          `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note              []*Annotation    `bson:"note,omitempty" json:"note,omitempty"`
}

// OtherFamilyMemberHistory is a helper type to use the default implementations of Marshall and Unmarshal
type OtherFamilyMemberHistory FamilyMemberHistory

// MarshalJSON marshals the given FamilyMemberHistory as JSON into a byte slice
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherFamilyMemberHistory
	}{
		OtherFamilyMemberHistory: OtherFamilyMemberHistory(r),
		ResourceType:             "FamilyMemberHistory",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into FamilyMemberHistory
func (r *FamilyMemberHistory) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherFamilyMemberHistory)(r)); err != nil {
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
func (r FamilyMemberHistory) GetResourceType() ResourceType {
	return ResourceTypeFamilyMemberHistory
}
