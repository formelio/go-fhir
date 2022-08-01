package fhir

import (
	"bytes"
	"encoding/json"
)

// Questionnaire is documented here http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource          `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string              `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []*Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string              `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string              `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string              `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string              `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string              `bson:"purpose,omitempty" json:"purpose,omitempty"`
	ApprovalDate      *string              `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string              `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []*UsageContext      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Contact           []*ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code              []*Coding            `bson:"code,omitempty" json:"code,omitempty"`
	SubjectType       []*ResourceType      `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Item              []*QuestionnaireItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItem struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                         `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition        *string                        `bson:"definition,omitempty" json:"definition,omitempty"`
	Code              []*Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Prefix            *string                        `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text              *string                        `bson:"text,omitempty" json:"text,omitempty"`
	Type              QuestionnaireItemType          `bson:"type,omitempty" json:"type,omitempty"`
	EnableWhen        []*QuestionnaireItemEnableWhen `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	Required          *bool                          `bson:"required,omitempty" json:"required,omitempty"`
	Repeats           *bool                          `bson:"repeats,omitempty" json:"repeats,omitempty"`
	ReadOnly          *bool                          `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	MaxLength         *int                           `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Options           *Reference                     `bson:"options,omitempty" json:"options,omitempty"`
	Option            []*QuestionnaireItemOption     `bson:"option,omitempty" json:"option,omitempty"`
	InitialBoolean    *bool                          `bson:"initialBoolean,omitempty" json:"initialBoolean,omitempty"`
	InitialDecimal    *float64                       `bson:"initialDecimal,omitempty" json:"initialDecimal,omitempty"`
	InitialInteger    *int                           `bson:"initialInteger,omitempty" json:"initialInteger,omitempty"`
	InitialDate       *string                        `bson:"initialDate,omitempty" json:"initialDate,omitempty"`
	InitialDateTime   *string                        `bson:"initialDateTime,omitempty" json:"initialDateTime,omitempty"`
	InitialTime       *string                        `bson:"initialTime,omitempty" json:"initialTime,omitempty"`
	InitialString     *string                        `bson:"initialString,omitempty" json:"initialString,omitempty"`
	InitialUri        *string                        `bson:"initialUri,omitempty" json:"initialUri,omitempty"`
	InitialAttachment *Attachment                    `bson:"initialAttachment,omitempty" json:"initialAttachment,omitempty"`
	InitialCoding     *Coding                        `bson:"initialCoding,omitempty" json:"initialCoding,omitempty"`
	InitialQuantity   *Quantity                      `bson:"initialQuantity,omitempty" json:"initialQuantity,omitempty"`
	InitialReference  *Reference                     `bson:"initialReference,omitempty" json:"initialReference,omitempty"`
	Item              []*QuestionnaireItem           `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItemEnableWhen struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Question          string       `bson:"question,omitempty" json:"question,omitempty"`
	HasAnswer         *bool        `bson:"hasAnswer,omitempty" json:"hasAnswer,omitempty"`
	AnswerBoolean     *bool        `bson:"answerBoolean,omitempty" json:"answerBoolean,omitempty"`
	AnswerDecimal     *float64     `bson:"answerDecimal,omitempty" json:"answerDecimal,omitempty"`
	AnswerInteger     *int         `bson:"answerInteger,omitempty" json:"answerInteger,omitempty"`
	AnswerDate        *string      `bson:"answerDate,omitempty" json:"answerDate,omitempty"`
	AnswerDateTime    *string      `bson:"answerDateTime,omitempty" json:"answerDateTime,omitempty"`
	AnswerTime        *string      `bson:"answerTime,omitempty" json:"answerTime,omitempty"`
	AnswerString      *string      `bson:"answerString,omitempty" json:"answerString,omitempty"`
	AnswerUri         *string      `bson:"answerUri,omitempty" json:"answerUri,omitempty"`
	AnswerAttachment  *Attachment  `bson:"answerAttachment,omitempty" json:"answerAttachment,omitempty"`
	AnswerCoding      *Coding      `bson:"answerCoding,omitempty" json:"answerCoding,omitempty"`
	AnswerQuantity    *Quantity    `bson:"answerQuantity,omitempty" json:"answerQuantity,omitempty"`
	AnswerReference   *Reference   `bson:"answerReference,omitempty" json:"answerReference,omitempty"`
}
type QuestionnaireItemOption struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueInteger      *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate         *string      `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueTime         *string      `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString       *string      `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueCoding       *Coding      `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
}

// OtherQuestionnaire is a helper type to use the default implementations of Marshall and Unmarshal
type OtherQuestionnaire Questionnaire

// MarshalJSON marshals the given Questionnaire as JSON into a byte slice
func (r Questionnaire) MarshalJSON() ([]byte, error) {
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
		OtherQuestionnaire
	}{
		OtherQuestionnaire: OtherQuestionnaire(r),
		ResourceType:       "Questionnaire",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Questionnaire
func (r *Questionnaire) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherQuestionnaire)(r)); err != nil {
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
func (r Questionnaire) GetResourceType() ResourceType {
	return ResourceTypeQuestionnaire
}
