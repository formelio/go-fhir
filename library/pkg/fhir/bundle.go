package fhir

import "encoding/json"

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	Id            *string       `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string       `bson:"language,omitempty" json:"language,omitempty"`
	Identifier    *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type          string        `bson:"type" json:"type"`
	Total         *int          `bson:"total,omitempty" json:"total,omitempty"`
	Link          []BundleLink  `bson:"link,omitempty" json:"link,omitempty"`
	Entry         []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
}
type BundleLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation          string      `bson:"relation" json:"relation"`
	Url               string      `bson:"url" json:"url"`
}
type BundleEntry struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl           *string              `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource          json.RawMessage      `bson:"resource,omitempty" json:"resource,omitempty"`
	Search            *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request           *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response          *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
type BundleEntrySearch struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              *string     `bson:"mode,omitempty" json:"mode,omitempty"`
	Score             *string     `bson:"score,omitempty" json:"score,omitempty"`
}
type BundleEntryRequest struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            string      `bson:"method" json:"method"`
	Url               string      `bson:"url" json:"url"`
	IfNoneMatch       *string     `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist       *string     `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}
type BundleEntryResponse struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string          `bson:"status" json:"status"`
	Location          *string         `bson:"location,omitempty" json:"location,omitempty"`
	Etag              *string         `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified      *string         `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Outcome           json.RawMessage `bson:"outcome,omitempty" json:"outcome,omitempty"`
}
type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice
func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(r),
		ResourceType: "Bundle",
	})
}

// UnmarshalBundle unmarshals a Bundle.
func UnmarshalBundle(b []byte) (Bundle, error) {
	var bundle Bundle
	if err := json.Unmarshal(b, &bundle); err != nil {
		return bundle, err
	}
	return bundle, nil
}
