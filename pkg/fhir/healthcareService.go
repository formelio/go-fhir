package fhir

import "encoding/json"

// HealthcareService is documented here http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                     *string                          `bson:"id" json:"id"`
	Meta                   *Meta                            `bson:"meta" json:"meta"`
	ImplicitRules          *string                          `bson:"implicitRules" json:"implicitRules"`
	Language               *string                          `bson:"language" json:"language"`
	Text                   *Narrative                       `bson:"text" json:"text"`
	Contained              []json.RawMessage                `bson:"contained" json:"contained"`
	Extension              []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier             []Identifier                     `bson:"identifier" json:"identifier"`
	Active                 *bool                            `bson:"active" json:"active"`
	ProvidedBy             *Reference                       `bson:"providedBy" json:"providedBy"`
	Category               *CodeableConcept                 `bson:"category" json:"category"`
	Type                   []CodeableConcept                `bson:"type" json:"type"`
	Specialty              []CodeableConcept                `bson:"specialty" json:"specialty"`
	Location               []Reference                      `bson:"location" json:"location"`
	Name                   *string                          `bson:"name" json:"name"`
	Comment                *string                          `bson:"comment" json:"comment"`
	ExtraDetails           *string                          `bson:"extraDetails" json:"extraDetails"`
	Photo                  *Attachment                      `bson:"photo" json:"photo"`
	Telecom                []ContactPoint                   `bson:"telecom" json:"telecom"`
	CoverageArea           []Reference                      `bson:"coverageArea" json:"coverageArea"`
	ServiceProvisionCode   []CodeableConcept                `bson:"serviceProvisionCode" json:"serviceProvisionCode"`
	Eligibility            *CodeableConcept                 `bson:"eligibility" json:"eligibility"`
	EligibilityNote        *string                          `bson:"eligibilityNote" json:"eligibilityNote"`
	ProgramName            []string                         `bson:"programName" json:"programName"`
	Characteristic         []CodeableConcept                `bson:"characteristic" json:"characteristic"`
	ReferralMethod         []CodeableConcept                `bson:"referralMethod" json:"referralMethod"`
	AppointmentRequired    *bool                            `bson:"appointmentRequired" json:"appointmentRequired"`
	AvailableTime          []HealthcareServiceAvailableTime `bson:"availableTime" json:"availableTime"`
	NotAvailable           []HealthcareServiceNotAvailable  `bson:"notAvailable" json:"notAvailable"`
	AvailabilityExceptions *string                          `bson:"availabilityExceptions" json:"availabilityExceptions"`
	Endpoint               []Reference                      `bson:"endpoint" json:"endpoint"`
}
type HealthcareServiceAvailableTime struct {
	Id                 *string      `bson:"id" json:"id"`
	Extension          []Extension  `bson:"extension" json:"extension"`
	ModifierExtension  []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek" json:"daysOfWeek"`
	AllDay             *bool        `bson:"allDay" json:"allDay"`
	AvailableStartTime *string      `bson:"availableStartTime" json:"availableStartTime"`
	AvailableEndTime   *string      `bson:"availableEndTime" json:"availableEndTime"`
}
type HealthcareServiceNotAvailable struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       string      `bson:"description,omitempty" json:"description,omitempty"`
	During            *Period     `bson:"during" json:"during"`
}
type OtherHealthcareService HealthcareService

// MarshalJSON marshals the given HealthcareService as JSON into a byte slice
func (r HealthcareService) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHealthcareService
		ResourceType string `json:"resourceType"`
	}{
		OtherHealthcareService: OtherHealthcareService(r),
		ResourceType:           "HealthcareService",
	})
}

// UnmarshalHealthcareService unmarshalls a HealthcareService.
func UnmarshalHealthcareService(b []byte) (HealthcareService, error) {
	var healthcareService HealthcareService
	if err := json.Unmarshal(b, &healthcareService); err != nil {
		return healthcareService, err
	}
	return healthcareService, nil
}
