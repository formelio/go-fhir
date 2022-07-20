package fhir

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        Identifier                `bson:"identifier" json:"identifier"`
	Type              CodeableConcept           `bson:"type" json:"type"`
	Unit              *CodeableConcept          `bson:"unit,omitempty" json:"unit,omitempty"`
	Source            *Reference                `bson:"source,omitempty" json:"source,omitempty"`
	Parent            *Reference                `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus *string                   `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Color             *string                   `bson:"color,omitempty" json:"color,omitempty"`
	Category          string                    `bson:"category" json:"category"`
	MeasurementPeriod *Timing                   `bson:"measurementPeriod,omitempty" json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibration `bson:"calibration,omitempty" json:"calibration,omitempty"`
}
type DeviceMetricCalibration struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *string     `bson:"type,omitempty" json:"type,omitempty"`
	State             *string     `bson:"state,omitempty" json:"state,omitempty"`
	Time              *string     `bson:"time,omitempty" json:"time,omitempty"`
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

// UnmarshalDeviceMetric unmarshals a DeviceMetric.
func UnmarshalDeviceMetric(b []byte) (DeviceMetric, error) {
	var deviceMetric DeviceMetric
	if err := json.Unmarshal(b, &deviceMetric); err != nil {
		return deviceMetric, err
	}
	return deviceMetric, nil
}
