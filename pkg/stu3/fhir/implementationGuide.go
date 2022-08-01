package fhir

import (
	"bytes"
	"encoding/json"
)

// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage                `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                      `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                           `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                          `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	FhirVersion       *string                          `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Dependency        []*ImplementationGuideDependency `bson:"dependency,omitempty" json:"dependency,omitempty"`
	Package           []*ImplementationGuidePackage    `bson:"package,omitempty" json:"package,omitempty"`
	Global            []*ImplementationGuideGlobal     `bson:"global,omitempty" json:"global,omitempty"`
	Binary            []*string                        `bson:"binary,omitempty" json:"binary,omitempty"`
	Page              *ImplementationGuidePage         `bson:"page,omitempty" json:"page,omitempty"`
}
type ImplementationGuideDependency struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              GuideDependencyType `bson:"type,omitempty" json:"type,omitempty"`
	Uri               string              `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ImplementationGuidePackage struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                               `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                              `bson:"description,omitempty" json:"description,omitempty"`
	Resource          []ImplementationGuidePackageResource `bson:"resource,omitempty" json:"resource,omitempty"`
}
type ImplementationGuidePackageResource struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Example           bool         `bson:"example,omitempty" json:"example,omitempty"`
	Name              *string      `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Acronym           *string      `bson:"acronym,omitempty" json:"acronym,omitempty"`
	SourceUri         *string      `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference   *Reference   `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	ExampleFor        *Reference   `bson:"exampleFor,omitempty" json:"exampleFor,omitempty"`
}
type ImplementationGuideGlobal struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType `bson:"type,omitempty" json:"type,omitempty"`
	Profile           Reference    `bson:"profile,omitempty" json:"profile,omitempty"`
}
type ImplementationGuidePage struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Source            string                     `bson:"source,omitempty" json:"source,omitempty"`
	Title             string                     `bson:"title,omitempty" json:"title,omitempty"`
	Kind              GuidePageKind              `bson:"kind,omitempty" json:"kind,omitempty"`
	Type              []*ResourceType            `bson:"type,omitempty" json:"type,omitempty"`
	Package           []*string                  `bson:"package,omitempty" json:"package,omitempty"`
	Format            *string                    `bson:"format,omitempty" json:"format,omitempty"`
	Page              []*ImplementationGuidePage `bson:"page,omitempty" json:"page,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherImplementationGuide
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
	return buffer.Bytes(), err
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
