package fhir

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                *string                        `bson:"id" json:"id"`
	Meta              *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules     *string                        `bson:"implicitRules" json:"implicitRules"`
	Language          *string                        `bson:"language" json:"language"`
	Text              *Narrative                     `bson:"text" json:"text"`
	RawContained      []json.RawMessage              `bson:"contained" json:"contained"`
	Contained         []IResource                    `bson:"-" json:"-"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Unit              *CodeableConcept               `bson:"unit" json:"unit"`
	Source            *Reference                     `bson:"source" json:"source"`
	Parent            *Reference                     `bson:"parent" json:"parent"`
	OperationalStatus *DeviceMetricOperationalStatus `bson:"operationalStatus" json:"operationalStatus"`
	Color             *DeviceMetricColor             `bson:"color" json:"color"`
	Category          DeviceMetricCategory           `bson:"category,omitempty" json:"category,omitempty"`
	MeasurementPeriod *Timing                        `bson:"measurementPeriod" json:"measurementPeriod"`
	Calibration       []DeviceMetricCalibration      `bson:"calibration" json:"calibration"`
}
type DeviceMetricCalibration struct {
	Id                *string                       `bson:"id" json:"id"`
	Extension         []Extension                   `bson:"extension" json:"extension"`
	ModifierExtension []Extension                   `bson:"modifierExtension" json:"modifierExtension"`
	Type              *DeviceMetricCalibrationType  `bson:"type" json:"type"`
	State             *DeviceMetricCalibrationState `bson:"state" json:"state"`
	Time              *string                       `bson:"time" json:"time"`
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
