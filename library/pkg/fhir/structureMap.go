package fhir

import "encoding/json"

// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMap struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                  `bson:"url" json:"url"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                  `bson:"name" json:"name"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status            string                  `bson:"status" json:"status"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Structure         []StructureMapStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	Import            []string                `bson:"import,omitempty" json:"import,omitempty"`
	Group             []StructureMapGroup     `bson:"group" json:"group"`
}
type StructureMapStructure struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string      `bson:"url" json:"url"`
	Mode              string      `bson:"mode" json:"mode"`
	Alias             *string     `bson:"alias,omitempty" json:"alias,omitempty"`
	Documentation     *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroup struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                   `bson:"name" json:"name"`
	Extends           *string                  `bson:"extends,omitempty" json:"extends,omitempty"`
	TypeMode          string                   `bson:"typeMode" json:"typeMode"`
	Documentation     *string                  `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Input             []StructureMapGroupInput `bson:"input" json:"input"`
	Rule              []StructureMapGroupRule  `bson:"rule" json:"rule"`
}
type StructureMapGroupInput struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Type              *string     `bson:"type,omitempty" json:"type,omitempty"`
	Mode              string      `bson:"mode" json:"mode"`
	Documentation     *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRule struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                           `bson:"name" json:"name"`
	Source            []StructureMapGroupRuleSource    `bson:"source" json:"source"`
	Target            []StructureMapGroupRuleTarget    `bson:"target,omitempty" json:"target,omitempty"`
	Rule              []StructureMapGroupRule          `bson:"rule,omitempty" json:"rule,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Documentation     *string                          `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRuleSource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           string      `bson:"context" json:"context"`
	Min               *int        `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string     `bson:"max,omitempty" json:"max,omitempty"`
	Type              *string     `bson:"type,omitempty" json:"type,omitempty"`
	Element           *string     `bson:"element,omitempty" json:"element,omitempty"`
	ListMode          *string     `bson:"listMode,omitempty" json:"listMode,omitempty"`
	Variable          *string     `bson:"variable,omitempty" json:"variable,omitempty"`
	Condition         *string     `bson:"condition,omitempty" json:"condition,omitempty"`
	Check             *string     `bson:"check,omitempty" json:"check,omitempty"`
}
type StructureMapGroupRuleTarget struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           *string                                `bson:"context,omitempty" json:"context,omitempty"`
	ContextType       *string                                `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Element           *string                                `bson:"element,omitempty" json:"element,omitempty"`
	Variable          *string                                `bson:"variable,omitempty" json:"variable,omitempty"`
	ListMode          []string                               `bson:"listMode,omitempty" json:"listMode,omitempty"`
	ListRuleId        *string                                `bson:"listRuleId,omitempty" json:"listRuleId,omitempty"`
	Transform         *string                                `bson:"transform,omitempty" json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type StructureMapGroupRuleDependent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Variable          []string    `bson:"variable" json:"variable"`
}
type OtherStructureMap StructureMap

// MarshalJSON marshals the given StructureMap as JSON into a byte slice
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

// UnmarshalStructureMap unmarshals a StructureMap.
func UnmarshalStructureMap(b []byte) (StructureMap, error) {
	var structureMap StructureMap
	if err := json.Unmarshal(b, &structureMap); err != nil {
		return structureMap, err
	}
	return structureMap, nil
}
