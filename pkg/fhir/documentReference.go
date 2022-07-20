package fhir

import "encoding/json"

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id" json:"id"`
	Meta              *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules     *string                      `bson:"implicitRules" json:"implicitRules"`
	Language          *string                      `bson:"language" json:"language"`
	Text              *Narrative                   `bson:"text" json:"text"`
	Contained         []json.RawMessage            `bson:"contained" json:"contained"`
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
type OtherDocumentReference DocumentReference

// MarshalJSON marshals the given DocumentReference as JSON into a byte slice
func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentReference: OtherDocumentReference(r),
		ResourceType:           "DocumentReference",
	})
}

// UnmarshalDocumentReference unmarshalls a DocumentReference.
func UnmarshalDocumentReference(b []byte) (DocumentReference, error) {
	var documentReference DocumentReference
	if err := json.Unmarshal(b, &documentReference); err != nil {
		return documentReference, err
	}
	return documentReference, nil
}
