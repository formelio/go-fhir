package fhir

import "encoding/json"

// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier            `bson:"identifier" json:"identifier"`
	Status            CompositionStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Type              CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Class             *CodeableConcept       `bson:"class" json:"class"`
	Subject           Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference             `bson:"encounter" json:"encounter"`
	Date              string                 `bson:"date,omitempty" json:"date,omitempty"`
	Author            []Reference            `bson:"author,omitempty" json:"author,omitempty"`
	Title             string                 `bson:"title,omitempty" json:"title,omitempty"`
	Confidentiality   *string                `bson:"confidentiality" json:"confidentiality"`
	Attester          []CompositionAttester  `bson:"attester" json:"attester"`
	Custodian         *Reference             `bson:"custodian" json:"custodian"`
	RelatesTo         []CompositionRelatesTo `bson:"relatesTo" json:"relatesTo"`
	Event             []CompositionEvent     `bson:"event" json:"event"`
	Section           []CompositionSection   `bson:"section" json:"section"`
}
type CompositionAttester struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Mode              []CompositionAttestationMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Time              *string                      `bson:"time" json:"time"`
	Party             *Reference                   `bson:"party" json:"party"`
}
type CompositionRelatesTo struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Code              DocumentRelationshipType `bson:"code,omitempty" json:"code,omitempty"`
	TargetIdentifier  *Identifier              `bson:"targetIdentifier,omitempty" json:"targetIdentifier,omitempty"`
	TargetReference   *Reference               `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
}
type CompositionEvent struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Code              []CodeableConcept `bson:"code" json:"code"`
	Period            *Period           `bson:"period" json:"period"`
	Detail            []Reference       `bson:"detail" json:"detail"`
}
type CompositionSection struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Title             *string              `bson:"title" json:"title"`
	Code              *CodeableConcept     `bson:"code" json:"code"`
	Text              *Narrative           `bson:"text" json:"text"`
	Mode              *ListMode            `bson:"mode" json:"mode"`
	OrderedBy         *CodeableConcept     `bson:"orderedBy" json:"orderedBy"`
	Entry             []Reference          `bson:"entry" json:"entry"`
	EmptyReason       *CodeableConcept     `bson:"emptyReason" json:"emptyReason"`
	Section           []CompositionSection `bson:"section" json:"section"`
}
type OtherComposition Composition

// MarshalJSON marshals the given Composition as JSON into a byte slice
func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}

// UnmarshalComposition unmarshalls a Composition.
func UnmarshalComposition(b []byte) (Composition, error) {
	var composition Composition
	if err := json.Unmarshal(b, &composition); err != nil {
		return composition, err
	}
	return composition, nil
}
