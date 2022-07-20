package fhir

import "encoding/json"

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string                  `bson:"id" json:"id"`
	Meta              *Meta                    `bson:"meta" json:"meta"`
	ImplicitRules     *string                  `bson:"implicitRules" json:"implicitRules"`
	Language          *string                  `bson:"language" json:"language"`
	Text              *Narrative               `bson:"text" json:"text"`
	Contained         []json.RawMessage        `bson:"contained" json:"contained"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier             `bson:"identifier" json:"identifier"`
	Status            DeviceUseStatementStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject           Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	WhenUsed          *Period                  `bson:"whenUsed" json:"whenUsed"`
	TimingTiming      *Timing                  `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod      *Period                  `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime    *string                  `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	RecordedOn        *string                  `bson:"recordedOn" json:"recordedOn"`
	Source            *Reference               `bson:"source" json:"source"`
	Device            Reference                `bson:"device,omitempty" json:"device,omitempty"`
	Indication        []CodeableConcept        `bson:"indication" json:"indication"`
	BodySite          *CodeableConcept         `bson:"bodySite" json:"bodySite"`
	Note              []Annotation             `bson:"note" json:"note"`
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

// UnmarshalDeviceUseStatement unmarshalls a DeviceUseStatement.
func UnmarshalDeviceUseStatement(b []byte) (DeviceUseStatement, error) {
	var deviceUseStatement DeviceUseStatement
	if err := json.Unmarshal(b, &deviceUseStatement); err != nil {
		return deviceUseStatement, err
	}
	return deviceUseStatement, nil
}
