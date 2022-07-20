package fhir

import "encoding/json"

// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
	Id                   *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string           `bson:"status,omitempty" json:"status,omitempty"`
	OperationalStatus    *Coding           `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Name                 *string           `bson:"name,omitempty" json:"name,omitempty"`
	Alias                []string          `bson:"alias,omitempty" json:"alias,omitempty"`
	Description          *string           `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 *string           `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                 *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Telecom              []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address              *Address          `bson:"address,omitempty" json:"address,omitempty"`
	PhysicalType         *CodeableConcept  `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Position             *LocationPosition `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization *Reference        `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf               *Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Endpoint             []Reference       `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type LocationPosition struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Longitude         string      `bson:"longitude" json:"longitude"`
	Latitude          string      `bson:"latitude" json:"latitude"`
	Altitude          *string     `bson:"altitude,omitempty" json:"altitude,omitempty"`
}
type OtherLocation Location

// MarshalJSON marshals the given Location as JSON into a byte slice
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}

// UnmarshalLocation unmarshals a Location.
func UnmarshalLocation(b []byte) (Location, error) {
	var location Location
	if err := json.Unmarshal(b, &location); err != nil {
		return location, err
	}
	return location, nil
}
