package fhir

import (
	"bytes"
	"encoding/json"
)

// ConceptMap is documented here http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMap struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string            `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string            `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail   `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	SourceUri         *string            `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference   *Reference         `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	TargetUri         *string            `bson:"targetUri,omitempty" json:"targetUri,omitempty"`
	TargetReference   *Reference         `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
	Group             []*ConceptMapGroup `bson:"group,omitempty" json:"group,omitempty"`
}
type ConceptMapGroup struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Source            *string                  `bson:"source,omitempty" json:"source,omitempty"`
	SourceVersion     *string                  `bson:"sourceVersion,omitempty" json:"sourceVersion,omitempty"`
	Target            *string                  `bson:"target,omitempty" json:"target,omitempty"`
	TargetVersion     *string                  `bson:"targetVersion,omitempty" json:"targetVersion,omitempty"`
	Element           []ConceptMapGroupElement `bson:"element,omitempty" json:"element,omitempty"`
	Unmapped          *ConceptMapGroupUnmapped `bson:"unmapped,omitempty" json:"unmapped,omitempty"`
}
type ConceptMapGroupElement struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *string                         `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                         `bson:"display,omitempty" json:"display,omitempty"`
	Target            []*ConceptMapGroupElementTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type ConceptMapGroupElementTarget struct {
	Id                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *string                                  `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                  `bson:"display,omitempty" json:"display,omitempty"`
	Equivalence       *ConceptMapEquivalence                   `bson:"equivalence,omitempty" json:"equivalence,omitempty"`
	Comment           *string                                  `bson:"comment,omitempty" json:"comment,omitempty"`
	DependsOn         []*ConceptMapGroupElementTargetDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Product           []*ConceptMapGroupElementTargetDependsOn `bson:"product,omitempty" json:"product,omitempty"`
}
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          string       `bson:"property,omitempty" json:"property,omitempty"`
	System            *string      `bson:"system,omitempty" json:"system,omitempty"`
	Code              string       `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string      `bson:"display,omitempty" json:"display,omitempty"`
}
type ConceptMapGroupUnmapped struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              ConceptMapGroupUnmappedMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Code              *string                     `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                     `bson:"display,omitempty" json:"display,omitempty"`
	Url               *string                     `bson:"url,omitempty" json:"url,omitempty"`
}

// OtherConceptMap is a helper type to use the default implementations of Marshall and Unmarshal
type OtherConceptMap ConceptMap

// MarshalJSON marshals the given ConceptMap as JSON into a byte slice
func (r ConceptMap) MarshalJSON() ([]byte, error) {
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
		OtherConceptMap
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into ConceptMap
func (r *ConceptMap) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherConceptMap)(r)); err != nil {
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
func (r ConceptMap) GetResourceType() ResourceType {
	return ResourceTypeConceptMap
}
