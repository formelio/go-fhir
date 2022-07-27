package fhir

import "encoding/json"

// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMap struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	RawContained      []json.RawMessage       `bson:"contained" json:"contained"`
	Contained         []IResource             `bson:"-" json:"-"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                  `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier            `bson:"identifier" json:"identifier"`
	Version           *string                 `bson:"version" json:"version"`
	Name              string                  `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                 `bson:"title" json:"title"`
	Status            PublicationStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                   `bson:"experimental" json:"experimental"`
	Date              *string                 `bson:"date" json:"date"`
	Publisher         *string                 `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail         `bson:"contact" json:"contact"`
	Description       *string                 `bson:"description" json:"description"`
	UseContext        []UsageContext          `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                 `bson:"purpose" json:"purpose"`
	Copyright         *string                 `bson:"copyright" json:"copyright"`
	Structure         []StructureMapStructure `bson:"structure" json:"structure"`
	Import            []string                `bson:"import" json:"import"`
	Group             []StructureMapGroup     `bson:"group,omitempty" json:"group,omitempty"`
}
type StructureMapStructure struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                `bson:"url,omitempty" json:"url,omitempty"`
	Mode              StructureMapModelMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Alias             *string               `bson:"alias" json:"alias"`
	Documentation     *string               `bson:"documentation" json:"documentation"`
}
type StructureMapGroup struct {
	Id                *string                   `bson:"id" json:"id"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                    `bson:"name,omitempty" json:"name,omitempty"`
	Extends           *string                   `bson:"extends" json:"extends"`
	TypeMode          StructureMapGroupTypeMode `bson:"typeMode,omitempty" json:"typeMode,omitempty"`
	Documentation     *string                   `bson:"documentation" json:"documentation"`
	Input             []StructureMapGroupInput  `bson:"input,omitempty" json:"input,omitempty"`
	Rule              []StructureMapGroupRule   `bson:"rule,omitempty" json:"rule,omitempty"`
}
type StructureMapGroupInput struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                `bson:"name,omitempty" json:"name,omitempty"`
	Type              *string               `bson:"type" json:"type"`
	Mode              StructureMapInputMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation     *string               `bson:"documentation" json:"documentation"`
}
type StructureMapGroupRule struct {
	Id                *string                          `bson:"id" json:"id"`
	Extension         []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Source            []StructureMapGroupRuleSource    `bson:"source,omitempty" json:"source,omitempty"`
	Target            []StructureMapGroupRuleTarget    `bson:"target" json:"target"`
	Rule              []StructureMapGroupRule          `bson:"rule" json:"rule"`
	Dependent         []StructureMapGroupRuleDependent `bson:"dependent" json:"dependent"`
	Documentation     *string                          `bson:"documentation" json:"documentation"`
}
type StructureMapGroupRuleSource struct {
	Id                          *string                     `bson:"id" json:"id"`
	Extension                   []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension           []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Context                     string                      `bson:"context,omitempty" json:"context,omitempty"`
	Min                         *int                        `bson:"min" json:"min"`
	Max                         *string                     `bson:"max" json:"max"`
	Type                        *string                     `bson:"type" json:"type"`
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
	Element                     *string                     `bson:"element" json:"element"`
	ListMode                    *StructureMapSourceListMode `bson:"listMode" json:"listMode"`
	Variable                    *string                     `bson:"variable" json:"variable"`
	Condition                   *string                     `bson:"condition" json:"condition"`
	Check                       *string                     `bson:"check" json:"check"`
}
type StructureMapGroupRuleTarget struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	Context           *string                                `bson:"context" json:"context"`
	ContextType       *StructureMapContextType               `bson:"contextType" json:"contextType"`
	Element           *string                                `bson:"element" json:"element"`
	Variable          *string                                `bson:"variable" json:"variable"`
	ListMode          []StructureMapTargetListMode           `bson:"listMode" json:"listMode"`
	ListRuleId        *string                                `bson:"listRuleId" json:"listRuleId"`
	Transform         *StructureMapTransform                 `bson:"transform" json:"transform"`
	Parameter         []StructureMapGroupRuleTargetParameter `bson:"parameter" json:"parameter"`
}
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ValueId           *string     `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *float64    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
}
type StructureMapGroupRuleDependent struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
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
