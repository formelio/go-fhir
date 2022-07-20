package fhir

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                *string                        `bson:"id" json:"id"`
	Meta              *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules     *string                        `bson:"implicitRules" json:"implicitRules"`
	Language          *string                        `bson:"language" json:"language"`
	Text              *Narrative                     `bson:"text" json:"text"`
	Contained         []json.RawMessage              `bson:"contained" json:"contained"`
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
type OtherDeviceMetric DeviceMetric

// MarshalJSON marshals the given DeviceMetric as JSON into a byte slice
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetric
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceMetric: OtherDeviceMetric(r),
		ResourceType:      "DeviceMetric",
	})
}

// UnmarshalDeviceMetric unmarshalls a DeviceMetric.
func UnmarshalDeviceMetric(b []byte) (DeviceMetric, error) {
	var deviceMetric DeviceMetric
	if err := json.Unmarshal(b, &deviceMetric); err != nil {
		return deviceMetric, err
	}
	return deviceMetric, nil
}
