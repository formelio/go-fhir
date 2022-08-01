package fhir

import "encoding/json"

// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
	Id                   *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *LocationStatus   `bson:"status,omitempty" json:"status,omitempty"`
	OperationalStatus    *Coding           `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Name                 *string           `bson:"name,omitempty" json:"name,omitempty"`
	Alias                []string          `bson:"alias,omitempty" json:"alias,omitempty"`
	Description          *string           `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 *LocationMode     `bson:"mode,omitempty" json:"mode,omitempty"`
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
	Longitude         float64     `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude          float64     `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Altitude          *float64    `bson:"altitude,omitempty" json:"altitude,omitempty"`
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
