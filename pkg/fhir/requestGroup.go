package fhir

import "encoding/json"

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                    *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource          `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition            []Reference          `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn               []Reference          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference          `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                RequestStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Intent                RequestIntent        `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority              *RequestPriority     `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject               *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Context               *Reference           `bson:"context,omitempty" json:"context,omitempty"`
	AuthoredOn            *string              `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Author                *Reference           `bson:"author,omitempty" json:"author,omitempty"`
	ReasonCodeableConcept *CodeableConcept     `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference           `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Action                []RequestGroupAction `bson:"action,omitempty" json:"action,omitempty"`
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
	TimingDateTime      *string                           `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                         `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                            `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []Reference                       `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior           `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior          `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior           `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior           `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior        `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Action              []RequestGroupAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type RequestGroupActionCondition struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `bson:"kind,omitempty" json:"kind,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string             `bson:"expression,omitempty" json:"expression,omitempty"`
}
type RequestGroupActionRelatedAction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string                 `bson:"actionId,omitempty" json:"actionId,omitempty"`
	Relationship      ActionRelationshipType `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration    *Duration              `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}

// OtherRequestGroup is a helper type to use the default implementations of Marshall and Unmarshal
type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
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
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalJSON unmarshals the given byte slice into RequestGroup
func (r *RequestGroup) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherRequestGroup)(r)); err != nil {
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
func (r RequestGroup) GetResourceType() ResourceType {
	return ResourceTypeRequestGroup
}
