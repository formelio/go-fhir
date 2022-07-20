package fhir

import "encoding/json"

// FamilyMemberHistory is documented here http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                *string                        `bson:"id" json:"id"`
	Meta              *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules     *string                        `bson:"implicitRules" json:"implicitRules"`
	Language          *string                        `bson:"language" json:"language"`
	Text              *Narrative                     `bson:"text" json:"text"`
	Contained         []json.RawMessage              `bson:"contained" json:"contained"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier                   `bson:"identifier" json:"identifier"`
	Definition        []Reference                    `bson:"definition" json:"definition"`
	Status            FamilyHistoryStatus            `bson:"status,omitempty" json:"status,omitempty"`
	NotDone           *bool                          `bson:"notDone" json:"notDone"`
	NotDoneReason     *CodeableConcept               `bson:"notDoneReason" json:"notDoneReason"`
	Patient           Reference                      `bson:"patient,omitempty" json:"patient,omitempty"`
	Date              *string                        `bson:"date" json:"date"`
	Name              *string                        `bson:"name" json:"name"`
	Relationship      CodeableConcept                `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Gender            *AdministrativeGender          `bson:"gender" json:"gender"`
	BornPeriod        *Period                        `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate          *string                        `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString        *string                        `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge            *Age                           `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange          *Range                         `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString         *string                        `bson:"ageString,omitempty" json:"ageString,omitempty"`
	EstimatedAge      *bool                          `bson:"estimatedAge" json:"estimatedAge"`
	DeceasedBoolean   *bool                          `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge       *Age                           `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange     *Range                         `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate      *string                        `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString    *string                        `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	ReasonCode        []CodeableConcept              `bson:"reasonCode" json:"reasonCode"`
	ReasonReference   []Reference                    `bson:"reasonReference" json:"reasonReference"`
	Note              []Annotation                   `bson:"note" json:"note"`
	Condition         []FamilyMemberHistoryCondition `bson:"condition" json:"condition"`
}
type FamilyMemberHistoryCondition struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Code              CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Outcome           *CodeableConcept `bson:"outcome" json:"outcome"`
	OnsetAge          *Age             `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange        *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetPeriod       *Period          `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString       *string          `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note              []Annotation     `bson:"note" json:"note"`
}
type OtherFamilyMemberHistory FamilyMemberHistory

// MarshalJSON marshals the given FamilyMemberHistory as JSON into a byte slice
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherFamilyMemberHistory: OtherFamilyMemberHistory(r),
		ResourceType:             "FamilyMemberHistory",
	})
}

// UnmarshalFamilyMemberHistory unmarshalls a FamilyMemberHistory.
func UnmarshalFamilyMemberHistory(b []byte) (FamilyMemberHistory, error) {
	var familyMemberHistory FamilyMemberHistory
	if err := json.Unmarshal(b, &familyMemberHistory); err != nil {
		return familyMemberHistory, err
	}
	return familyMemberHistory, nil
}
