package fhir

import "encoding/json"

// AdverseEvent is documented here http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category          *string                     `bson:"category,omitempty" json:"category,omitempty"`
	Type              *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Date              *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Reaction          []Reference                 `bson:"reaction,omitempty" json:"reaction,omitempty"`
	Location          *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Seriousness       *CodeableConcept            `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	Outcome           *CodeableConcept            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	SuspectEntity     []AdverseEventSuspectEntity `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	ReferenceDocument []Reference                 `bson:"referenceDocument,omitempty" json:"referenceDocument,omitempty"`
	Study             []Reference                 `bson:"study,omitempty" json:"study,omitempty"`
}
type AdverseEventSuspectEntity struct {
	Id                          *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Causality                   *string          `bson:"causality,omitempty" json:"causality,omitempty"`
	CausalityAssessment         *CodeableConcept `bson:"causalityAssessment,omitempty" json:"causalityAssessment,omitempty"`
	CausalityProductRelatedness *string          `bson:"causalityProductRelatedness,omitempty" json:"causalityProductRelatedness,omitempty"`
	CausalityMethod             *CodeableConcept `bson:"causalityMethod,omitempty" json:"causalityMethod,omitempty"`
	CausalityResult             *CodeableConcept `bson:"causalityResult,omitempty" json:"causalityResult,omitempty"`
}
type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalAdverseEvent unmarshals a AdverseEvent.
func UnmarshalAdverseEvent(b []byte) (AdverseEvent, error) {
	var adverseEvent AdverseEvent
	if err := json.Unmarshal(b, &adverseEvent); err != nil {
		return adverseEvent, err
	}
	return adverseEvent, nil
}
