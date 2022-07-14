package fhir

import "encoding/json"

// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                 `bson:"status" json:"status"`
	Type              CodeableConcept        `bson:"type" json:"type"`
	Class             *CodeableConcept       `bson:"class,omitempty" json:"class,omitempty"`
	Subject           Reference              `bson:"subject" json:"subject"`
	Encounter         *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              string                 `bson:"date" json:"date"`
	Title             string                 `bson:"title" json:"title"`
	Confidentiality   *string                `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian         *Reference             `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `bson:"event,omitempty" json:"event,omitempty"`
	Section           []CompositionSection   `bson:"section,omitempty" json:"section,omitempty"`
}
type CompositionAttester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              []string    `bson:"mode" json:"mode"`
	Time              *string     `bson:"time,omitempty" json:"time,omitempty"`
}
type CompositionRelatesTo struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
}
type CompositionEvent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type CompositionSection struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Mode              *string              `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy         *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry             []Reference          `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section           []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
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

// UnmarshalComposition unmarshals a Composition.
func UnmarshalComposition(b []byte) (Composition, error) {
	var composition Composition
	if err := json.Unmarshal(b, &composition); err != nil {
		return composition, err
	}
	return composition, nil
}
