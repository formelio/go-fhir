package fhir

import "encoding/json"

// CapabilityStatement is documented here http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatement struct {
	Id                  *string                            `bson:"id" json:"id"`
	Meta                *Meta                              `bson:"meta" json:"meta"`
	ImplicitRules       *string                            `bson:"implicitRules" json:"implicitRules"`
	Language            *string                            `bson:"language" json:"language"`
	Text                *Narrative                         `bson:"text" json:"text"`
	RawContained        []json.RawMessage                  `bson:"contained" json:"contained"`
	Contained           []IResource                        `bson:"-" json:"-"`
	Extension           []Extension                        `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                        `bson:"modifierExtension" json:"modifierExtension"`
	Url                 *string                            `bson:"url" json:"url"`
	Version             *string                            `bson:"version" json:"version"`
	Name                *string                            `bson:"name" json:"name"`
	Title               *string                            `bson:"title" json:"title"`
	Status              PublicationStatus                  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental        *bool                              `bson:"experimental" json:"experimental"`
	Date                string                             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           *string                            `bson:"publisher" json:"publisher"`
	Contact             []ContactDetail                    `bson:"contact" json:"contact"`
	Description         *string                            `bson:"description" json:"description"`
	UseContext          []UsageContext                     `bson:"useContext" json:"useContext"`
	Jurisdiction        []CodeableConcept                  `bson:"jurisdiction" json:"jurisdiction"`
	Purpose             *string                            `bson:"purpose" json:"purpose"`
	Copyright           *string                            `bson:"copyright" json:"copyright"`
	Kind                CapabilityStatementKind            `bson:"kind,omitempty" json:"kind,omitempty"`
	Instantiates        []string                           `bson:"instantiates" json:"instantiates"`
	Software            *CapabilityStatementSoftware       `bson:"software" json:"software"`
	Implementation      *CapabilityStatementImplementation `bson:"implementation" json:"implementation"`
	FhirVersion         string                             `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	AcceptUnknown       UnknownContentCode                 `bson:"acceptUnknown,omitempty" json:"acceptUnknown,omitempty"`
	Format              []string                           `bson:"format,omitempty" json:"format,omitempty"`
	PatchFormat         []string                           `bson:"patchFormat" json:"patchFormat"`
	ImplementationGuide []string                           `bson:"implementationGuide" json:"implementationGuide"`
	Profile             []Reference                        `bson:"profile" json:"profile"`
	Rest                []CapabilityStatementRest          `bson:"rest" json:"rest"`
	Messaging           []CapabilityStatementMessaging     `bson:"messaging" json:"messaging"`
	Document            []CapabilityStatementDocument      `bson:"document" json:"document"`
}
type CapabilityStatementSoftware struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Version           *string     `bson:"version" json:"version"`
	ReleaseDate       *string     `bson:"releaseDate" json:"releaseDate"`
}
type CapabilityStatementImplementation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       string      `bson:"description,omitempty" json:"description,omitempty"`
	Url               *string     `bson:"url" json:"url"`
}
type CapabilityStatementRest struct {
	Id                *string                                      `bson:"id" json:"id"`
	Extension         []Extension                                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                  `bson:"modifierExtension" json:"modifierExtension"`
	Mode              RestfulCapabilityMode                        `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string                                      `bson:"documentation" json:"documentation"`
	Security          *CapabilityStatementRestSecurity             `bson:"security" json:"security"`
	Resource          []CapabilityStatementRestResource            `bson:"resource" json:"resource"`
	Interaction       []CapabilityStatementRestInteraction         `bson:"interaction" json:"interaction"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `bson:"searchParam" json:"searchParam"`
	Operation         []CapabilityStatementRestOperation           `bson:"operation" json:"operation"`
	Compartment       []string                                     `bson:"compartment" json:"compartment"`
}
type CapabilityStatementRestSecurity struct {
	Id                *string                                      `bson:"id" json:"id"`
	Extension         []Extension                                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                  `bson:"modifierExtension" json:"modifierExtension"`
	Cors              *bool                                        `bson:"cors" json:"cors"`
	Service           []CodeableConcept                            `bson:"service" json:"service"`
	Description       *string                                      `bson:"description" json:"description"`
	Certificate       []CapabilityStatementRestSecurityCertificate `bson:"certificate" json:"certificate"`
}
type CapabilityStatementRestSecurityCertificate struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Type              *string     `bson:"type" json:"type"`
	Blob              *string     `bson:"blob" json:"blob"`
}
type CapabilityStatementRestResource struct {
	Id                *string                                      `bson:"id" json:"id"`
	Extension         []Extension                                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                  `bson:"modifierExtension" json:"modifierExtension"`
	Type              ResourceType                                 `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *Reference                                   `bson:"profile" json:"profile"`
	Documentation     *string                                      `bson:"documentation" json:"documentation"`
	Interaction       []CapabilityStatementRestResourceInteraction `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Versioning        *ResourceVersionPolicy                       `bson:"versioning" json:"versioning"`
	ReadHistory       *bool                                        `bson:"readHistory" json:"readHistory"`
	UpdateCreate      *bool                                        `bson:"updateCreate" json:"updateCreate"`
	ConditionalCreate *bool                                        `bson:"conditionalCreate" json:"conditionalCreate"`
	ConditionalRead   *ConditionalReadStatus                       `bson:"conditionalRead" json:"conditionalRead"`
	ConditionalUpdate *bool                                        `bson:"conditionalUpdate" json:"conditionalUpdate"`
	ConditionalDelete *ConditionalDeleteStatus                     `bson:"conditionalDelete" json:"conditionalDelete"`
	ReferencePolicy   []ReferenceHandlingPolicy                    `bson:"referencePolicy" json:"referencePolicy"`
	SearchInclude     []string                                     `bson:"searchInclude" json:"searchInclude"`
	SearchRevInclude  []string                                     `bson:"searchRevInclude" json:"searchRevInclude"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `bson:"searchParam" json:"searchParam"`
}
type CapabilityStatementRestResourceInteraction struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Code              TypeRestfulInteraction `bson:"code,omitempty" json:"code,omitempty"`
	Documentation     *string                `bson:"documentation" json:"documentation"`
}
type CapabilityStatementRestResourceSearchParam struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Name              string          `bson:"name,omitempty" json:"name,omitempty"`
	Definition        *string         `bson:"definition" json:"definition"`
	Type              SearchParamType `bson:"type,omitempty" json:"type,omitempty"`
	Documentation     *string         `bson:"documentation" json:"documentation"`
}
type CapabilityStatementRestInteraction struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	Code              SystemRestfulInteraction `bson:"code,omitempty" json:"code,omitempty"`
	Documentation     *string                  `bson:"documentation" json:"documentation"`
}
type CapabilityStatementRestOperation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Definition        Reference   `bson:"definition,omitempty" json:"definition,omitempty"`
}
type CapabilityStatementMessaging struct {
	Id                *string                                        `bson:"id" json:"id"`
	Extension         []Extension                                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                    `bson:"modifierExtension" json:"modifierExtension"`
	Endpoint          []CapabilityStatementMessagingEndpoint         `bson:"endpoint" json:"endpoint"`
	ReliableCache     *int                                           `bson:"reliableCache" json:"reliableCache"`
	Documentation     *string                                        `bson:"documentation" json:"documentation"`
	SupportedMessage  []CapabilityStatementMessagingSupportedMessage `bson:"supportedMessage" json:"supportedMessage"`
	Event             []CapabilityStatementMessagingEvent            `bson:"event" json:"event"`
}
type CapabilityStatementMessagingEndpoint struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Protocol          Coding      `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Address           string      `bson:"address,omitempty" json:"address,omitempty"`
}
type CapabilityStatementMessagingSupportedMessage struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Mode              EventCapabilityMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Definition        Reference           `bson:"definition,omitempty" json:"definition,omitempty"`
}
type CapabilityStatementMessagingEvent struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Code              Coding                       `bson:"code,omitempty" json:"code,omitempty"`
	Category          *MessageSignificanceCategory `bson:"category" json:"category"`
	Mode              EventCapabilityMode          `bson:"mode,omitempty" json:"mode,omitempty"`
	Focus             ResourceType                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Request           Reference                    `bson:"request,omitempty" json:"request,omitempty"`
	Response          Reference                    `bson:"response,omitempty" json:"response,omitempty"`
	Documentation     *string                      `bson:"documentation" json:"documentation"`
}
type CapabilityStatementDocument struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Mode              DocumentMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string      `bson:"documentation" json:"documentation"`
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
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherCapabilityStatement: OtherCapabilityStatement(r),
		ResourceType:             "CapabilityStatement",
	})
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
