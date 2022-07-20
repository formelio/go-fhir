package fhir

import "encoding/json"

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
type DocumentManifest struct {
	Id                *string                   `bson:"id" json:"id"`
	Meta              *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules     *string                   `bson:"implicitRules" json:"implicitRules"`
	Language          *string                   `bson:"language" json:"language"`
	Text              *Narrative                `bson:"text" json:"text"`
	Contained         []json.RawMessage         `bson:"contained" json:"contained"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	MasterIdentifier  *Identifier               `bson:"masterIdentifier" json:"masterIdentifier"`
	Identifier        []Identifier              `bson:"identifier" json:"identifier"`
	Status            DocumentReferenceStatus   `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept          `bson:"type" json:"type"`
	Subject           *Reference                `bson:"subject" json:"subject"`
	Created           *string                   `bson:"created" json:"created"`
	Author            []Reference               `bson:"author" json:"author"`
	Recipient         []Reference               `bson:"recipient" json:"recipient"`
	Source            *string                   `bson:"source" json:"source"`
	Description       *string                   `bson:"description" json:"description"`
	Content           []DocumentManifestContent `bson:"content,omitempty" json:"content,omitempty"`
	Related           []DocumentManifestRelated `bson:"related" json:"related"`
}
type DocumentManifestContent struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	PAttachment       *Attachment `bson:"pAttachment,omitempty" json:"pAttachment,omitempty"`
	PReference        *Reference  `bson:"pReference,omitempty" json:"pReference,omitempty"`
}
type DocumentManifestRelated struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier `bson:"identifier" json:"identifier"`
	Ref               *Reference  `bson:"ref" json:"ref"`
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

// UnmarshalDocumentManifest unmarshalls a DocumentManifest.
func UnmarshalDocumentManifest(b []byte) (DocumentManifest, error) {
	var documentManifest DocumentManifest
	if err := json.Unmarshal(b, &documentManifest); err != nil {
		return documentManifest, err
	}
	return documentManifest, nil
}
