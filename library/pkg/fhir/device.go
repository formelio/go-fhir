package fhir

import "encoding/json"

// Device is documented here http://hl7.org/fhir/StructureDefinition/Device
type Device struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Udi               *DeviceUdi        `bson:"udi,omitempty" json:"udi,omitempty"`
	Status            *string           `bson:"status,omitempty" json:"status,omitempty"`
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
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DeviceIdentifier  *string     `bson:"deviceIdentifier,omitempty" json:"deviceIdentifier,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Jurisdiction      *string     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	CarrierHRF        *string     `bson:"carrierHRF,omitempty" json:"carrierHRF,omitempty"`
	CarrierAIDC       *string     `bson:"carrierAIDC,omitempty" json:"carrierAIDC,omitempty"`
	Issuer            *string     `bson:"issuer,omitempty" json:"issuer,omitempty"`
	EntryType         *string     `bson:"entryType,omitempty" json:"entryType,omitempty"`
}
type OtherDevice Device

// MarshalJSON marshals the given Device as JSON into a byte slice
func (r Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDevice
		ResourceType string `json:"resourceType"`
	}{
		OtherDevice:  OtherDevice(r),
		ResourceType: "Device",
	})
}

// UnmarshalDevice unmarshals a Device.
func UnmarshalDevice(b []byte) (Device, error) {
	var device Device
	if err := json.Unmarshal(b, &device); err != nil {
		return device, err
	}
	return device, nil
}
