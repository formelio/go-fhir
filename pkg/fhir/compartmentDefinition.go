package fhir

import "encoding/json"

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	Id                *string                         `bson:"id" json:"id"`
	Meta              *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules     *string                         `bson:"implicitRules" json:"implicitRules"`
	Language          *string                         `bson:"language" json:"language"`
	Text              *Narrative                      `bson:"text" json:"text"`
	Contained         []json.RawMessage               `bson:"contained" json:"contained"`
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
type OtherCompartmentDefinition CompartmentDefinition

// MarshalJSON marshals the given CompartmentDefinition as JSON into a byte slice
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
}

// UnmarshalCompartmentDefinition unmarshalls a CompartmentDefinition.
func UnmarshalCompartmentDefinition(b []byte) (CompartmentDefinition, error) {
	var compartmentDefinition CompartmentDefinition
	if err := json.Unmarshal(b, &compartmentDefinition); err != nil {
		return compartmentDefinition, err
	}
	return compartmentDefinition, nil
}
