package fhir

import "encoding/json"

// DeviceComponent is documented here http://hl7.org/fhir/StructureDefinition/DeviceComponent
type DeviceComponent struct {
	Id                      *string                                  `bson:"id" json:"id"`
	Meta                    *Meta                                    `bson:"meta" json:"meta"`
	ImplicitRules           *string                                  `bson:"implicitRules" json:"implicitRules"`
	Language                *string                                  `bson:"language" json:"language"`
	Text                    *Narrative                               `bson:"text" json:"text"`
	RawContained            []json.RawMessage                        `bson:"contained" json:"contained"`
	Contained               []IResource                              `bson:"-" json:"-"`
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

// OtherDeviceComponent is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDeviceComponent DeviceComponent

// MarshalJSON marshals the given DeviceComponent as JSON into a byte slice
func (r DeviceComponent) MarshalJSON() ([]byte, error) {
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
		OtherDeviceComponent
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceComponent: OtherDeviceComponent(r),
		ResourceType:         "DeviceComponent",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DeviceComponent
func (r *DeviceComponent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDeviceComponent)(r)); err != nil {
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
func (r DeviceComponent) GetResourceType() ResourceType {
	return ResourceTypeDeviceComponent
}
