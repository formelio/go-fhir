package fhir

import "encoding/json"

// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	RawContained         []json.RawMessage `bson:"contained" json:"contained"`
	Contained            []IResource       `bson:"-" json:"-"`
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

// OtherLocation is a helper type to use the default implementations of Marshall and Unmarshal
type OtherLocation Location

// MarshalJSON marshals the given Location as JSON into a byte slice
func (r Location) MarshalJSON() ([]byte, error) {
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
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Location
func (r *Location) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherLocation)(r)); err != nil {
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
func (r Location) GetResourceType() ResourceType {
	return ResourceTypeLocation
}
