package fhir

import "encoding/json"

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
type DocumentManifest struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource               `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier               `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            DocumentReferenceStatus   `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Subject           *Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	Created           *string                   `bson:"created,omitempty" json:"created,omitempty"`
	Author            []Reference               `bson:"author,omitempty" json:"author,omitempty"`
	Recipient         []Reference               `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Source            *string                   `bson:"source,omitempty" json:"source,omitempty"`
	Description       *string                   `bson:"description,omitempty" json:"description,omitempty"`
	Content           []DocumentManifestContent `bson:"content,omitempty" json:"content,omitempty"`
	Related           []DocumentManifestRelated `bson:"related,omitempty" json:"related,omitempty"`
}
type DocumentManifestContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PAttachment       *Attachment `bson:"pAttachment,omitempty" json:"pAttachment,omitempty"`
	PReference        *Reference  `bson:"pReference,omitempty" json:"pReference,omitempty"`
}
type DocumentManifestRelated struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref               *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

// OtherDocumentManifest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDocumentManifest DocumentManifest

// MarshalJSON marshals the given DocumentManifest as JSON into a byte slice
func (r DocumentManifest) MarshalJSON() ([]byte, error) {
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
		OtherDocumentManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentManifest: OtherDocumentManifest(r),
		ResourceType:          "DocumentManifest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DocumentManifest
func (r *DocumentManifest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDocumentManifest)(r)); err != nil {
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
func (r DocumentManifest) GetResourceType() ResourceType {
	return ResourceTypeDocumentManifest
}
