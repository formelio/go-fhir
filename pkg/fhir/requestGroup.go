package fhir

import "encoding/json"

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                    *string              `bson:"id" json:"id"`
	Meta                  *Meta                `bson:"meta" json:"meta"`
	ImplicitRules         *string              `bson:"implicitRules" json:"implicitRules"`
	Language              *string              `bson:"language" json:"language"`
	Text                  *Narrative           `bson:"text" json:"text"`
	Contained             []json.RawMessage    `bson:"contained" json:"contained"`
	Extension             []Extension          `bson:"extension" json:"extension"`
	ModifierExtension     []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            []Identifier         `bson:"identifier" json:"identifier"`
	Definition            []Reference          `bson:"definition" json:"definition"`
	BasedOn               []Reference          `bson:"basedOn" json:"basedOn"`
	Replaces              []Reference          `bson:"replaces" json:"replaces"`
	GroupIdentifier       *Identifier          `bson:"groupIdentifier" json:"groupIdentifier"`
	Status                RequestStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Intent                RequestIntent        `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority              *RequestPriority     `bson:"priority" json:"priority"`
	Subject               *Reference           `bson:"subject" json:"subject"`
	Context               *Reference           `bson:"context" json:"context"`
	AuthoredOn            *string              `bson:"authoredOn" json:"authoredOn"`
	Author                *Reference           `bson:"author" json:"author"`
	ReasonCodeableConcept *CodeableConcept     `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference           `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation         `bson:"note" json:"note"`
	Action                []RequestGroupAction `bson:"action" json:"action"`
}
type RequestGroupAction struct {
	Id                  *string                           `bson:"id" json:"id"`
	Extension           []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	Label               *string                           `bson:"label" json:"label"`
	Title               *string                           `bson:"title" json:"title"`
	Description         *string                           `bson:"description" json:"description"`
	TextEquivalent      *string                           `bson:"textEquivalent" json:"textEquivalent"`
	Code                []CodeableConcept                 `bson:"code" json:"code"`
	Documentation       []RelatedArtifact                 `bson:"documentation" json:"documentation"`
	Condition           []RequestGroupActionCondition     `bson:"condition" json:"condition"`
	RelatedAction       []RequestGroupActionRelatedAction `bson:"relatedAction" json:"relatedAction"`
	TimingDateTime      *string                           `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                         `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                            `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []Reference                       `bson:"participant" json:"participant"`
	Type                *Coding                           `bson:"type" json:"type"`
	GroupingBehavior    *ActionGroupingBehavior           `bson:"groupingBehavior" json:"groupingBehavior"`
	SelectionBehavior   *ActionSelectionBehavior          `bson:"selectionBehavior" json:"selectionBehavior"`
	RequiredBehavior    *ActionRequiredBehavior           `bson:"requiredBehavior" json:"requiredBehavior"`
	PrecheckBehavior    *ActionPrecheckBehavior           `bson:"precheckBehavior" json:"precheckBehavior"`
	CardinalityBehavior *ActionCardinalityBehavior        `bson:"cardinalityBehavior" json:"cardinalityBehavior"`
	Resource            *Reference                        `bson:"resource" json:"resource"`
	Action              []RequestGroupAction              `bson:"action" json:"action"`
}
type RequestGroupActionCondition struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Kind              ActionConditionKind `bson:"kind,omitempty" json:"kind,omitempty"`
	Description       *string             `bson:"description" json:"description"`
	Language          *string             `bson:"language" json:"language"`
	Expression        *string             `bson:"expression" json:"expression"`
}
type RequestGroupActionRelatedAction struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	ActionId          string                 `bson:"actionId,omitempty" json:"actionId,omitempty"`
	Relationship      ActionRelationshipType `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration    *Duration              `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
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

// UnmarshalRequestGroup unmarshalls a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}
