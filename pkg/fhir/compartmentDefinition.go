package fhir

import "encoding/json"

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
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
	Name              string                          `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                         `bson:"title" json:"title"`
	Status            PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                           `bson:"experimental" json:"experimental"`
	Date              *string                         `bson:"date" json:"date"`
	Publisher         *string                         `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                 `bson:"contact" json:"contact"`
	Description       *string                         `bson:"description" json:"description"`
	Purpose           *string                         `bson:"purpose" json:"purpose"`
	UseContext        []UsageContext                  `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept               `bson:"jurisdiction" json:"jurisdiction"`
	Code              CompartmentType                 `bson:"code,omitempty" json:"code,omitempty"`
	Search            bool                            `bson:"search,omitempty" json:"search,omitempty"`
	Resource          []CompartmentDefinitionResource `bson:"resource" json:"resource"`
}
type CompartmentDefinitionResource struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Code              ResourceType `bson:"code,omitempty" json:"code,omitempty"`
	Param             []string     `bson:"param" json:"param"`
	Documentation     *string      `bson:"documentation" json:"documentation"`
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
