package fhir

import "encoding/json"

// FamilyMemberHistory is documented here http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                         `bson:"status" json:"status"`
	NotDone           *bool                          `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason     *CodeableConcept               `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Patient           Reference                      `bson:"patient" json:"patient"`
	Date              *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Name              *string                        `bson:"name,omitempty" json:"name,omitempty"`
	Relationship      CodeableConcept                `bson:"relationship" json:"relationship"`
	Gender            *string                        `bson:"gender,omitempty" json:"gender,omitempty"`
	EstimatedAge      *bool                          `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	ReasonCode        []CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Note              []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	Condition         []FamilyMemberHistoryCondition `bson:"condition,omitempty" json:"condition,omitempty"`
}
type FamilyMemberHistoryCondition struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `bson:"code" json:"code"`
	Outcome           *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
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

// UnmarshalFamilyMemberHistory unmarshals a FamilyMemberHistory.
func UnmarshalFamilyMemberHistory(b []byte) (FamilyMemberHistory, error) {
	var familyMemberHistory FamilyMemberHistory
	if err := json.Unmarshal(b, &familyMemberHistory); err != nil {
		return familyMemberHistory, err
	}
	return familyMemberHistory, nil
}
