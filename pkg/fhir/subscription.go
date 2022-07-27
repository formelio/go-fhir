package fhir

import "encoding/json"

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `bson:"id" json:"id"`
	Meta              *Meta               `bson:"meta" json:"meta"`
	ImplicitRules     *string             `bson:"implicitRules" json:"implicitRules"`
	Language          *string             `bson:"language" json:"language"`
	Text              *Narrative          `bson:"text" json:"text"`
	RawContained      []json.RawMessage   `bson:"contained" json:"contained"`
	Contained         []IResource         `bson:"-" json:"-"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Status            SubscriptionStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Contact           []ContactPoint      `bson:"contact" json:"contact"`
	End               *string             `bson:"end" json:"end"`
	Reason            string              `bson:"reason,omitempty" json:"reason,omitempty"`
	Criteria          string              `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Error             *string             `bson:"error" json:"error"`
	Channel           SubscriptionChannel `bson:"channel,omitempty" json:"channel,omitempty"`
	Tag               []Coding            `bson:"tag" json:"tag"`
}
type SubscriptionChannel struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Type              SubscriptionChannelType `bson:"type,omitempty" json:"type,omitempty"`
	Endpoint          *string                 `bson:"endpoint" json:"endpoint"`
	Payload           *string                 `bson:"payload" json:"payload"`
	Header            []string                `bson:"header" json:"header"`
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
