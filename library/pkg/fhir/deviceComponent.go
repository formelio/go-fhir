package fhir

import "encoding/json"

// DeviceComponent is documented here http://hl7.org/fhir/StructureDefinition/DeviceComponent
type DeviceComponent struct {
	Id                      *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                               `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              Identifier                               `bson:"identifier" json:"identifier"`
	Type                    CodeableConcept                          `bson:"type" json:"type"`
	LastSystemChange        *string                                  `bson:"lastSystemChange,omitempty" json:"lastSystemChange,omitempty"`
	Source                  *Reference                               `bson:"source,omitempty" json:"source,omitempty"`
	Parent                  *Reference                               `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus       []CodeableConcept                        `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	ParameterGroup          *CodeableConcept                         `bson:"parameterGroup,omitempty" json:"parameterGroup,omitempty"`
	MeasurementPrinciple    *string                                  `bson:"measurementPrinciple,omitempty" json:"measurementPrinciple,omitempty"`
	ProductionSpecification []DeviceComponentProductionSpecification `bson:"productionSpecification,omitempty" json:"productionSpecification,omitempty"`
	LanguageCode            *CodeableConcept                         `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
}
type DeviceComponentProductionSpecification struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SpecType          *CodeableConcept `bson:"specType,omitempty" json:"specType,omitempty"`
	ComponentId       *Identifier      `bson:"componentId,omitempty" json:"componentId,omitempty"`
	ProductionSpec    *string          `bson:"productionSpec,omitempty" json:"productionSpec,omitempty"`
}
type OtherDeviceComponent DeviceComponent

// MarshalJSON marshals the given DeviceComponent as JSON into a byte slice
func (r DeviceComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceComponent
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceComponent: OtherDeviceComponent(r),
		ResourceType:         "DeviceComponent",
	})
}

// UnmarshalDeviceComponent unmarshals a DeviceComponent.
func UnmarshalDeviceComponent(b []byte) (DeviceComponent, error) {
	var deviceComponent DeviceComponent
	if err := json.Unmarshal(b, &deviceComponent); err != nil {
		return deviceComponent, err
	}
	return deviceComponent, nil
}
