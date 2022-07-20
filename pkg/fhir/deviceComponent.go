package fhir

import "encoding/json"

// DeviceComponent is documented here http://hl7.org/fhir/StructureDefinition/DeviceComponent
type DeviceComponent struct {
	Id                      *string                                  `bson:"id" json:"id"`
	Meta                    *Meta                                    `bson:"meta" json:"meta"`
	ImplicitRules           *string                                  `bson:"implicitRules" json:"implicitRules"`
	Language                *string                                  `bson:"language" json:"language"`
	Text                    *Narrative                               `bson:"text" json:"text"`
	Contained               []json.RawMessage                        `bson:"contained" json:"contained"`
	Extension               []Extension                              `bson:"extension" json:"extension"`
	ModifierExtension       []Extension                              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier              Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                    CodeableConcept                          `bson:"type,omitempty" json:"type,omitempty"`
	LastSystemChange        *string                                  `bson:"lastSystemChange" json:"lastSystemChange"`
	Source                  *Reference                               `bson:"source" json:"source"`
	Parent                  *Reference                               `bson:"parent" json:"parent"`
	OperationalStatus       []CodeableConcept                        `bson:"operationalStatus" json:"operationalStatus"`
	ParameterGroup          *CodeableConcept                         `bson:"parameterGroup" json:"parameterGroup"`
	MeasurementPrinciple    *MeasmntPrinciple                        `bson:"measurementPrinciple" json:"measurementPrinciple"`
	ProductionSpecification []DeviceComponentProductionSpecification `bson:"productionSpecification" json:"productionSpecification"`
	LanguageCode            *CodeableConcept                         `bson:"languageCode" json:"languageCode"`
}
type DeviceComponentProductionSpecification struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	SpecType          *CodeableConcept `bson:"specType" json:"specType"`
	ComponentId       *Identifier      `bson:"componentId" json:"componentId"`
	ProductionSpec    *string          `bson:"productionSpec" json:"productionSpec"`
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

// UnmarshalDeviceComponent unmarshalls a DeviceComponent.
func UnmarshalDeviceComponent(b []byte) (DeviceComponent, error) {
	var deviceComponent DeviceComponent
	if err := json.Unmarshal(b, &deviceComponent); err != nil {
		return deviceComponent, err
	}
	return deviceComponent, nil
}
