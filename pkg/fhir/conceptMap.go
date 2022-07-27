package fhir

import "encoding/json"

// ConceptMap is documented here http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMap struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string           `bson:"url" json:"url"`
	Identifier        *Identifier       `bson:"identifier" json:"identifier"`
	Version           *string           `bson:"version" json:"version"`
	Name              *string           `bson:"name" json:"name"`
	Title             *string           `bson:"title" json:"title"`
	Status            PublicationStatus `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool             `bson:"experimental" json:"experimental"`
	Date              *string           `bson:"date" json:"date"`
	Publisher         *string           `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail   `bson:"contact" json:"contact"`
	Description       *string           `bson:"description" json:"description"`
	UseContext        []UsageContext    `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string           `bson:"purpose" json:"purpose"`
	Copyright         *string           `bson:"copyright" json:"copyright"`
	SourceUri         *string           `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference   *Reference        `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	TargetUri         *string           `bson:"targetUri,omitempty" json:"targetUri,omitempty"`
	TargetReference   *Reference        `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
	Group             []ConceptMapGroup `bson:"group" json:"group"`
}
type ConceptMapGroup struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Source            *string                  `bson:"source" json:"source"`
	SourceVersion     *string                  `bson:"sourceVersion" json:"sourceVersion"`
	Target            *string                  `bson:"target" json:"target"`
	TargetVersion     *string                  `bson:"targetVersion" json:"targetVersion"`
	Element           []ConceptMapGroupElement `bson:"element,omitempty" json:"element,omitempty"`
	Unmapped          *ConceptMapGroupUnmapped `bson:"unmapped" json:"unmapped"`
}
type ConceptMapGroupElement struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Code              *string                        `bson:"code" json:"code"`
	Display           *string                        `bson:"display" json:"display"`
	Target            []ConceptMapGroupElementTarget `bson:"target" json:"target"`
}
type ConceptMapGroupElementTarget struct {
	Id                *string                                 `bson:"id" json:"id"`
	Extension         []Extension                             `bson:"extension" json:"extension"`
	ModifierExtension []Extension                             `bson:"modifierExtension" json:"modifierExtension"`
	Code              *string                                 `bson:"code" json:"code"`
	Display           *string                                 `bson:"display" json:"display"`
	Equivalence       *ConceptMapEquivalence                  `bson:"equivalence" json:"equivalence"`
	Comment           *string                                 `bson:"comment" json:"comment"`
	DependsOn         []ConceptMapGroupElementTargetDependsOn `bson:"dependsOn" json:"dependsOn"`
	Product           []ConceptMapGroupElementTargetDependsOn `bson:"product" json:"product"`
}
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Property          string      `bson:"property,omitempty" json:"property,omitempty"`
	System            *string     `bson:"system" json:"system"`
	Code              string      `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string     `bson:"display" json:"display"`
}
type ConceptMapGroupUnmapped struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Mode              ConceptMapGroupUnmappedMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Code              *string                     `bson:"code" json:"code"`
	Display           *string                     `bson:"display" json:"display"`
	Url               *string                     `bson:"url" json:"url"`
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
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType"`
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
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
