package fhir

import "encoding/json"

// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Questionnaire     *Reference                  `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status            string                      `bson:"status" json:"status"`
	Subject           *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Authored          *string                     `bson:"authored,omitempty" json:"authored,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireResponseItem struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                            `bson:"linkId" json:"linkId"`
	Definition        *string                           `bson:"definition,omitempty" json:"definition,omitempty"`
	Text              *string                           `bson:"text,omitempty" json:"text,omitempty"`
	Subject           *Reference                        `bson:"subject,omitempty" json:"subject,omitempty"`
	Answer            []QuestionnaireResponseItemAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	Item              []QuestionnaireResponseItem       `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireResponseItemAnswer struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
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

// UnmarshalQuestionnaireResponse unmarshals a QuestionnaireResponse.
func UnmarshalQuestionnaireResponse(b []byte) (QuestionnaireResponse, error) {
	var questionnaireResponse QuestionnaireResponse
	if err := json.Unmarshal(b, &questionnaireResponse); err != nil {
		return questionnaireResponse, err
	}
	return questionnaireResponse, nil
}
