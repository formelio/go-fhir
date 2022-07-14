package fhir

import "encoding/json"

// Questionnaire is documented here http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string             `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string             `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string             `bson:"title,omitempty" json:"title,omitempty"`
	Status            string              `bson:"status" json:"status"`
	Experimental      *bool               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	ApprovalDate      *string             `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string             `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []UsageContext      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Contact           []ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code              []Coding            `bson:"code,omitempty" json:"code,omitempty"`
	SubjectType       []string            `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Item              []QuestionnaireItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItem struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                        `bson:"linkId" json:"linkId"`
	Definition        *string                       `bson:"definition,omitempty" json:"definition,omitempty"`
	Code              []Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Prefix            *string                       `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text              *string                       `bson:"text,omitempty" json:"text,omitempty"`
	Type              string                        `bson:"type" json:"type"`
	EnableWhen        []QuestionnaireItemEnableWhen `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	Required          *bool                         `bson:"required,omitempty" json:"required,omitempty"`
	Repeats           *bool                         `bson:"repeats,omitempty" json:"repeats,omitempty"`
	ReadOnly          *bool                         `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	MaxLength         *int                          `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Options           *Reference                    `bson:"options,omitempty" json:"options,omitempty"`
	Option            []QuestionnaireItemOption     `bson:"option,omitempty" json:"option,omitempty"`
	Item              []QuestionnaireItem           `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItemEnableWhen struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Question          string      `bson:"question" json:"question"`
	HasAnswer         *bool       `bson:"hasAnswer,omitempty" json:"hasAnswer,omitempty"`
}
type QuestionnaireItemOption struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherQuestionnaire Questionnaire

// MarshalJSON marshals the given Questionnaire as JSON into a byte slice
func (r Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaire
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaire: OtherQuestionnaire(r),
		ResourceType:       "Questionnaire",
	})
}

// UnmarshalQuestionnaire unmarshals a Questionnaire.
func UnmarshalQuestionnaire(b []byte) (Questionnaire, error) {
	var questionnaire Questionnaire
	if err := json.Unmarshal(b, &questionnaire); err != nil {
		return questionnaire, err
	}
	return questionnaire, nil
}
