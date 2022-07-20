package fhir

import "encoding/json"

// Questionnaire is documented here http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                *string             `bson:"id" json:"id"`
	Meta              *Meta               `bson:"meta" json:"meta"`
	ImplicitRules     *string             `bson:"implicitRules" json:"implicitRules"`
	Language          *string             `bson:"language" json:"language"`
	Text              *Narrative          `bson:"text" json:"text"`
	Contained         []json.RawMessage   `bson:"contained" json:"contained"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string             `bson:"url" json:"url"`
	Identifier        []Identifier        `bson:"identifier" json:"identifier"`
	Version           *string             `bson:"version" json:"version"`
	Name              *string             `bson:"name" json:"name"`
	Title             *string             `bson:"title" json:"title"`
	Status            PublicationStatus   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool               `bson:"experimental" json:"experimental"`
	Date              *string             `bson:"date" json:"date"`
	Publisher         *string             `bson:"publisher" json:"publisher"`
	Description       *string             `bson:"description" json:"description"`
	Purpose           *string             `bson:"purpose" json:"purpose"`
	ApprovalDate      *string             `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate    *string             `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod   *Period             `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext        []UsageContext      `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept   `bson:"jurisdiction" json:"jurisdiction"`
	Contact           []ContactDetail     `bson:"contact" json:"contact"`
	Copyright         *string             `bson:"copyright" json:"copyright"`
	Code              []Coding            `bson:"code" json:"code"`
	SubjectType       []ResourceType      `bson:"subjectType" json:"subjectType"`
	Item              []QuestionnaireItem `bson:"item" json:"item"`
}
type QuestionnaireItem struct {
	Id                *string                       `bson:"id" json:"id"`
	Extension         []Extension                   `bson:"extension" json:"extension"`
	ModifierExtension []Extension                   `bson:"modifierExtension" json:"modifierExtension"`
	LinkId            string                        `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition        *string                       `bson:"definition" json:"definition"`
	Code              []Coding                      `bson:"code" json:"code"`
	Prefix            *string                       `bson:"prefix" json:"prefix"`
	Text              *string                       `bson:"text" json:"text"`
	Type              QuestionnaireItemType         `bson:"type,omitempty" json:"type,omitempty"`
	EnableWhen        []QuestionnaireItemEnableWhen `bson:"enableWhen" json:"enableWhen"`
	Required          *bool                         `bson:"required" json:"required"`
	Repeats           *bool                         `bson:"repeats" json:"repeats"`
	ReadOnly          *bool                         `bson:"readOnly" json:"readOnly"`
	MaxLength         *int                          `bson:"maxLength" json:"maxLength"`
	Options           *Reference                    `bson:"options" json:"options"`
	Option            []QuestionnaireItemOption     `bson:"option" json:"option"`
	InitialBoolean    *bool                         `bson:"initialBoolean,omitempty" json:"initialBoolean,omitempty"`
	InitialDecimal    *float64                      `bson:"initialDecimal,omitempty" json:"initialDecimal,omitempty"`
	InitialInteger    *int                          `bson:"initialInteger,omitempty" json:"initialInteger,omitempty"`
	InitialDate       *string                       `bson:"initialDate,omitempty" json:"initialDate,omitempty"`
	InitialDateTime   *string                       `bson:"initialDateTime,omitempty" json:"initialDateTime,omitempty"`
	InitialTime       *string                       `bson:"initialTime,omitempty" json:"initialTime,omitempty"`
	InitialString     *string                       `bson:"initialString,omitempty" json:"initialString,omitempty"`
	InitialUri        *string                       `bson:"initialUri,omitempty" json:"initialUri,omitempty"`
	InitialAttachment *Attachment                   `bson:"initialAttachment,omitempty" json:"initialAttachment,omitempty"`
	InitialCoding     *Coding                       `bson:"initialCoding,omitempty" json:"initialCoding,omitempty"`
	InitialQuantity   *Quantity                     `bson:"initialQuantity,omitempty" json:"initialQuantity,omitempty"`
	InitialReference  *Reference                    `bson:"initialReference,omitempty" json:"initialReference,omitempty"`
	Item              []QuestionnaireItem           `bson:"item" json:"item"`
}
type QuestionnaireItemEnableWhen struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Question          string      `bson:"question,omitempty" json:"question,omitempty"`
	HasAnswer         *bool       `bson:"hasAnswer" json:"hasAnswer"`
	AnswerBoolean     *bool       `bson:"answerBoolean,omitempty" json:"answerBoolean,omitempty"`
	AnswerDecimal     *float64    `bson:"answerDecimal,omitempty" json:"answerDecimal,omitempty"`
	AnswerInteger     *int        `bson:"answerInteger,omitempty" json:"answerInteger,omitempty"`
	AnswerDate        *string     `bson:"answerDate,omitempty" json:"answerDate,omitempty"`
	AnswerDateTime    *string     `bson:"answerDateTime,omitempty" json:"answerDateTime,omitempty"`
	AnswerTime        *string     `bson:"answerTime,omitempty" json:"answerTime,omitempty"`
	AnswerString      *string     `bson:"answerString,omitempty" json:"answerString,omitempty"`
	AnswerUri         *string     `bson:"answerUri,omitempty" json:"answerUri,omitempty"`
	AnswerAttachment  *Attachment `bson:"answerAttachment,omitempty" json:"answerAttachment,omitempty"`
	AnswerCoding      *Coding     `bson:"answerCoding,omitempty" json:"answerCoding,omitempty"`
	AnswerQuantity    *Quantity   `bson:"answerQuantity,omitempty" json:"answerQuantity,omitempty"`
	AnswerReference   *Reference  `bson:"answerReference,omitempty" json:"answerReference,omitempty"`
}
type QuestionnaireItemOption struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate         *string     `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueTime         *string     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueCoding       *Coding     `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
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

// UnmarshalQuestionnaire unmarshalls a Questionnaire.
func UnmarshalQuestionnaire(b []byte) (Questionnaire, error) {
	var questionnaire Questionnaire
	if err := json.Unmarshal(b, &questionnaire); err != nil {
		return questionnaire, err
	}
	return questionnaire, nil
}
