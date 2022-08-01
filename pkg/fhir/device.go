package fhir

import "encoding/json"

// Device is documented here http://hl7.org/fhir/StructureDefinition/Device
type Device struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Udi               *DeviceUdi        `bson:"udi,omitempty" json:"udi,omitempty"`
	Status            *FHIRDeviceStatus `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	LotNumber         *string           `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Manufacturer      *string           `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	ManufactureDate   *string           `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	ExpirationDate    *string           `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Model             *string           `bson:"model,omitempty" json:"model,omitempty"`
	Version           *string           `bson:"version,omitempty" json:"version,omitempty"`
	Patient           *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	Owner             *Reference        `bson:"owner,omitempty" json:"owner,omitempty"`
	Contact           []ContactPoint    `bson:"contact,omitempty" json:"contact,omitempty"`
	Location          *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	Url               *string           `bson:"url,omitempty" json:"url,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	Safety            []CodeableConcept `bson:"safety,omitempty" json:"safety,omitempty"`
}
type DeviceUdi struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DeviceIdentifier  *string       `bson:"deviceIdentifier,omitempty" json:"deviceIdentifier,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Jurisdiction      *string       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	CarrierHRF        *string       `bson:"carrierHRF,omitempty" json:"carrierHRF,omitempty"`
	CarrierAIDC       *string       `bson:"carrierAIDC,omitempty" json:"carrierAIDC,omitempty"`
	Issuer            *string       `bson:"issuer,omitempty" json:"issuer,omitempty"`
	EntryType         *UDIEntryType `bson:"entryType,omitempty" json:"entryType,omitempty"`
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
