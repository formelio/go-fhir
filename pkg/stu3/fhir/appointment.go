package fhir

import (
	"bytes"
	"encoding/json"
)

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
	Id                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage        `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource              `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []*Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []*Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                AppointmentStatus        `bson:"status,omitempty" json:"status,omitempty"`
	ServiceCategory       *CodeableConcept         `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType           []*CodeableConcept       `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty             []*CodeableConcept       `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept         `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Reason                []*CodeableConcept       `bson:"reason,omitempty" json:"reason,omitempty"`
	Indication            []*Reference             `bson:"indication,omitempty" json:"indication,omitempty"`
	Priority              *int                     `bson:"priority,omitempty" json:"priority,omitempty"`
	Description           *string                  `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInformation []*Reference             `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Start                 *string                  `bson:"start,omitempty" json:"start,omitempty"`
	End                   *string                  `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration       *int                     `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot                  []*Reference             `bson:"slot,omitempty" json:"slot,omitempty"`
	Created               *string                  `bson:"created,omitempty" json:"created,omitempty"`
	Comment               *string                  `bson:"comment,omitempty" json:"comment,omitempty"`
	IncomingReferral      []*Reference             `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant           []AppointmentParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	RequestedPeriod       []*Period                `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}
type AppointmentParticipant struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []*CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Actor             *Reference           `bson:"actor,omitempty" json:"actor,omitempty"`
	Required          *ParticipantRequired `bson:"required,omitempty" json:"required,omitempty"`
	Status            ParticipationStatus  `bson:"status,omitempty" json:"status,omitempty"`
}

// OtherAppointment is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAppointment Appointment

// MarshalJSON marshals the given Appointment as JSON into a byte slice
func (r Appointment) MarshalJSON() ([]byte, error) {
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
		OtherAppointment
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Appointment
func (r *Appointment) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAppointment)(r)); err != nil {
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
func (r Appointment) GetResourceType() ResourceType {
	return ResourceTypeAppointment
}
