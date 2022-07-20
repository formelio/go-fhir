package fhir

import "encoding/json"

// ConceptMap is documented here http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMap struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
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
type OtherConceptMap ConceptMap

// MarshalJSON marshals the given ConceptMap as JSON into a byte slice
func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType"`
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
}

// UnmarshalConceptMap unmarshalls a ConceptMap.
func UnmarshalConceptMap(b []byte) (ConceptMap, error) {
	var conceptMap ConceptMap
	if err := json.Unmarshal(b, &conceptMap); err != nil {
		return conceptMap, err
	}
	return conceptMap, nil
}
