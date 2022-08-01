package fhir

import "encoding/json"

// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMap struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                  `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                  `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Structure         []StructureMapStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	Import            []string                `bson:"import,omitempty" json:"import,omitempty"`
	Group             []StructureMapGroup     `bson:"group,omitempty" json:"group,omitempty"`
}
type StructureMapStructure struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                `bson:"url,omitempty" json:"url,omitempty"`
	Mode              StructureMapModelMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Alias             *string               `bson:"alias,omitempty" json:"alias,omitempty"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroup struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                    `bson:"name,omitempty" json:"name,omitempty"`
	Extends           *string                   `bson:"extends,omitempty" json:"extends,omitempty"`
	TypeMode          StructureMapGroupTypeMode `bson:"typeMode,omitempty" json:"typeMode,omitempty"`
	Documentation     *string                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Input             []StructureMapGroupInput  `bson:"input,omitempty" json:"input,omitempty"`
	Rule              []StructureMapGroupRule   `bson:"rule,omitempty" json:"rule,omitempty"`
}
type StructureMapGroupInput struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                `bson:"name,omitempty" json:"name,omitempty"`
	Type              *string               `bson:"type,omitempty" json:"type,omitempty"`
	Mode              StructureMapInputMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRule struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Source            []StructureMapGroupRuleSource    `bson:"source,omitempty" json:"source,omitempty"`
	Target            []StructureMapGroupRuleTarget    `bson:"target,omitempty" json:"target,omitempty"`
	Rule              []StructureMapGroupRule          `bson:"rule,omitempty" json:"rule,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Documentation     *string                          `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRuleSource struct {
	Id                          *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context                     string                      `bson:"context,omitempty" json:"context,omitempty"`
	Min                         *int                        `bson:"min,omitempty" json:"min,omitempty"`
	Max                         *string                     `bson:"max,omitempty" json:"max,omitempty"`
	Type                        *string                     `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValueBase64Binary    *string                     `bson:"defaultValueBase64Binary,omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean         *bool                       `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCode            *string                     `bson:"defaultValueCode,omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueDate            *string                     `bson:"defaultValueDate,omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime        *string                     `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal         *float64                    `bson:"defaultValueDecimal,omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueId              *string                     `bson:"defaultValueId,omitempty" json:"defaultValueId,omitempty"`
	DefaultValueInstant         *string                     `bson:"defaultValueInstant,omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger         *int                        `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown        *string                     `bson:"defaultValueMarkdown,omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid             *string                     `bson:"defaultValueOid,omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt     *int                        `bson:"defaultValuePositiveInt,omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString          *string                     `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime            *string                     `bson:"defaultValueTime,omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt     *int                        `bson:"defaultValueUnsignedInt,omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri             *string                     `bson:"defaultValueUri,omitempty" json:"defaultValueUri,omitempty"`
	DefaultValueAddress         *Address                    `bson:"defaultValueAddress,omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAge             *Age                        `bson:"defaultValueAge,omitempty" json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation      *Annotation                 `bson:"defaultValueAnnotation,omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment      *Attachment                 `bson:"defaultValueAttachment,omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept *CodeableConcept            `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding          *Coding                     `bson:"defaultValueCoding,omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint    *ContactPoint               `bson:"defaultValueContactPoint,omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount           *Count                      `bson:"defaultValueCount,omitempty" json:"defaultValueCount,omitempty"`
	DefaultValueDistance        *Distance                   `bson:"defaultValueDistance,omitempty" json:"defaultValueDistance,omitempty"`
	DefaultValueDuration        *Duration                   `bson:"defaultValueDuration,omitempty" json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName       *HumanName                  `bson:"defaultValueHumanName,omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier      *Identifier                 `bson:"defaultValueIdentifier,omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney           *Money                      `bson:"defaultValueMoney,omitempty" json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod          *Period                     `bson:"defaultValuePeriod,omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity        *Quantity                   `bson:"defaultValueQuantity,omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange           *Range                      `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio           *Ratio                      `bson:"defaultValueRatio,omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueReference       *Reference                  `bson:"defaultValueReference,omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData     *SampledData                `bson:"defaultValueSampledData,omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature       *Signature                  `bson:"defaultValueSignature,omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueTiming          *Timing                     `bson:"defaultValueTiming,omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueMeta            *Meta                       `bson:"defaultValueMeta,omitempty" json:"defaultValueMeta,omitempty"`
	Element                     *string                     `bson:"element,omitempty" json:"element,omitempty"`
	ListMode                    *StructureMapSourceListMode `bson:"listMode,omitempty" json:"listMode,omitempty"`
	Variable                    *string                     `bson:"variable,omitempty" json:"variable,omitempty"`
	Condition                   *string                     `bson:"condition,omitempty" json:"condition,omitempty"`
	Check                       *string                     `bson:"check,omitempty" json:"check,omitempty"`
}
type StructureMapGroupRuleTarget struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           *string                                `bson:"context,omitempty" json:"context,omitempty"`
	ContextType       *StructureMapContextType               `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Element           *string                                `bson:"element,omitempty" json:"element,omitempty"`
	Variable          *string                                `bson:"variable,omitempty" json:"variable,omitempty"`
	ListMode          []StructureMapTargetListMode           `bson:"listMode,omitempty" json:"listMode,omitempty"`
	ListRuleId        *string                                `bson:"listRuleId,omitempty" json:"listRuleId,omitempty"`
	Transform         *StructureMapTransform                 `bson:"transform,omitempty" json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueId           *string     `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *float64    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
}
type StructureMapGroupRuleDependent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Variable          []string    `bson:"variable,omitempty" json:"variable,omitempty"`
}

// OtherStructureMap is a helper type to use the default implementations of Marshall and Unmarshal
type OtherStructureMap StructureMap

// MarshalJSON marshals the given StructureMap as JSON into a byte slice
func (r StructureMap) MarshalJSON() ([]byte, error) {
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
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

// UnmarshalJSON unmarshals the given byte slice into StructureMap
func (r *StructureMap) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherStructureMap)(r)); err != nil {
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
func (r StructureMap) GetResourceType() ResourceType {
	return ResourceTypeStructureMap
}
