package fhir

import "encoding/json"

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
type DocumentManifest struct {
	Id                *string                   `bson:"id" json:"id"`
	Meta              *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules     *string                   `bson:"implicitRules" json:"implicitRules"`
	Language          *string                   `bson:"language" json:"language"`
	Text              *Narrative                `bson:"text" json:"text"`
	RawContained      []json.RawMessage         `bson:"contained" json:"contained"`
	Contained         []IResource               `bson:"-" json:"-"`
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
