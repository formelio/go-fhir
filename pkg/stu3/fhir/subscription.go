package fhir

import "encoding/json"

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource         `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            SubscriptionStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Contact           []ContactPoint      `bson:"contact,omitempty" json:"contact,omitempty"`
	End               *string             `bson:"end,omitempty" json:"end,omitempty"`
	Reason            string              `bson:"reason,omitempty" json:"reason,omitempty"`
	Criteria          string              `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Error             *string             `bson:"error,omitempty" json:"error,omitempty"`
	Channel           SubscriptionChannel `bson:"channel,omitempty" json:"channel,omitempty"`
	Tag               []Coding            `bson:"tag,omitempty" json:"tag,omitempty"`
}
type SubscriptionChannel struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              SubscriptionChannelType `bson:"type,omitempty" json:"type,omitempty"`
	Endpoint          *string                 `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Payload           *string                 `bson:"payload,omitempty" json:"payload,omitempty"`
	Header            []string                `bson:"header,omitempty" json:"header,omitempty"`
}

// OtherSubscription is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSubscription Subscription

// MarshalJSON marshals the given Subscription as JSON into a byte slice
func (r Subscription) MarshalJSON() ([]byte, error) {
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
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Subscription
func (r *Subscription) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSubscription)(r)); err != nil {
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
func (r Subscription) GetResourceType() ResourceType {
	return ResourceTypeSubscription
}
