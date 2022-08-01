package fhir

import "encoding/json"

// DeviceComponent is documented here http://hl7.org/fhir/StructureDefinition/DeviceComponent
type DeviceComponent struct {
	Id                      *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                               `bson:"text,omitempty" json:"text,omitempty"`
	RawContained            []json.RawMessage                        `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained               []IResource                              `bson:"-,omitempty" json:"-,omitempty"`
	Extension               []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                    CodeableConcept                          `bson:"type,omitempty" json:"type,omitempty"`
	LastSystemChange        *string                                  `bson:"lastSystemChange,omitempty" json:"lastSystemChange,omitempty"`
	Source                  *Reference                               `bson:"source,omitempty" json:"source,omitempty"`
	Parent                  *Reference                               `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus       []CodeableConcept                        `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	ParameterGroup          *CodeableConcept                         `bson:"parameterGroup,omitempty" json:"parameterGroup,omitempty"`
	MeasurementPrinciple    *MeasmntPrinciple                        `bson:"measurementPrinciple,omitempty" json:"measurementPrinciple,omitempty"`
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
