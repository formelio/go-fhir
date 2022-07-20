package fhir

import "encoding/json"

// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	Contained         []json.RawMessage           `bson:"contained" json:"contained"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier                 `bson:"identifier" json:"identifier"`
	BasedOn           []Reference                 `bson:"basedOn" json:"basedOn"`
	Parent            []Reference                 `bson:"parent" json:"parent"`
	Questionnaire     *Reference                  `bson:"questionnaire" json:"questionnaire"`
	Status            QuestionnaireResponseStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject           *Reference                  `bson:"subject" json:"subject"`
	Context           *Reference                  `bson:"context" json:"context"`
	Authored          *string                     `bson:"authored" json:"authored"`
	Author            *Reference                  `bson:"author" json:"author"`
	Source            *Reference                  `bson:"source" json:"source"`
	Item              []QuestionnaireResponseItem `bson:"item" json:"item"`
}
type QuestionnaireResponseItem struct {
	Id                *string                           `bson:"id" json:"id"`
	Extension         []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	LinkId            string                            `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition        *string                           `bson:"definition" json:"definition"`
	Text              *string                           `bson:"text" json:"text"`
	Subject           *Reference                        `bson:"subject" json:"subject"`
	Answer            []QuestionnaireResponseItemAnswer `bson:"answer" json:"answer"`
	Item              []QuestionnaireResponseItem       `bson:"item" json:"item"`
}
type QuestionnaireResponseItemAnswer struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	ValueBoolean      *bool                       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal      *float64                    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger      *int                        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate         *string                     `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime     *string                     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueTime         *string                     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString       *string                     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri          *string                     `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment   *Attachment                 `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding       *Coding                     `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity     *Quantity                   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference    *Reference                  `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item" json:"item"`
}
type OtherQuestionnaireResponse QuestionnaireResponse

// MarshalJSON marshals the given QuestionnaireResponse as JSON into a byte slice
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
}

// UnmarshalQuestionnaireResponse unmarshalls a QuestionnaireResponse.
func UnmarshalQuestionnaireResponse(b []byte) (QuestionnaireResponse, error) {
	var questionnaireResponse QuestionnaireResponse
	if err := json.Unmarshal(b, &questionnaireResponse); err != nil {
		return questionnaireResponse, err
	}
	return questionnaireResponse, nil
}
