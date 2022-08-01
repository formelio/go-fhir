package fhir

import "encoding/json"

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	Id            *string       `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string       `bson:"language,omitempty" json:"language,omitempty"`
	Identifier    *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type          BundleType    `bson:"type,omitempty" json:"type,omitempty"`
	Total         *int          `bson:"total,omitempty" json:"total,omitempty"`
	Link          []BundleLink  `bson:"link,omitempty" json:"link,omitempty"`
	Entry         []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
}
type BundleLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation          string      `bson:"relation,omitempty" json:"relation,omitempty"`
	Url               string      `bson:"url,omitempty" json:"url,omitempty"`
}
type BundleEntry struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl           *string              `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	RawResource       json.RawMessage      `bson:"resource,omitempty" json:"resource,omitempty"`
	Resource          IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Search            *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request           *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response          *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
type BundleEntrySearch struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              *SearchEntryMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Score             *float64         `bson:"score,omitempty" json:"score,omitempty"`
}
type BundleEntryRequest struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            HTTPVerb    `bson:"method,omitempty" json:"method,omitempty"`
	Url               string      `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch       *string     `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist       *string     `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}
type BundleEntryResponse struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string          `bson:"status,omitempty" json:"status,omitempty"`
	Location          *string         `bson:"location,omitempty" json:"location,omitempty"`
	Etag              *string         `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified      *string         `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	RawOutcome        json.RawMessage `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Outcome           IResource       `bson:"-,omitempty" json:"-,omitempty"`
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
