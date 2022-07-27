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
	RawResource       json.RawMessage      `bson:"resource" json:"resource"`
	Resource          IResource            `bson:"-" json:"-"`
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
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Status            string          `bson:"status,omitempty" json:"status,omitempty"`
	Location          *string         `bson:"location" json:"location"`
	Etag              *string         `bson:"etag" json:"etag"`
	LastModified      *string         `bson:"lastModified" json:"lastModified"`
	RawOutcome        json.RawMessage `bson:"outcome" json:"outcome"`
	Outcome           IResource       `bson:"-" json:"-"`
}

// OtherBundleEntryResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBundleEntryResponse BundleEntryResponse

func (r *BundleEntryResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBundleEntryResponse)(r)); err != nil {
		return err
	}
	var err error
	r.Outcome, err = UnmarshalResource(r.RawOutcome)
	if err != nil {
		return err
	}
	r.RawOutcome = nil
	return nil
}
func (r *BundleEntryResponse) MarshalJSON() ([]byte, error) {
	if r.Outcome != nil {
		Outcome, err := json.Marshal(r.Outcome)
		if err != nil {
			return nil, err
		}
		r.RawOutcome = Outcome
	}
	return json.Marshal((*OtherBundleEntryResponse)(r))
}

// OtherBundleEntry is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBundleEntry BundleEntry

func (r *BundleEntry) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBundleEntry)(r)); err != nil {
		return err
	}
	var err error
	r.Resource, err = UnmarshalResource(r.RawResource)
	if err != nil {
		return err
	}
	r.RawResource = nil
	return nil
}
func (r *BundleEntry) MarshalJSON() ([]byte, error) {
	if r.Resource != nil {
		Resource, err := json.Marshal(r.Resource)
		if err != nil {
			return nil, err
		}
		r.RawResource = Resource
	}
	return json.Marshal((*OtherBundleEntry)(r))
}

// OtherBundle is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice
func (r Bundle) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(r),
		ResourceType: "Bundle",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Bundle
func (r *Bundle) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBundle)(r)); err != nil {
		return err
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r Bundle) GetResourceType() ResourceType {
	return ResourceTypeBundle
}
