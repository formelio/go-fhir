package fhir

import (
	"bytes"
	"encoding/json"
)

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage                `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource                      `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []*Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                            `bson:"active,omitempty" json:"active,omitempty"`
	Period                 *Period                          `bson:"period,omitempty" json:"period,omitempty"`
	Practitioner           *Reference                       `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Organization           *Reference                       `bson:"organization,omitempty" json:"organization,omitempty"`
	Code                   []*CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Specialty              []*CodeableConcept               `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []*Reference                     `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService      []*Reference                     `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	Telecom                []*ContactPoint                  `bson:"telecom,omitempty" json:"telecom,omitempty"`
	AvailableTime          []*PractitionerRoleAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []*PractitionerRoleNotAvailable  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                          `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []*Reference                     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type PractitionerRoleAvailableTime struct {
	Id                 *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []*Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []*Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DaysOfWeek         []*DaysOfWeek `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *string       `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *string       `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}
type PractitionerRoleNotAvailable struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string       `bson:"description,omitempty" json:"description,omitempty"`
	During            *Period      `bson:"during,omitempty" json:"during,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherPractitionerRole
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
	return buffer.Bytes(), err
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
