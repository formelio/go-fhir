package fhir

import "encoding/json"

// PlanDefinition is documented here http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string                `bson:"url" json:"url"`
	Identifier        []Identifier           `bson:"identifier" json:"identifier"`
	Version           *string                `bson:"version" json:"version"`
	Name              *string                `bson:"name" json:"name"`
	Title             *string                `bson:"title" json:"title"`
	Type              *CodeableConcept       `bson:"type" json:"type"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                  `bson:"experimental" json:"experimental"`
	Date              *string                `bson:"date" json:"date"`
	Publisher         *string                `bson:"publisher" json:"publisher"`
	Description       *string                `bson:"description" json:"description"`
	Purpose           *string                `bson:"purpose" json:"purpose"`
	Usage             *string                `bson:"usage" json:"usage"`
	ApprovalDate      *string                `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate    *string                `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod   *Period                `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext        []UsageContext         `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction" json:"jurisdiction"`
	Topic             []CodeableConcept      `bson:"topic" json:"topic"`
	Contributor       []Contributor          `bson:"contributor" json:"contributor"`
	Contact           []ContactDetail        `bson:"contact" json:"contact"`
	Copyright         *string                `bson:"copyright" json:"copyright"`
	RelatedArtifact   []RelatedArtifact      `bson:"relatedArtifact" json:"relatedArtifact"`
	Library           []Reference            `bson:"library" json:"library"`
	Goal              []PlanDefinitionGoal   `bson:"goal" json:"goal"`
	Action            []PlanDefinitionAction `bson:"action" json:"action"`
}
type PlanDefinitionGoal struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Category          *CodeableConcept           `bson:"category" json:"category"`
	Description       CodeableConcept            `bson:"description,omitempty" json:"description,omitempty"`
	Priority          *CodeableConcept           `bson:"priority" json:"priority"`
	Start             *CodeableConcept           `bson:"start" json:"start"`
	Addresses         []CodeableConcept          `bson:"addresses" json:"addresses"`
	Documentation     []RelatedArtifact          `bson:"documentation" json:"documentation"`
	Target            []PlanDefinitionGoalTarget `bson:"target" json:"target"`
}
type PlanDefinitionGoalTarget struct {
	Id                    *string          `bson:"id" json:"id"`
	Extension             []Extension      `bson:"extension" json:"extension"`
	ModifierExtension     []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Measure               *CodeableConcept `bson:"measure" json:"measure"`
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty"`
	Due                   *Duration        `bson:"due" json:"due"`
}
type PlanDefinitionAction struct {
	Id                  *string                             `bson:"id" json:"id"`
	Extension           []Extension                         `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                         `bson:"modifierExtension" json:"modifierExtension"`
	Label               *string                             `bson:"label" json:"label"`
	Title               *string                             `bson:"title" json:"title"`
	Description         *string                             `bson:"description" json:"description"`
	TextEquivalent      *string                             `bson:"textEquivalent" json:"textEquivalent"`
	Code                []CodeableConcept                   `bson:"code" json:"code"`
	Reason              []CodeableConcept                   `bson:"reason" json:"reason"`
	Documentation       []RelatedArtifact                   `bson:"documentation" json:"documentation"`
	GoalId              []string                            `bson:"goalId" json:"goalId"`
	TriggerDefinition   []TriggerDefinition                 `bson:"triggerDefinition" json:"triggerDefinition"`
	Condition           []PlanDefinitionActionCondition     `bson:"condition" json:"condition"`
	Input               []DataRequirement                   `bson:"input" json:"input"`
	Output              []DataRequirement                   `bson:"output" json:"output"`
	RelatedAction       []PlanDefinitionActionRelatedAction `bson:"relatedAction" json:"relatedAction"`
	TimingDateTime      *string                             `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                             `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                           `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                              `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                             `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []PlanDefinitionActionParticipant   `bson:"participant" json:"participant"`
	Type                *Coding                             `bson:"type" json:"type"`
	GroupingBehavior    *ActionGroupingBehavior             `bson:"groupingBehavior" json:"groupingBehavior"`
	SelectionBehavior   *ActionSelectionBehavior            `bson:"selectionBehavior" json:"selectionBehavior"`
	RequiredBehavior    *ActionRequiredBehavior             `bson:"requiredBehavior" json:"requiredBehavior"`
	PrecheckBehavior    *ActionPrecheckBehavior             `bson:"precheckBehavior" json:"precheckBehavior"`
	CardinalityBehavior *ActionCardinalityBehavior          `bson:"cardinalityBehavior" json:"cardinalityBehavior"`
	Definition          *Reference                          `bson:"definition" json:"definition"`
	Transform           *Reference                          `bson:"transform" json:"transform"`
	DynamicValue        []PlanDefinitionActionDynamicValue  `bson:"dynamicValue" json:"dynamicValue"`
	Action              []PlanDefinitionAction              `bson:"action" json:"action"`
}
type PlanDefinitionActionCondition struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Kind              ActionConditionKind `bson:"kind,omitempty" json:"kind,omitempty"`
	Description       *string             `bson:"description" json:"description"`
	Language          *string             `bson:"language" json:"language"`
	Expression        *string             `bson:"expression" json:"expression"`
}
type PlanDefinitionActionRelatedAction struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	ActionId          string                 `bson:"actionId,omitempty" json:"actionId,omitempty"`
	Relationship      ActionRelationshipType `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration    *Duration              `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}
type PlanDefinitionActionParticipant struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Type              ActionParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Role              *CodeableConcept      `bson:"role" json:"role"`
}
type PlanDefinitionActionDynamicValue struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       *string     `bson:"description" json:"description"`
	Path              *string     `bson:"path" json:"path"`
	Language          *string     `bson:"language" json:"language"`
	Expression        *string     `bson:"expression" json:"expression"`
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
	return json.Marshal(struct {
		OtherPlanDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPlanDefinition: OtherPlanDefinition(r),
		ResourceType:        "PlanDefinition",
	})
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
