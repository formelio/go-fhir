package fhir

import "encoding/json"

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id" json:"id"`
	Meta              *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules     *string                      `bson:"implicitRules" json:"implicitRules"`
	Language          *string                      `bson:"language" json:"language"`
	Text              *Narrative                   `bson:"text" json:"text"`
	RawContained      []json.RawMessage            `bson:"contained" json:"contained"`
	Contained         []IResource                  `bson:"-" json:"-"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	MasterIdentifier  *Identifier                  `bson:"masterIdentifier" json:"masterIdentifier"`
	Identifier        []Identifier                 `bson:"identifier" json:"identifier"`
	Status            DocumentReferenceStatus      `bson:"status,omitempty" json:"status,omitempty"`
	DocStatus         *CompositionStatus           `bson:"docStatus" json:"docStatus"`
	Type              CodeableConcept              `bson:"type,omitempty" json:"type,omitempty"`
	Class             *CodeableConcept             `bson:"class" json:"class"`
	Subject           *Reference                   `bson:"subject" json:"subject"`
	Created           *string                      `bson:"created" json:"created"`
	Indexed           string                       `bson:"indexed,omitempty" json:"indexed,omitempty"`
	Author            []Reference                  `bson:"author" json:"author"`
	Authenticator     *Reference                   `bson:"authenticator" json:"authenticator"`
	Custodian         *Reference                   `bson:"custodian" json:"custodian"`
	RelatesTo         []DocumentReferenceRelatesTo `bson:"relatesTo" json:"relatesTo"`
	Description       *string                      `bson:"description" json:"description"`
	SecurityLabel     []CodeableConcept            `bson:"securityLabel" json:"securityLabel"`
	Content           []DocumentReferenceContent   `bson:"content,omitempty" json:"content,omitempty"`
	Context           *DocumentReferenceContext    `bson:"context" json:"context"`
}
type DocumentReferenceRelatesTo struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Code              DocumentRelationshipType `bson:"code,omitempty" json:"code,omitempty"`
	Target            Reference                `bson:"target,omitempty" json:"target,omitempty"`
}
type DocumentReferenceContent struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Attachment        Attachment  `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format            *Coding     `bson:"format" json:"format"`
}
type DocumentReferenceContext struct {
	Id                *string                           `bson:"id" json:"id"`
	Extension         []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	Encounter         *Reference                        `bson:"encounter" json:"encounter"`
	Event             []CodeableConcept                 `bson:"event" json:"event"`
	Period            *Period                           `bson:"period" json:"period"`
	FacilityType      *CodeableConcept                  `bson:"facilityType" json:"facilityType"`
	PracticeSetting   *CodeableConcept                  `bson:"practiceSetting" json:"practiceSetting"`
	SourcePatientInfo *Reference                        `bson:"sourcePatientInfo" json:"sourcePatientInfo"`
	Related           []DocumentReferenceContextRelated `bson:"related" json:"related"`
}
type DocumentReferenceContextRelated struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier `bson:"identifier" json:"identifier"`
	Ref               *Reference  `bson:"ref" json:"ref"`
}

// OtherDocumentReference is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDocumentReference DocumentReference

// MarshalJSON marshals the given DocumentReference as JSON into a byte slice
func (r DocumentReference) MarshalJSON() ([]byte, error) {
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
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DocumentReference
func (r *DocumentReference) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDocumentReference)(r)); err != nil {
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
func (r DocumentReference) GetResourceType() ResourceType {
	return ResourceTypeDocumentReference
}
