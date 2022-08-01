package fhir

import "encoding/json"

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage            `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                  `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier                  `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            DocumentReferenceStatus      `bson:"status,omitempty" json:"status,omitempty"`
	DocStatus         *CompositionStatus           `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	Type              CodeableConcept              `bson:"type,omitempty" json:"type,omitempty"`
	Class             *CodeableConcept             `bson:"class,omitempty" json:"class,omitempty"`
	Subject           *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Created           *string                      `bson:"created,omitempty" json:"created,omitempty"`
	Indexed           string                       `bson:"indexed,omitempty" json:"indexed,omitempty"`
	Author            []Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Authenticator     *Reference                   `bson:"authenticator,omitempty" json:"authenticator,omitempty"`
	Custodian         *Reference                   `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `bson:"content,omitempty" json:"content,omitempty"`
	Context           *DocumentReferenceContext    `bson:"context,omitempty" json:"context,omitempty"`
}
type DocumentReferenceRelatesTo struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `bson:"code,omitempty" json:"code,omitempty"`
	Target            Reference                `bson:"target,omitempty" json:"target,omitempty"`
}
type DocumentReferenceContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Attachment        Attachment  `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format            *Coding     `bson:"format,omitempty" json:"format,omitempty"`
}
type DocumentReferenceContext struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Encounter         *Reference                        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept                 `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period                           `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept                  `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept                  `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference                        `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []DocumentReferenceContextRelated `bson:"related,omitempty" json:"related,omitempty"`
}
type DocumentReferenceContextRelated struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref               *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
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
