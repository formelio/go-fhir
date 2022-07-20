package fhir

import "encoding/json"

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
type DocumentManifest struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier               `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                    `bson:"status" json:"status"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Created           *string                   `bson:"created,omitempty" json:"created,omitempty"`
	Source            *string                   `bson:"source,omitempty" json:"source,omitempty"`
	Description       *string                   `bson:"description,omitempty" json:"description,omitempty"`
	Content           []DocumentManifestContent `bson:"content" json:"content"`
	Related           []DocumentManifestRelated `bson:"related,omitempty" json:"related,omitempty"`
}
type DocumentManifestContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type DocumentManifestRelated struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref               *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}
type OtherDocumentManifest DocumentManifest

// MarshalJSON marshals the given DocumentManifest as JSON into a byte slice
func (r DocumentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentManifest: OtherDocumentManifest(r),
		ResourceType:          "DocumentManifest",
	})
}

// UnmarshalDocumentManifest unmarshals a DocumentManifest.
func UnmarshalDocumentManifest(b []byte) (DocumentManifest, error) {
	var documentManifest DocumentManifest
	if err := json.Unmarshal(b, &documentManifest); err != nil {
		return documentManifest, err
	}
	return documentManifest, nil
}
