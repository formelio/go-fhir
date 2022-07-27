package fhir

import "encoding/json"

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                         `bson:"id" json:"id"`
	Meta                   *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules          *string                         `bson:"implicitRules" json:"implicitRules"`
	Language               *string                         `bson:"language" json:"language"`
	Text                   *Narrative                      `bson:"text" json:"text"`
	RawContained           []json.RawMessage               `bson:"contained" json:"contained"`
	Contained              []IResource                     `bson:"-" json:"-"`
	Extension              []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Identifier             []Identifier                    `bson:"identifier" json:"identifier"`
	Active                 *bool                           `bson:"active" json:"active"`
	Period                 *Period                         `bson:"period" json:"period"`
	Practitioner           *Reference                      `bson:"practitioner" json:"practitioner"`
	Organization           *Reference                      `bson:"organization" json:"organization"`
	Code                   []CodeableConcept               `bson:"code" json:"code"`
	Specialty              []CodeableConcept               `bson:"specialty" json:"specialty"`
	Location               []Reference                     `bson:"location" json:"location"`
	HealthcareService      []Reference                     `bson:"healthcareService" json:"healthcareService"`
	Telecom                []ContactPoint                  `bson:"telecom" json:"telecom"`
	AvailableTime          []PractitionerRoleAvailableTime `bson:"availableTime" json:"availableTime"`
	NotAvailable           []PractitionerRoleNotAvailable  `bson:"notAvailable" json:"notAvailable"`
	AvailabilityExceptions *string                         `bson:"availabilityExceptions" json:"availabilityExceptions"`
	Endpoint               []Reference                     `bson:"endpoint" json:"endpoint"`
}
type PractitionerRoleAvailableTime struct {
	Id                 *string      `bson:"id" json:"id"`
	Extension          []Extension  `bson:"extension" json:"extension"`
	ModifierExtension  []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek" json:"daysOfWeek"`
	AllDay             *bool        `bson:"allDay" json:"allDay"`
	AvailableStartTime *string      `bson:"availableStartTime" json:"availableStartTime"`
	AvailableEndTime   *string      `bson:"availableEndTime" json:"availableEndTime"`
}
type PractitionerRoleNotAvailable struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       string      `bson:"description,omitempty" json:"description,omitempty"`
	During            *Period     `bson:"during" json:"during"`
}

// OtherPractitionerRole is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPractitionerRole PractitionerRole

// MarshalJSON marshals the given PractitionerRole as JSON into a byte slice
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
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
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

// UnmarshalJSON unmarshals the given byte slice into PractitionerRole
func (r *PractitionerRole) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPractitionerRole)(r)); err != nil {
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
func (r PractitionerRole) GetResourceType() ResourceType {
	return ResourceTypePractitionerRole
}
