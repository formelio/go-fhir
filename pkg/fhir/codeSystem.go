package fhir

import "encoding/json"

// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string              `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string              `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string              `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Status            string               `bson:"status" json:"status"`
	Experimental      *bool                `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string              `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail      `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string              `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext       `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept    `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string              `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CaseSensitive     *bool                `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	ValueSet          *string              `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	HierarchyMeaning  *string              `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                `bson:"compositional,omitempty" json:"compositional,omitempty"`
	VersionNeeded     *bool                `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	Content           string               `bson:"content" json:"content"`
	Count             *int                 `bson:"count,omitempty" json:"count,omitempty"`
	Filter            []CodeSystemFilter   `bson:"filter,omitempty" json:"filter,omitempty"`
	Property          []CodeSystemProperty `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []CodeSystemConcept  `bson:"concept,omitempty" json:"concept,omitempty"`
}
type CodeSystemFilter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Operator          []string    `bson:"operator" json:"operator"`
	Value             string      `bson:"value" json:"value"`
}
type CodeSystemProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Type              string      `bson:"type" json:"type"`
}
type CodeSystemConcept struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                         `bson:"code" json:"code"`
	Display           *string                        `bson:"display,omitempty" json:"display,omitempty"`
	Definition        *string                        `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []CodeSystemConcept            `bson:"concept,omitempty" json:"concept,omitempty"`
}
type CodeSystemConceptDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
	Value             string      `bson:"value" json:"value"`
}
type CodeSystemConceptProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
}
type OtherCodeSystem CodeSystem

// MarshalJSON marshals the given CodeSystem as JSON into a byte slice
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

// UnmarshalCodeSystem unmarshals a CodeSystem.
func UnmarshalCodeSystem(b []byte) (CodeSystem, error) {
	var codeSystem CodeSystem
	if err := json.Unmarshal(b, &codeSystem); err != nil {
		return codeSystem, err
	}
	return codeSystem, nil
}
