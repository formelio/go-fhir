package fhir

import "encoding/json"

// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Kind              NamingSystemType       `bson:"kind,omitempty" json:"kind,omitempty"`
	Date              string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail        `bson:"contact" json:"contact"`
	Responsible       *string                `bson:"responsible" json:"responsible"`
	Type              *CodeableConcept       `bson:"type" json:"type"`
	Description       *string                `bson:"description" json:"description"`
	UseContext        []UsageContext         `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction" json:"jurisdiction"`
	Usage             *string                `bson:"usage" json:"usage"`
	UniqueId          []NamingSystemUniqueId `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
	ReplacedBy        *Reference             `bson:"replacedBy" json:"replacedBy"`
}
type NamingSystemUniqueId struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Type              NamingSystemIdentifierType `bson:"type,omitempty" json:"type,omitempty"`
	Value             string                     `bson:"value,omitempty" json:"value,omitempty"`
	Preferred         *bool                      `bson:"preferred" json:"preferred"`
	Comment           *string                    `bson:"comment" json:"comment"`
	Period            *Period                    `bson:"period" json:"period"`
}

// OtherNamingSystem is a helper type to use the default implementations of Marshall and Unmarshal
type OtherNamingSystem NamingSystem

// MarshalJSON marshals the given NamingSystem as JSON into a byte slice
func (r NamingSystem) MarshalJSON() ([]byte, error) {
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
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}

// UnmarshalJSON unmarshals the given byte slice into NamingSystem
func (r *NamingSystem) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherNamingSystem)(r)); err != nil {
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
func (r NamingSystem) GetResourceType() ResourceType {
	return ResourceTypeNamingSystem
}
