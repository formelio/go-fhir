package fhir

import "encoding/json"

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	Id            *string       `bson:"id" json:"id"`
	Meta          *Meta         `bson:"meta" json:"meta"`
	ImplicitRules *string       `bson:"implicitRules" json:"implicitRules"`
	Language      *string       `bson:"language" json:"language"`
	Identifier    *Identifier   `bson:"identifier" json:"identifier"`
	Type          BundleType    `bson:"type,omitempty" json:"type,omitempty"`
	Total         *int          `bson:"total" json:"total"`
	Link          []BundleLink  `bson:"link" json:"link"`
	Entry         []BundleEntry `bson:"entry" json:"entry"`
	Signature     *Signature    `bson:"signature" json:"signature"`
}
type BundleLink struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Relation          string      `bson:"relation,omitempty" json:"relation,omitempty"`
	Url               string      `bson:"url,omitempty" json:"url,omitempty"`
}
type BundleEntry struct {
	Id                *string              `bson:"id" json:"id"`
	Extension         []Extension          `bson:"extension" json:"extension"`
	ModifierExtension []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Link              []BundleLink         `bson:"link" json:"link"`
	FullUrl           *string              `bson:"fullUrl" json:"fullUrl"`
	Resource          *json.RawMessage     `bson:"resource" json:"resource"`
	Search            *BundleEntrySearch   `bson:"search" json:"search"`
	Request           *BundleEntryRequest  `bson:"request" json:"request"`
	Response          *BundleEntryResponse `bson:"response" json:"response"`
}
type BundleEntrySearch struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Mode              *SearchEntryMode `bson:"mode" json:"mode"`
	Score             *float64         `bson:"score" json:"score"`
}
type BundleEntryRequest struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Method            HTTPVerb    `bson:"method,omitempty" json:"method,omitempty"`
	Url               string      `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch       *string     `bson:"ifNoneMatch" json:"ifNoneMatch"`
	IfModifiedSince   *string     `bson:"ifModifiedSince" json:"ifModifiedSince"`
	IfMatch           *string     `bson:"ifMatch" json:"ifMatch"`
	IfNoneExist       *string     `bson:"ifNoneExist" json:"ifNoneExist"`
}
type BundleEntryResponse struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Status            string           `bson:"status,omitempty" json:"status,omitempty"`
	Location          *string          `bson:"location" json:"location"`
	Etag              *string          `bson:"etag" json:"etag"`
	LastModified      *string          `bson:"lastModified" json:"lastModified"`
	Outcome           *json.RawMessage `bson:"outcome" json:"outcome"`
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

// UnmarshalBundle unmarshalls a Bundle.
func UnmarshalBundle(b []byte) (Bundle, error) {
	var bundle Bundle
	if err := json.Unmarshal(b, &bundle); err != nil {
		return bundle, err
	}
	return bundle, nil
}
