package fhir

import "encoding/json"

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string              `bson:"status" json:"status"`
	Contact           []ContactPoint      `bson:"contact,omitempty" json:"contact,omitempty"`
	End               *string             `bson:"end,omitempty" json:"end,omitempty"`
	Reason            string              `bson:"reason" json:"reason"`
	Criteria          string              `bson:"criteria" json:"criteria"`
	Error             *string             `bson:"error,omitempty" json:"error,omitempty"`
	Channel           SubscriptionChannel `bson:"channel" json:"channel"`
	Tag               []Coding            `bson:"tag,omitempty" json:"tag,omitempty"`
}
type SubscriptionChannel struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Endpoint          *string     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Payload           *string     `bson:"payload,omitempty" json:"payload,omitempty"`
	Header            []string    `bson:"header,omitempty" json:"header,omitempty"`
}
type OtherSubscription Subscription

// MarshalJSON marshals the given Subscription as JSON into a byte slice
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// UnmarshalSubscription unmarshals a Subscription.
func UnmarshalSubscription(b []byte) (Subscription, error) {
	var subscription Subscription
	if err := json.Unmarshal(b, &subscription); err != nil {
		return subscription, err
	}
	return subscription, nil
}
