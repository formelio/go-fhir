package fhir

import "encoding/json"

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                          `bson:"url" json:"url"`
	Name              string                          `bson:"name" json:"name"`
	Title             *string                         `bson:"title,omitempty" json:"title,omitempty"`
	Status            string                          `bson:"status" json:"status"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string                         `bson:"purpose,omitempty" json:"purpose,omitempty"`
	UseContext        []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Code              string                          `bson:"code" json:"code"`
	Search            bool                            `bson:"search" json:"search"`
	Resource          []CompartmentDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
}
type CompartmentDefinitionResource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Param             []string    `bson:"param,omitempty" json:"param,omitempty"`
	Documentation     *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
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

// UnmarshalCompartmentDefinition unmarshals a CompartmentDefinition.
func UnmarshalCompartmentDefinition(b []byte) (CompartmentDefinition, error) {
	var compartmentDefinition CompartmentDefinition
	if err := json.Unmarshal(b, &compartmentDefinition); err != nil {
		return compartmentDefinition, err
	}
	return compartmentDefinition, nil
}
