package fhir

import "encoding/json"

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                          `bson:"url,omitempty" json:"url,omitempty"`
	Name              string                          `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                         `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string                         `bson:"purpose,omitempty" json:"purpose,omitempty"`
	UseContext        []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Code              CompartmentType                 `bson:"code,omitempty" json:"code,omitempty"`
	Search            bool                            `bson:"search,omitempty" json:"search,omitempty"`
	Resource          []CompartmentDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
}
type CompartmentDefinitionResource struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              ResourceType `bson:"code,omitempty" json:"code,omitempty"`
	Param             []string     `bson:"param,omitempty" json:"param,omitempty"`
	Documentation     *string      `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// OtherCompartmentDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCompartmentDefinition CompartmentDefinition

// MarshalJSON marshals the given CompartmentDefinition as JSON into a byte slice
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
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
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
}

// UnmarshalJSON unmarshals the given byte slice into CompartmentDefinition
func (r *CompartmentDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCompartmentDefinition)(r)); err != nil {
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
func (r CompartmentDefinition) GetResourceType() ResourceType {
	return ResourceTypeCompartmentDefinition
}
