package fhir

import "encoding/json"

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                         `bson:"id" json:"id"`
	Meta                   *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules          *string                         `bson:"implicitRules" json:"implicitRules"`
	Language               *string                         `bson:"language" json:"language"`
	Text                   *Narrative                      `bson:"text" json:"text"`
	Contained              []json.RawMessage               `bson:"contained" json:"contained"`
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
type OtherPractitionerRole PractitionerRole

// MarshalJSON marshals the given PractitionerRole as JSON into a byte slice
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

// UnmarshalPractitionerRole unmarshalls a PractitionerRole.
func UnmarshalPractitionerRole(b []byte) (PractitionerRole, error) {
	var practitionerRole PractitionerRole
	if err := json.Unmarshal(b, &practitionerRole); err != nil {
		return practitionerRole, err
	}
	return practitionerRole, nil
}
