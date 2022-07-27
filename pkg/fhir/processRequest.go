package fhir

import "encoding/json"

// ProcessRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcessRequest
type ProcessRequest struct {
	Id                *string              `bson:"id" json:"id"`
	Meta              *Meta                `bson:"meta" json:"meta"`
	ImplicitRules     *string              `bson:"implicitRules" json:"implicitRules"`
	Language          *string              `bson:"language" json:"language"`
	Text              *Narrative           `bson:"text" json:"text"`
	RawContained      []json.RawMessage    `bson:"contained" json:"contained"`
	Contained         []IResource          `bson:"-" json:"-"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier         `bson:"identifier" json:"identifier"`
	Status            *string              `bson:"status" json:"status"`
	Action            *ActionList          `bson:"action" json:"action"`
	Target            *Reference           `bson:"target" json:"target"`
	Created           *string              `bson:"created" json:"created"`
	Provider          *Reference           `bson:"provider" json:"provider"`
	Organization      *Reference           `bson:"organization" json:"organization"`
	Request           *Reference           `bson:"request" json:"request"`
	Response          *Reference           `bson:"response" json:"response"`
	Nullify           *bool                `bson:"nullify" json:"nullify"`
	Reference         *string              `bson:"reference" json:"reference"`
	Item              []ProcessRequestItem `bson:"item" json:"item"`
	Include           []string             `bson:"include" json:"include"`
	Exclude           []string             `bson:"exclude" json:"exclude"`
	Period            *Period              `bson:"period" json:"period"`
}
type ProcessRequestItem struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    int         `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
}

// OtherProcessRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherProcessRequest ProcessRequest

// MarshalJSON marshals the given ProcessRequest as JSON into a byte slice
func (r ProcessRequest) MarshalJSON() ([]byte, error) {
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
		OtherProcessRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherProcessRequest: OtherProcessRequest(r),
		ResourceType:        "ProcessRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ProcessRequest
func (r *ProcessRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherProcessRequest)(r)); err != nil {
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
func (r ProcessRequest) GetResourceType() ResourceType {
	return ResourceTypeProcessRequest
}
