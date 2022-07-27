package fhir

import "encoding/json"

// Device is documented here http://hl7.org/fhir/StructureDefinition/Device
type Device struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Udi               *DeviceUdi        `bson:"udi" json:"udi"`
	Status            *FHIRDeviceStatus `bson:"status" json:"status"`
	Type              *CodeableConcept  `bson:"type" json:"type"`
	LotNumber         *string           `bson:"lotNumber" json:"lotNumber"`
	Manufacturer      *string           `bson:"manufacturer" json:"manufacturer"`
	ManufactureDate   *string           `bson:"manufactureDate" json:"manufactureDate"`
	ExpirationDate    *string           `bson:"expirationDate" json:"expirationDate"`
	Model             *string           `bson:"model" json:"model"`
	Version           *string           `bson:"version" json:"version"`
	Patient           *Reference        `bson:"patient" json:"patient"`
	Owner             *Reference        `bson:"owner" json:"owner"`
	Contact           []ContactPoint    `bson:"contact" json:"contact"`
	Location          *Reference        `bson:"location" json:"location"`
	Url               *string           `bson:"url" json:"url"`
	Note              []Annotation      `bson:"note" json:"note"`
	Safety            []CodeableConcept `bson:"safety" json:"safety"`
}
type DeviceUdi struct {
	Id                *string       `bson:"id" json:"id"`
	Extension         []Extension   `bson:"extension" json:"extension"`
	ModifierExtension []Extension   `bson:"modifierExtension" json:"modifierExtension"`
	DeviceIdentifier  *string       `bson:"deviceIdentifier" json:"deviceIdentifier"`
	Name              *string       `bson:"name" json:"name"`
	Jurisdiction      *string       `bson:"jurisdiction" json:"jurisdiction"`
	CarrierHRF        *string       `bson:"carrierHRF" json:"carrierHRF"`
	CarrierAIDC       *string       `bson:"carrierAIDC" json:"carrierAIDC"`
	Issuer            *string       `bson:"issuer" json:"issuer"`
	EntryType         *UDIEntryType `bson:"entryType" json:"entryType"`
}

// OtherDevice is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDevice Device

// MarshalJSON marshals the given Device as JSON into a byte slice
func (r Device) MarshalJSON() ([]byte, error) {
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
		OtherDevice
		ResourceType string `json:"resourceType"`
	}{
		OtherDevice:  OtherDevice(r),
		ResourceType: "Device",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Device
func (r *Device) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDevice)(r)); err != nil {
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
func (r Device) GetResourceType() ResourceType {
	return ResourceTypeDevice
}
