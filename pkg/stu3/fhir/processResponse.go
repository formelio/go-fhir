package fhir

import "encoding/json"

// ProcessResponse is documented here http://hl7.org/fhir/StructureDefinition/ProcessResponse
type ProcessResponse struct {
	Id                   *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage            `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource                  `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string                      `bson:"status,omitempty" json:"status,omitempty"`
	Created              *string                      `bson:"created,omitempty" json:"created,omitempty"`
	Organization         *Reference                   `bson:"organization,omitempty" json:"organization,omitempty"`
	Request              *Reference                   `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              *CodeableConcept             `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                      `bson:"disposition,omitempty" json:"disposition,omitempty"`
	RequestProvider      *Reference                   `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization  *Reference                   `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Form                 *CodeableConcept             `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ProcessResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	Error                []CodeableConcept            `bson:"error,omitempty" json:"error,omitempty"`
	CommunicationRequest []Reference                  `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
}
type ProcessResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
}

// OtherProcessResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherProcessResponse ProcessResponse

// MarshalJSON marshals the given ProcessResponse as JSON into a byte slice
func (r ProcessResponse) MarshalJSON() ([]byte, error) {
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
		OtherProcessResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherProcessResponse: OtherProcessResponse(r),
		ResourceType:         "ProcessResponse",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ProcessResponse
func (r *ProcessResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherProcessResponse)(r)); err != nil {
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
func (r ProcessResponse) GetResourceType() ResourceType {
	return ResourceTypeProcessResponse
}
