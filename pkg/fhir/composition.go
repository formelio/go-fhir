package fhir

import "encoding/json"

// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            CompositionStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Type              CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Class             *CodeableConcept       `bson:"class,omitempty" json:"class,omitempty"`
	Subject           Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              string                 `bson:"date,omitempty" json:"date,omitempty"`
	Author            []Reference            `bson:"author,omitempty" json:"author,omitempty"`
	Title             string                 `bson:"title,omitempty" json:"title,omitempty"`
	Confidentiality   *string                `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian         *Reference             `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `bson:"event,omitempty" json:"event,omitempty"`
	Section           []CompositionSection   `bson:"section,omitempty" json:"section,omitempty"`
}
type CompositionAttester struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              []CompositionAttestationMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Time              *string                      `bson:"time,omitempty" json:"time,omitempty"`
	Party             *Reference                   `bson:"party,omitempty" json:"party,omitempty"`
}
type CompositionRelatesTo struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `bson:"code,omitempty" json:"code,omitempty"`
	TargetIdentifier  *Identifier              `bson:"targetIdentifier,omitempty" json:"targetIdentifier,omitempty"`
	TargetReference   *Reference               `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
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
	Mode              *ListMode            `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy         *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry             []Reference          `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section           []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
}

// OtherComposition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherComposition Composition

// MarshalJSON marshals the given Composition as JSON into a byte slice
func (r Composition) MarshalJSON() ([]byte, error) {
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
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Composition
func (r *Composition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherComposition)(r)); err != nil {
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
func (r Composition) GetResourceType() ResourceType {
	return ResourceTypeComposition
}
