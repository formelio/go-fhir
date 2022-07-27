package fhir

import "encoding/json"

// DetectedIssue is documented here http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
	Id                *string                   `bson:"id" json:"id"`
	Meta              *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules     *string                   `bson:"implicitRules" json:"implicitRules"`
	Language          *string                   `bson:"language" json:"language"`
	Text              *Narrative                `bson:"text" json:"text"`
	RawContained      []json.RawMessage         `bson:"contained" json:"contained"`
	Contained         []IResource               `bson:"-" json:"-"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier               `bson:"identifier" json:"identifier"`
	Status            ObservationStatus         `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept          `bson:"category" json:"category"`
	Severity          *DetectedIssueSeverity    `bson:"severity" json:"severity"`
	Patient           *Reference                `bson:"patient" json:"patient"`
	Date              *string                   `bson:"date" json:"date"`
	Author            *Reference                `bson:"author" json:"author"`
	Implicated        []Reference               `bson:"implicated" json:"implicated"`
	Detail            *string                   `bson:"detail" json:"detail"`
	Reference         *string                   `bson:"reference" json:"reference"`
	Mitigation        []DetectedIssueMitigation `bson:"mitigation" json:"mitigation"`
}
type DetectedIssueMitigation struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Action            CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Date              *string         `bson:"date" json:"date"`
	Author            *Reference      `bson:"author" json:"author"`
}

// OtherDetectedIssue is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDetectedIssue DetectedIssue

// MarshalJSON marshals the given DetectedIssue as JSON into a byte slice
func (r DetectedIssue) MarshalJSON() ([]byte, error) {
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
		OtherDetectedIssue
		ResourceType string `json:"resourceType"`
	}{
		OtherDetectedIssue: OtherDetectedIssue(r),
		ResourceType:       "DetectedIssue",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DetectedIssue
func (r *DetectedIssue) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDetectedIssue)(r)); err != nil {
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
func (r DetectedIssue) GetResourceType() ResourceType {
	return ResourceTypeDetectedIssue
}
