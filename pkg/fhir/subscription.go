package fhir

import "encoding/json"

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `bson:"id" json:"id"`
	Meta              *Meta               `bson:"meta" json:"meta"`
	ImplicitRules     *string             `bson:"implicitRules" json:"implicitRules"`
	Language          *string             `bson:"language" json:"language"`
	Text              *Narrative          `bson:"text" json:"text"`
	Contained         []json.RawMessage   `bson:"contained" json:"contained"`
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

// UnmarshalSubscription unmarshalls a Subscription.
func UnmarshalSubscription(b []byte) (Subscription, error) {
	var subscription Subscription
	if err := json.Unmarshal(b, &subscription); err != nil {
		return subscription, err
	}
	return subscription, nil
}
