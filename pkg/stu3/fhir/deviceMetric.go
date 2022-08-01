package fhir

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage              `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                    `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Unit              *CodeableConcept               `bson:"unit,omitempty" json:"unit,omitempty"`
	Source            *Reference                     `bson:"source,omitempty" json:"source,omitempty"`
	Parent            *Reference                     `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus *DeviceMetricOperationalStatus `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Color             *DeviceMetricColor             `bson:"color,omitempty" json:"color,omitempty"`
	Category          DeviceMetricCategory           `bson:"category,omitempty" json:"category,omitempty"`
	MeasurementPeriod *Timing                        `bson:"measurementPeriod,omitempty" json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibration      `bson:"calibration,omitempty" json:"calibration,omitempty"`
}
type DeviceMetricCalibration struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *DeviceMetricCalibrationType  `bson:"type,omitempty" json:"type,omitempty"`
	State             *DeviceMetricCalibrationState `bson:"state,omitempty" json:"state,omitempty"`
	Time              *string                       `bson:"time,omitempty" json:"time,omitempty"`
}

// OtherDeviceMetric is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDeviceMetric DeviceMetric

// MarshalJSON marshals the given DeviceMetric as JSON into a byte slice
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
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
		OtherDeviceMetric
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceMetric: OtherDeviceMetric(r),
		ResourceType:      "DeviceMetric",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DeviceMetric
func (r *DeviceMetric) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDeviceMetric)(r)); err != nil {
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
func (r DeviceMetric) GetResourceType() ResourceType {
	return ResourceTypeDeviceMetric
}
