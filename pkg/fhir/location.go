package fhir

import "encoding/json"

// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	Contained            []json.RawMessage `bson:"contained" json:"contained"`
	Extension            []Extension       `bson:"extension" json:"extension"`
	ModifierExtension    []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier      `bson:"identifier" json:"identifier"`
	Status               *LocationStatus   `bson:"status" json:"status"`
	OperationalStatus    *Coding           `bson:"operationalStatus" json:"operationalStatus"`
	Name                 *string           `bson:"name" json:"name"`
	Alias                []string          `bson:"alias" json:"alias"`
	Description          *string           `bson:"description" json:"description"`
	Mode                 *LocationMode     `bson:"mode" json:"mode"`
	Type                 *CodeableConcept  `bson:"type" json:"type"`
	Telecom              []ContactPoint    `bson:"telecom" json:"telecom"`
	Address              *Address          `bson:"address" json:"address"`
	PhysicalType         *CodeableConcept  `bson:"physicalType" json:"physicalType"`
	Position             *LocationPosition `bson:"position" json:"position"`
	ManagingOrganization *Reference        `bson:"managingOrganization" json:"managingOrganization"`
	PartOf               *Reference        `bson:"partOf" json:"partOf"`
	Endpoint             []Reference       `bson:"endpoint" json:"endpoint"`
}
type LocationPosition struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Longitude         float64     `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude          float64     `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Altitude          *float64    `bson:"altitude" json:"altitude"`
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

// UnmarshalLocation unmarshalls a Location.
func UnmarshalLocation(b []byte) (Location, error) {
	var location Location
	if err := json.Unmarshal(b, &location); err != nil {
		return location, err
	}
	return location, nil
}
