package fhir

import (
	"bytes"
	"encoding/json"
)

// EnrollmentRequest is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string           `bson:"status,omitempty" json:"status,omitempty"`
	Created           *string           `bson:"created,omitempty" json:"created,omitempty"`
	Insurer           *Reference        `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider          *Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization      *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Subject           *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Coverage          *Reference        `bson:"coverage,omitempty" json:"coverage,omitempty"`
}

// OtherEnrollmentRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEnrollmentRequest EnrollmentRequest

// MarshalJSON marshals the given EnrollmentRequest as JSON into a byte slice
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
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
		OtherEnrollmentRequest
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into EnrollmentRequest
func (r *EnrollmentRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEnrollmentRequest)(r)); err != nil {
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
func (r EnrollmentRequest) GetResourceType() ResourceType {
	return ResourceTypeEnrollmentRequest
}
