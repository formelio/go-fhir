package fhir

import "encoding/json"

// PlanDefinition is documented here http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                `bson:"title,omitempty" json:"title,omitempty"`
	Type              *CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Status            string                 `bson:"status" json:"status"`
	Experimental      *bool                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate      *string                `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic             []CodeableConcept      `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor       []Contributor          `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact           []ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string                `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact   []RelatedArtifact      `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library           []Reference            `bson:"library,omitempty" json:"library,omitempty"`
	Goal              []PlanDefinitionGoal   `bson:"goal,omitempty" json:"goal,omitempty"`
	Action            []PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionGoal struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Description       CodeableConcept            `bson:"description" json:"description"`
	Priority          *CodeableConcept           `bson:"priority,omitempty" json:"priority,omitempty"`
	Start             *CodeableConcept           `bson:"start,omitempty" json:"start,omitempty"`
	Addresses         []CodeableConcept          `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Documentation     []RelatedArtifact          `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Target            []PlanDefinitionGoalTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type PlanDefinitionGoalTarget struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Measure           *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	Due               *Duration        `bson:"due,omitempty" json:"due,omitempty"`
}
type PlanDefinitionAction struct {
	Id                  *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label               *string                             `bson:"label,omitempty" json:"label,omitempty"`
	Title               *string                             `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                             `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                             `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Code                []CodeableConcept                   `bson:"code,omitempty" json:"code,omitempty"`
	Reason              []CodeableConcept                   `bson:"reason,omitempty" json:"reason,omitempty"`
	Documentation       []RelatedArtifact                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	GoalId              []string                            `bson:"goalId,omitempty" json:"goalId,omitempty"`
	TriggerDefinition   []TriggerDefinition                 `bson:"triggerDefinition,omitempty" json:"triggerDefinition,omitempty"`
	Condition           []PlanDefinitionActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	Input               []DataRequirement                   `bson:"input,omitempty" json:"input,omitempty"`
	Output              []DataRequirement                   `bson:"output,omitempty" json:"output,omitempty"`
	RelatedAction       []PlanDefinitionActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	Participant         []PlanDefinitionActionParticipant   `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *Coding                             `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *string                             `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *string                             `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *string                             `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *string                             `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *string                             `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Transform           *Reference                          `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue        []PlanDefinitionActionDynamicValue  `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	Action              []PlanDefinitionAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type PlanDefinitionActionCondition struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              string      `bson:"kind" json:"kind"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
}
type PlanDefinitionActionRelatedAction struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string      `bson:"actionId" json:"actionId"`
	Relationship      string      `bson:"relationship" json:"relationship"`
}
type PlanDefinitionActionParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string           `bson:"type" json:"type"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type PlanDefinitionActionDynamicValue struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
}
type OtherPlanDefinition PlanDefinition

// MarshalJSON marshals the given PlanDefinition as JSON into a byte slice
func (r PlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPlanDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPlanDefinition: OtherPlanDefinition(r),
		ResourceType:        "PlanDefinition",
	})
}

// UnmarshalPlanDefinition unmarshals a PlanDefinition.
func UnmarshalPlanDefinition(b []byte) (PlanDefinition, error) {
	var planDefinition PlanDefinition
	if err := json.Unmarshal(b, &planDefinition); err != nil {
		return planDefinition, err
	}
	return planDefinition, nil
}
