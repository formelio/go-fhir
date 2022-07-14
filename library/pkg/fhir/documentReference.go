package fhir

import "encoding/json"

// DocumentReference is documented here http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier                  `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                       `bson:"status" json:"status"`
	DocStatus         *string                      `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	Type              CodeableConcept              `bson:"type" json:"type"`
	Class             *CodeableConcept             `bson:"class,omitempty" json:"class,omitempty"`
	Created           *string                      `bson:"created,omitempty" json:"created,omitempty"`
	Indexed           string                       `bson:"indexed" json:"indexed"`
	Custodian         *Reference                   `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []DocumentReferenceRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel     []CodeableConcept            `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content           []DocumentReferenceContent   `bson:"content" json:"content"`
	Context           *DocumentReferenceContext    `bson:"context,omitempty" json:"context,omitempty"`
}
type DocumentReferenceRelatesTo struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Target            Reference   `bson:"target" json:"target"`
}
type DocumentReferenceContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Attachment        Attachment  `bson:"attachment" json:"attachment"`
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

// UnmarshalDocumentReference unmarshals a DocumentReference.
func UnmarshalDocumentReference(b []byte) (DocumentReference, error) {
	var documentReference DocumentReference
	if err := json.Unmarshal(b, &documentReference); err != nil {
		return documentReference, err
	}
	return documentReference, nil
}
