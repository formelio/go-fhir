package fhir

import "encoding/json"

// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	Id                *string                         `bson:"id" json:"id"`
	Meta              *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules     *string                         `bson:"implicitRules" json:"implicitRules"`
	Language          *string                         `bson:"language" json:"language"`
	Text              *Narrative                      `bson:"text" json:"text"`
	RawContained      []json.RawMessage               `bson:"contained" json:"contained"`
	Contained         []IResource                     `bson:"-" json:"-"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                          `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                         `bson:"version" json:"version"`
	Name              string                          `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                           `bson:"experimental" json:"experimental"`
	Date              *string                         `bson:"date" json:"date"`
	Publisher         *string                         `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                 `bson:"contact" json:"contact"`
	Description       *string                         `bson:"description" json:"description"`
	UseContext        []UsageContext                  `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept               `bson:"jurisdiction" json:"jurisdiction"`
	Copyright         *string                         `bson:"copyright" json:"copyright"`
	FhirVersion       *string                         `bson:"fhirVersion" json:"fhirVersion"`
	Dependency        []ImplementationGuideDependency `bson:"dependency" json:"dependency"`
	Package           []ImplementationGuidePackage    `bson:"package" json:"package"`
	Global            []ImplementationGuideGlobal     `bson:"global" json:"global"`
	Binary            []string                        `bson:"binary" json:"binary"`
	Page              *ImplementationGuidePage        `bson:"page" json:"page"`
}
type ImplementationGuideDependency struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Type              GuideDependencyType `bson:"type,omitempty" json:"type,omitempty"`
	Uri               string              `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ImplementationGuidePackage struct {
	Id                *string                              `bson:"id" json:"id"`
	Extension         []Extension                          `bson:"extension" json:"extension"`
	ModifierExtension []Extension                          `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                               `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                              `bson:"description" json:"description"`
	Resource          []ImplementationGuidePackageResource `bson:"resource,omitempty" json:"resource,omitempty"`
}
type ImplementationGuidePackageResource struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Example           bool        `bson:"example,omitempty" json:"example,omitempty"`
	Name              *string     `bson:"name" json:"name"`
	Description       *string     `bson:"description" json:"description"`
	Acronym           *string     `bson:"acronym" json:"acronym"`
	SourceUri         *string     `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference   *Reference  `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	ExampleFor        *Reference  `bson:"exampleFor" json:"exampleFor"`
}
type ImplementationGuideGlobal struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Type              ResourceType `bson:"type,omitempty" json:"type,omitempty"`
	Profile           Reference    `bson:"profile,omitempty" json:"profile,omitempty"`
}
type ImplementationGuidePage struct {
	Id                *string                   `bson:"id" json:"id"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Source            string                    `bson:"source,omitempty" json:"source,omitempty"`
	Title             string                    `bson:"title,omitempty" json:"title,omitempty"`
	Kind              GuidePageKind             `bson:"kind,omitempty" json:"kind,omitempty"`
	Type              []ResourceType            `bson:"type" json:"type"`
	Package           []string                  `bson:"package" json:"package"`
	Format            *string                   `bson:"format" json:"format"`
	Page              []ImplementationGuidePage `bson:"page" json:"page"`
}

// OtherImplementationGuide is a helper type to use the default implementations of Marshall and Unmarshal
type OtherImplementationGuide ImplementationGuide

// MarshalJSON marshals the given ImplementationGuide as JSON into a byte slice
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
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
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ImplementationGuide
func (r *ImplementationGuide) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherImplementationGuide)(r)); err != nil {
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
func (r ImplementationGuide) GetResourceType() ResourceType {
	return ResourceTypeImplementationGuide
}
