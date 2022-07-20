package fhir

import "encoding/json"

// ServiceDefinition is documented here http://hl7.org/fhir/StructureDefinition/ServiceDefinition
type ServiceDefinition struct {
	Id                  *string             `bson:"id" json:"id"`
	Meta                *Meta               `bson:"meta" json:"meta"`
	ImplicitRules       *string             `bson:"implicitRules" json:"implicitRules"`
	Language            *string             `bson:"language" json:"language"`
	Text                *Narrative          `bson:"text" json:"text"`
	Contained           []json.RawMessage   `bson:"contained" json:"contained"`
	Extension           []Extension         `bson:"extension" json:"extension"`
	ModifierExtension   []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Url                 *string             `bson:"url" json:"url"`
	Identifier          []Identifier        `bson:"identifier" json:"identifier"`
	Version             *string             `bson:"version" json:"version"`
	Name                *string             `bson:"name" json:"name"`
	Title               *string             `bson:"title" json:"title"`
	Status              PublicationStatus   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental        *bool               `bson:"experimental" json:"experimental"`
	Date                *string             `bson:"date" json:"date"`
	Publisher           *string             `bson:"publisher" json:"publisher"`
	Description         *string             `bson:"description" json:"description"`
	Purpose             *string             `bson:"purpose" json:"purpose"`
	Usage               *string             `bson:"usage" json:"usage"`
	ApprovalDate        *string             `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate      *string             `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod     *Period             `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext          []UsageContext      `bson:"useContext" json:"useContext"`
	Jurisdiction        []CodeableConcept   `bson:"jurisdiction" json:"jurisdiction"`
	Topic               []CodeableConcept   `bson:"topic" json:"topic"`
	Contributor         []Contributor       `bson:"contributor" json:"contributor"`
	Contact             []ContactDetail     `bson:"contact" json:"contact"`
	Copyright           *string             `bson:"copyright" json:"copyright"`
	RelatedArtifact     []RelatedArtifact   `bson:"relatedArtifact" json:"relatedArtifact"`
	Trigger             []TriggerDefinition `bson:"trigger" json:"trigger"`
	DataRequirement     []DataRequirement   `bson:"dataRequirement" json:"dataRequirement"`
	OperationDefinition *Reference          `bson:"operationDefinition" json:"operationDefinition"`
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

// UnmarshalServiceDefinition unmarshalls a ServiceDefinition.
func UnmarshalServiceDefinition(b []byte) (ServiceDefinition, error) {
	var serviceDefinition ServiceDefinition
	if err := json.Unmarshal(b, &serviceDefinition); err != nil {
		return serviceDefinition, err
	}
	return serviceDefinition, nil
}
