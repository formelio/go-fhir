package fhir

import "encoding/json"

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string            `bson:"status" json:"status"`
	WhenUsed          *Period           `bson:"whenUsed,omitempty" json:"whenUsed,omitempty"`
	RecordedOn        *string           `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Device            Reference         `bson:"device" json:"device"`
	Indication        []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	BodySite          *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherDeviceUseStatement DeviceUseStatement

// MarshalJSON marshals the given DeviceUseStatement as JSON into a byte slice
func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUseStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUseStatement: OtherDeviceUseStatement(r),
		ResourceType:            "DeviceUseStatement",
	})
}

// UnmarshalDeviceUseStatement unmarshals a DeviceUseStatement.
func UnmarshalDeviceUseStatement(b []byte) (DeviceUseStatement, error) {
	var deviceUseStatement DeviceUseStatement
	if err := json.Unmarshal(b, &deviceUseStatement); err != nil {
		return deviceUseStatement, err
	}
	return deviceUseStatement, nil
}
