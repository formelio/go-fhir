package fhir

import "encoding/json"

// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	RawContained      []json.RawMessage           `bson:"contained" json:"contained"`
	Contained         []IResource                 `bson:"-" json:"-"`
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

// OtherQuestionnaireResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherQuestionnaireResponse QuestionnaireResponse

// MarshalJSON marshals the given QuestionnaireResponse as JSON into a byte slice
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
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
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
}

// UnmarshalJSON unmarshals the given byte slice into QuestionnaireResponse
func (r *QuestionnaireResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherQuestionnaireResponse)(r)); err != nil {
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
func (r QuestionnaireResponse) GetResourceType() ResourceType {
	return ResourceTypeQuestionnaireResponse
}
