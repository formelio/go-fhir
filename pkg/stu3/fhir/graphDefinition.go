package fhir

import (
	"bytes"
	"encoding/json"
)

// GraphDefinition is documented here http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail       `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Start             ResourceType           `bson:"start,omitempty" json:"start,omitempty"`
	Profile           *string                `bson:"profile,omitempty" json:"profile,omitempty"`
	Link              []*GraphDefinitionLink `bson:"link,omitempty" json:"link,omitempty"`
}
type GraphDefinitionLink struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path              string                      `bson:"path,omitempty" json:"path,omitempty"`
	SliceName         *string                     `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	Min               *int                        `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string                     `bson:"max,omitempty" json:"max,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Target            []GraphDefinitionLinkTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type GraphDefinitionLinkTarget struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType                            `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *string                                 `bson:"profile,omitempty" json:"profile,omitempty"`
	Compartment       []*GraphDefinitionLinkTargetCompartment `bson:"compartment,omitempty" json:"compartment,omitempty"`
	Link              []*GraphDefinitionLink                  `bson:"link,omitempty" json:"link,omitempty"`
}
type GraphDefinitionLinkTargetCompartment struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CompartmentType      `bson:"code,omitempty" json:"code,omitempty"`
	Rule              GraphCompartmentRule `bson:"rule,omitempty" json:"rule,omitempty"`
	Expression        *string              `bson:"expression,omitempty" json:"expression,omitempty"`
	Description       *string              `bson:"description,omitempty" json:"description,omitempty"`
}

// OtherGraphDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherGraphDefinition GraphDefinition

// MarshalJSON marshals the given GraphDefinition as JSON into a byte slice
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
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
		OtherGraphDefinition
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into GraphDefinition
func (r *GraphDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherGraphDefinition)(r)); err != nil {
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
func (r GraphDefinition) GetResourceType() ResourceType {
	return ResourceTypeGraphDefinition
}
