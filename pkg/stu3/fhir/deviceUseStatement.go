package fhir

import "encoding/json"

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
type DeviceUseStatement struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage        `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource              `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            DeviceUseStatementStatus `bson:"status,omitempty" json:"status,omitempty"`
	Subject           Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	WhenUsed          *Period                  `bson:"whenUsed,omitempty" json:"whenUsed,omitempty"`
	TimingTiming      *Timing                  `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod      *Period                  `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime    *string                  `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	RecordedOn        *string                  `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Source            *Reference               `bson:"source,omitempty" json:"source,omitempty"`
	Device            Reference                `bson:"device,omitempty" json:"device,omitempty"`
	Indication        []CodeableConcept        `bson:"indication,omitempty" json:"indication,omitempty"`
	BodySite          *CodeableConcept         `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note              []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
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
