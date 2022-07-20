package fhir

import "encoding/json"

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition        []Reference          `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn           []Reference          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces          []Reference          `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier   *Identifier          `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status            string               `bson:"status" json:"status"`
	Intent            string               `bson:"intent" json:"intent"`
	Priority          *string              `bson:"priority,omitempty" json:"priority,omitempty"`
	AuthoredOn        *string              `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Note              []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Action            []RequestGroupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type RequestGroupAction struct {
	Id                  *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label               *string                           `bson:"label,omitempty" json:"label,omitempty"`
	Title               *string                           `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                           `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                           `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Code                []CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	Type                *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *string                           `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *string                           `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *string                           `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *string                           `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *string                           `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Action              []RequestGroupAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type RequestGroupActionCondition struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              string      `bson:"kind" json:"kind"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
}
type RequestGroupActionRelatedAction struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string      `bson:"actionId" json:"actionId"`
	Relationship      string      `bson:"relationship" json:"relationship"`
}
type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalRequestGroup unmarshals a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}
