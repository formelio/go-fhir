package fhir

import "encoding/json"

// GraphDefinition is documented here http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string               `bson:"url" json:"url"`
	Version           *string               `bson:"version" json:"version"`
	Name              string                `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus     `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                 `bson:"experimental" json:"experimental"`
	Date              *string               `bson:"date" json:"date"`
	Publisher         *string               `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail       `bson:"contact" json:"contact"`
	Description       *string               `bson:"description" json:"description"`
	UseContext        []UsageContext        `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept     `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string               `bson:"purpose" json:"purpose"`
	Start             ResourceType          `bson:"start,omitempty" json:"start,omitempty"`
	Profile           *string               `bson:"profile" json:"profile"`
	Link              []GraphDefinitionLink `bson:"link" json:"link"`
}
type GraphDefinitionLink struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Path              string                      `bson:"path,omitempty" json:"path,omitempty"`
	SliceName         *string                     `bson:"sliceName" json:"sliceName"`
	Min               *int                        `bson:"min" json:"min"`
	Max               *string                     `bson:"max" json:"max"`
	Description       *string                     `bson:"description" json:"description"`
	Target            []GraphDefinitionLinkTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type GraphDefinitionLinkTarget struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	Type              ResourceType                           `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *string                                `bson:"profile" json:"profile"`
	Compartment       []GraphDefinitionLinkTargetCompartment `bson:"compartment" json:"compartment"`
	Link              []GraphDefinitionLink                  `bson:"link" json:"link"`
}
type GraphDefinitionLinkTargetCompartment struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Code              CompartmentType      `bson:"code,omitempty" json:"code,omitempty"`
	Rule              GraphCompartmentRule `bson:"rule,omitempty" json:"rule,omitempty"`
	Expression        *string              `bson:"expression" json:"expression"`
	Description       *string              `bson:"description" json:"description"`
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
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
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
