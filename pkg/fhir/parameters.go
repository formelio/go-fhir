package fhir

import "encoding/json"

// Parameters is documented here http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
	Id            *string               `bson:"id" json:"id"`
	Meta          *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules *string               `bson:"implicitRules" json:"implicitRules"`
	Language      *string               `bson:"language" json:"language"`
	Parameter     []ParametersParameter `bson:"parameter" json:"parameter"`
}
type ParametersParameter struct {
	Id                   *string               `bson:"id" json:"id"`
	Extension            []Extension           `bson:"extension" json:"extension"`
	ModifierExtension    []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Name                 string                `bson:"name,omitempty" json:"name,omitempty"`
	ValueBase64Binary    *string               `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool                 `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            *string               `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDate            *string               `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *string               `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64              `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueId              *string               `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueInstant         *string               `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int                  `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        *string               `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueOid             *string               `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt     *int                  `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueString          *string               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *string               `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt     *int                  `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             *string               `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAddress         *Address              `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAge             *Age                  `bson:"valueAge,omitempty" json:"valueAge,omitempty"`
	ValueAnnotation      *Annotation           `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment           `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept *CodeableConcept      `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding               `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint         `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueCount           *Count                `bson:"valueCount,omitempty" json:"valueCount,omitempty"`
	ValueDistance        *Distance             `bson:"valueDistance,omitempty" json:"valueDistance,omitempty"`
	ValueDuration        *Duration             `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
	ValueHumanName       *HumanName            `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier      *Identifier           `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney           *Money                `bson:"valueMoney,omitempty" json:"valueMoney,omitempty"`
	ValuePeriod          *Period               `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity        *Quantity             `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range                `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference            `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData          `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature            `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueTiming          *Timing               `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueMeta            *Meta                 `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	RawResource          json.RawMessage       `bson:"resource" json:"resource"`
	Resource             IResource             `bson:"-" json:"-"`
	Part                 []ParametersParameter `bson:"part" json:"part"`
}

// OtherParametersParameter is a helper type to use the default implementations of Marshall and Unmarshal
type OtherParametersParameter ParametersParameter

func (r *ParametersParameter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherParametersParameter)(r)); err != nil {
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
func (r *ParametersParameter) MarshalJSON() ([]byte, error) {
	if r.Resource != nil {
		Resource, err := json.Marshal(r.Resource)
		if err != nil {
			return nil, err
		}
		r.RawResource = Resource
	}
	return json.Marshal((*OtherParametersParameter)(r))
}

// OtherParameters is a helper type to use the default implementations of Marshall and Unmarshal
type OtherParameters Parameters

// MarshalJSON marshals the given Parameters as JSON into a byte slice
func (r Parameters) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType"`
	}{
		OtherParameters: OtherParameters(r),
		ResourceType:    "Parameters",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Parameters
func (r *Parameters) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherParameters)(r)); err != nil {
		return err
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r Parameters) GetResourceType() ResourceType {
	return ResourceTypeParameters
}
