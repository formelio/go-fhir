package fhir

import "encoding/json"

// ServiceDefinition is documented here http://hl7.org/fhir/StructureDefinition/ServiceDefinition
type ServiceDefinition struct {
	Id                  *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                 *string             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier          []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version             *string             `bson:"version,omitempty" json:"version,omitempty"`
	Name                *string             `bson:"name,omitempty" json:"name,omitempty"`
	Title               *string             `bson:"title,omitempty" json:"title,omitempty"`
	Status              string              `bson:"status" json:"status"`
	Experimental        *bool               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                *string             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           *string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description         *string             `bson:"description,omitempty" json:"description,omitempty"`
	Purpose             *string             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage               *string             `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate        *string             `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate      *string             `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod     *Period             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext          []UsageContext      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic               []CodeableConcept   `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor         []Contributor       `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact             []ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright           *string             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact     []RelatedArtifact   `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Trigger             []TriggerDefinition `bson:"trigger,omitempty" json:"trigger,omitempty"`
	DataRequirement     []DataRequirement   `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
	OperationDefinition *Reference          `bson:"operationDefinition,omitempty" json:"operationDefinition,omitempty"`
}
type OtherServiceDefinition ServiceDefinition

// MarshalJSON marshals the given ServiceDefinition as JSON into a byte slice
func (r ServiceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherServiceDefinition: OtherServiceDefinition(r),
		ResourceType:           "ServiceDefinition",
	})
}

// UnmarshalServiceDefinition unmarshals a ServiceDefinition.
func UnmarshalServiceDefinition(b []byte) (ServiceDefinition, error) {
	var serviceDefinition ServiceDefinition
	if err := json.Unmarshal(b, &serviceDefinition); err != nil {
		return serviceDefinition, err
	}
	return serviceDefinition, nil
}
