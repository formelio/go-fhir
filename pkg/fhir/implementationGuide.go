package fhir

import "encoding/json"

// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                          `bson:"url" json:"url"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                          `bson:"name" json:"name"`
	Status            string                          `bson:"status" json:"status"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                         `bson:"copyright,omitempty" json:"copyright,omitempty"`
	FhirVersion       *string                         `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Dependency        []ImplementationGuideDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	Package           []ImplementationGuidePackage    `bson:"package,omitempty" json:"package,omitempty"`
	Global            []ImplementationGuideGlobal     `bson:"global,omitempty" json:"global,omitempty"`
	Binary            []string                        `bson:"binary,omitempty" json:"binary,omitempty"`
	Page              *ImplementationGuidePage        `bson:"page,omitempty" json:"page,omitempty"`
}
type ImplementationGuideDependency struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Uri               string      `bson:"uri" json:"uri"`
}
type ImplementationGuidePackage struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                               `bson:"name" json:"name"`
	Description       *string                              `bson:"description,omitempty" json:"description,omitempty"`
	Resource          []ImplementationGuidePackageResource `bson:"resource" json:"resource"`
}
type ImplementationGuidePackageResource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Example           bool        `bson:"example" json:"example"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Acronym           *string     `bson:"acronym,omitempty" json:"acronym,omitempty"`
	ExampleFor        *Reference  `bson:"exampleFor,omitempty" json:"exampleFor,omitempty"`
}
type ImplementationGuideGlobal struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Profile           Reference   `bson:"profile" json:"profile"`
}
type ImplementationGuidePage struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Source            string                    `bson:"source" json:"source"`
	Title             string                    `bson:"title" json:"title"`
	Kind              string                    `bson:"kind" json:"kind"`
	Type              []string                  `bson:"type,omitempty" json:"type,omitempty"`
	Package           []string                  `bson:"package,omitempty" json:"package,omitempty"`
	Format            *string                   `bson:"format,omitempty" json:"format,omitempty"`
	Page              []ImplementationGuidePage `bson:"page,omitempty" json:"page,omitempty"`
}
type OtherImplementationGuide ImplementationGuide

// MarshalJSON marshals the given ImplementationGuide as JSON into a byte slice
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

// UnmarshalImplementationGuide unmarshals a ImplementationGuide.
func UnmarshalImplementationGuide(b []byte) (ImplementationGuide, error) {
	var implementationGuide ImplementationGuide
	if err := json.Unmarshal(b, &implementationGuide); err != nil {
		return implementationGuide, err
	}
	return implementationGuide, nil
}
