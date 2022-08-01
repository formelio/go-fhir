package fhir

import (
	"bytes"
	"encoding/json"
)

// CapabilityStatement is documented here http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                  *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage                  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource                        `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []*Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                 *string                            `bson:"url,omitempty" json:"url,omitempty"`
	Version             *string                            `bson:"version,omitempty" json:"version,omitempty"`
	Name                *string                            `bson:"name,omitempty" json:"name,omitempty"`
	Title               *string                            `bson:"title,omitempty" json:"title,omitempty"`
	Status              PublicationStatus                  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental        *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                string                             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           *string                            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact             []*ContactDetail                   `bson:"contact,omitempty" json:"contact,omitempty"`
	Description         *string                            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext          []*UsageContext                    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []*CodeableConcept                 `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose             *string                            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright           *string                            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind                CapabilityStatementKind            `bson:"kind,omitempty" json:"kind,omitempty"`
	Instantiates        []*string                          `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	Software            *CapabilityStatementSoftware       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `bson:"implementation,omitempty" json:"implementation,omitempty"`
	FhirVersion         string                             `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	AcceptUnknown       UnknownContentCode                 `bson:"acceptUnknown,omitempty" json:"acceptUnknown,omitempty"`
	Format              []string                           `bson:"format,omitempty" json:"format,omitempty"`
	PatchFormat         []*string                          `bson:"patchFormat,omitempty" json:"patchFormat,omitempty"`
	ImplementationGuide []*string                          `bson:"implementationGuide,omitempty" json:"implementationGuide,omitempty"`
	Profile             []*Reference                       `bson:"profile,omitempty" json:"profile,omitempty"`
	Rest                []*CapabilityStatementRest         `bson:"rest,omitempty" json:"rest,omitempty"`
	Messaging           []*CapabilityStatementMessaging    `bson:"messaging,omitempty" json:"messaging,omitempty"`
	Document            []*CapabilityStatementDocument     `bson:"document,omitempty" json:"document,omitempty"`
}
type CapabilityStatementSoftware struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Version           *string      `bson:"version,omitempty" json:"version,omitempty"`
	ReleaseDate       *string      `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}
type CapabilityStatementImplementation struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string       `bson:"description,omitempty" json:"description,omitempty"`
	Url               *string      `bson:"url,omitempty" json:"url,omitempty"`
}
type CapabilityStatementRest struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              RestfulCapabilityMode                         `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string                                       `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Security          *CapabilityStatementRestSecurity              `bson:"security,omitempty" json:"security,omitempty"`
	Resource          []*CapabilityStatementRestResource            `bson:"resource,omitempty" json:"resource,omitempty"`
	Interaction       []*CapabilityStatementRestInteraction         `bson:"interaction,omitempty" json:"interaction,omitempty"`
	SearchParam       []*CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation         []*CapabilityStatementRestOperation           `bson:"operation,omitempty" json:"operation,omitempty"`
	Compartment       []*string                                     `bson:"compartment,omitempty" json:"compartment,omitempty"`
}
type CapabilityStatementRestSecurity struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Cors              *bool                                         `bson:"cors,omitempty" json:"cors,omitempty"`
	Service           []*CodeableConcept                            `bson:"service,omitempty" json:"service,omitempty"`
	Description       *string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Certificate       []*CapabilityStatementRestSecurityCertificate `bson:"certificate,omitempty" json:"certificate,omitempty"`
}
type CapabilityStatementRestSecurityCertificate struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *string      `bson:"type,omitempty" json:"type,omitempty"`
	Blob              *string      `bson:"blob,omitempty" json:"blob,omitempty"`
}
type CapabilityStatementRestResource struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType                                  `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *Reference                                    `bson:"profile,omitempty" json:"profile,omitempty"`
	Documentation     *string                                       `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction  `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Versioning        *ResourceVersionPolicy                        `bson:"versioning,omitempty" json:"versioning,omitempty"`
	ReadHistory       *bool                                         `bson:"readHistory,omitempty" json:"readHistory,omitempty"`
	UpdateCreate      *bool                                         `bson:"updateCreate,omitempty" json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                         `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty"`
	ConditionalRead   *ConditionalReadStatus                        `bson:"conditionalRead,omitempty" json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                         `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty"`
	ConditionalDelete *ConditionalDeleteStatus                      `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty"`
	ReferencePolicy   []*ReferenceHandlingPolicy                    `bson:"referencePolicy,omitempty" json:"referencePolicy,omitempty"`
	SearchInclude     []*string                                     `bson:"searchInclude,omitempty" json:"searchInclude,omitempty"`
	SearchRevInclude  []*string                                     `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty"`
	SearchParam       []*CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
}
type CapabilityStatementRestResourceInteraction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              TypeRestfulInteraction `bson:"code,omitempty" json:"code,omitempty"`
	Documentation     *string                `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string          `bson:"name,omitempty" json:"name,omitempty"`
	Definition        *string         `bson:"definition,omitempty" json:"definition,omitempty"`
	Type              SearchParamType `bson:"type,omitempty" json:"type,omitempty"`
	Documentation     *string         `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestInteraction struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              SystemRestfulInteraction `bson:"code,omitempty" json:"code,omitempty"`
	Documentation     *string                  `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementRestOperation struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Definition        Reference    `bson:"definition,omitempty" json:"definition,omitempty"`
}
type CapabilityStatementMessaging struct {
	Id                *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Endpoint          []*CapabilityStatementMessagingEndpoint         `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	ReliableCache     *int                                            `bson:"reliableCache,omitempty" json:"reliableCache,omitempty"`
	Documentation     *string                                         `bson:"documentation,omitempty" json:"documentation,omitempty"`
	SupportedMessage  []*CapabilityStatementMessagingSupportedMessage `bson:"supportedMessage,omitempty" json:"supportedMessage,omitempty"`
	Event             []*CapabilityStatementMessagingEvent            `bson:"event,omitempty" json:"event,omitempty"`
}
type CapabilityStatementMessagingEndpoint struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Protocol          Coding       `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Address           string       `bson:"address,omitempty" json:"address,omitempty"`
}
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              EventCapabilityMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Definition        Reference           `bson:"definition,omitempty" json:"definition,omitempty"`
}
type CapabilityStatementMessagingEvent struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              Coding                       `bson:"code,omitempty" json:"code,omitempty"`
	Category          *MessageSignificanceCategory `bson:"category,omitempty" json:"category,omitempty"`
	Mode              EventCapabilityMode          `bson:"mode,omitempty" json:"mode,omitempty"`
	Focus             ResourceType                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Request           Reference                    `bson:"request,omitempty" json:"request,omitempty"`
	Response          Reference                    `bson:"response,omitempty" json:"response,omitempty"`
	Documentation     *string                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type CapabilityStatementDocument struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              DocumentMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Profile           Reference    `bson:"profile,omitempty" json:"profile,omitempty"`
}

// OtherCapabilityStatement is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCapabilityStatement CapabilityStatement

// MarshalJSON marshals the given CapabilityStatement as JSON into a byte slice
func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherCapabilityStatement
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into CapabilityStatement
func (r *CapabilityStatement) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCapabilityStatement)(r)); err != nil {
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
func (r CapabilityStatement) GetResourceType() ResourceType {
	return ResourceTypeCapabilityStatement
}
