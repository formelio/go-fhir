package fhir

import (
	"bytes"
	"encoding/json"
)

// PlanDefinition is documented here http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                 `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []*Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Type              *CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Status            PublicationStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage             *string                 `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate      *string                 `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                 `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                 `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []*UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic             []*CodeableConcept      `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor       []*Contributor          `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact           []*ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact   []*RelatedArtifact      `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library           []*Reference            `bson:"library,omitempty" json:"library,omitempty"`
	Goal              []*PlanDefinitionGoal   `bson:"goal,omitempty" json:"goal,omitempty"`
	Action            []*PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionGoal struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept            `bson:"category,omitempty" json:"category,omitempty"`
	Description       CodeableConcept             `bson:"description,omitempty" json:"description,omitempty"`
	Priority          *CodeableConcept            `bson:"priority,omitempty" json:"priority,omitempty"`
	Start             *CodeableConcept            `bson:"start,omitempty" json:"start,omitempty"`
	Addresses         []*CodeableConcept          `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Documentation     []*RelatedArtifact          `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Target            []*PlanDefinitionGoalTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type PlanDefinitionGoalTarget struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty"`
	Due                   *Duration        `bson:"due,omitempty" json:"due,omitempty"`
}
type PlanDefinitionAction struct {
	Id                  *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []*Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label               *string                              `bson:"label,omitempty" json:"label,omitempty"`
	Title               *string                              `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                              `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                              `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Code                []*CodeableConcept                   `bson:"code,omitempty" json:"code,omitempty"`
	Reason              []*CodeableConcept                   `bson:"reason,omitempty" json:"reason,omitempty"`
	Documentation       []*RelatedArtifact                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	GoalId              []*string                            `bson:"goalId,omitempty" json:"goalId,omitempty"`
	TriggerDefinition   []*TriggerDefinition                 `bson:"triggerDefinition,omitempty" json:"triggerDefinition,omitempty"`
	Condition           []*PlanDefinitionActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	Input               []*DataRequirement                   `bson:"input,omitempty" json:"input,omitempty"`
	Output              []*DataRequirement                   `bson:"output,omitempty" json:"output,omitempty"`
	RelatedAction       []*PlanDefinitionActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	TimingDateTime      *string                              `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                              `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                            `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                               `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                              `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []*PlanDefinitionActionParticipant   `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *Coding                              `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior              `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior             `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior              `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior              `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior           `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Definition          *Reference                           `bson:"definition,omitempty" json:"definition,omitempty"`
	Transform           *Reference                           `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue        []*PlanDefinitionActionDynamicValue  `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	Action              []*PlanDefinitionAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionActionCondition struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `bson:"kind,omitempty" json:"kind,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string             `bson:"expression,omitempty" json:"expression,omitempty"`
}
type PlanDefinitionActionRelatedAction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string                 `bson:"actionId,omitempty" json:"actionId,omitempty"`
	Relationship      ActionRelationshipType `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration    *Duration              `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}
type PlanDefinitionActionParticipant struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ActionParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Role              *CodeableConcept      `bson:"role,omitempty" json:"role,omitempty"`
}
type PlanDefinitionActionDynamicValue struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Path              *string      `bson:"path,omitempty" json:"path,omitempty"`
	Language          *string      `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string      `bson:"expression,omitempty" json:"expression,omitempty"`
}

// OtherPlanDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPlanDefinition PlanDefinition

// MarshalJSON marshals the given PlanDefinition as JSON into a byte slice
func (r PlanDefinition) MarshalJSON() ([]byte, error) {
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
		OtherPlanDefinition
	}{
		OtherPlanDefinition: OtherPlanDefinition(r),
		ResourceType:        "PlanDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into PlanDefinition
func (r *PlanDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPlanDefinition)(r)); err != nil {
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
func (r PlanDefinition) GetResourceType() ResourceType {
	return ResourceTypePlanDefinition
}
