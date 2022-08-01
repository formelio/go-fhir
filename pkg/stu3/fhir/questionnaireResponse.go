package fhir

import "encoding/json"

// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent            []Reference                 `bson:"parent,omitempty" json:"parent,omitempty"`
	Questionnaire     *Reference                  `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status            QuestionnaireResponseStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject           *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context           *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	Authored          *string                     `bson:"authored,omitempty" json:"authored,omitempty"`
	Author            *Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Source            *Reference                  `bson:"source,omitempty" json:"source,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireResponseItem struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                            `bson:"linkId,omitempty" json:"linkId,omitempty"`
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
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
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
