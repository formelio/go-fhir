package fhir

import "encoding/json"

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string                  `bson:"id" json:"id"`
	Meta              *Meta                    `bson:"meta" json:"meta"`
	ImplicitRules     *string                  `bson:"implicitRules" json:"implicitRules"`
	Language          *string                  `bson:"language" json:"language"`
	Text              *Narrative               `bson:"text" json:"text"`
	RawContained      []json.RawMessage        `bson:"contained" json:"contained"`
	Contained         []IResource              `bson:"-" json:"-"`
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

// OtherDeviceUseStatement is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDeviceUseStatement DeviceUseStatement

// MarshalJSON marshals the given DeviceUseStatement as JSON into a byte slice
func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
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
		OtherDeviceUseStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUseStatement: OtherDeviceUseStatement(r),
		ResourceType:            "DeviceUseStatement",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DeviceUseStatement
func (r *DeviceUseStatement) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDeviceUseStatement)(r)); err != nil {
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
func (r DeviceUseStatement) GetResourceType() ResourceType {
	return ResourceTypeDeviceUseStatement
}
