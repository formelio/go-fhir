package fhir

import "encoding/json"

// Task is documented here http://hl7.org/fhir/StructureDefinition/Task
type Task struct {
	Id                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DefinitionUri       *string           `bson:"definitionUri,omitempty" json:"definitionUri,omitempty"`
	DefinitionReference *Reference        `bson:"definitionReference,omitempty" json:"definitionReference,omitempty"`
	BasedOn             []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier     *Identifier       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	PartOf              []Reference       `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status              TaskStatus        `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason        *CodeableConcept  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	BusinessStatus      *CodeableConcept  `bson:"businessStatus,omitempty" json:"businessStatus,omitempty"`
	Intent              RequestIntent     `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority            *RequestPriority  `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Description         *string           `bson:"description,omitempty" json:"description,omitempty"`
	Focus               *Reference        `bson:"focus,omitempty" json:"focus,omitempty"`
	For                 *Reference        `bson:"for,omitempty" json:"for,omitempty"`
	Context             *Reference        `bson:"context,omitempty" json:"context,omitempty"`
	ExecutionPeriod     *Period           `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty"`
	AuthoredOn          *string           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	LastModified        *string           `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Requester           *TaskRequester    `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType       []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Owner               *Reference        `bson:"owner,omitempty" json:"owner,omitempty"`
	Reason              *CodeableConcept  `bson:"reason,omitempty" json:"reason,omitempty"`
	Note                []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory     []Reference       `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Restriction         *TaskRestriction  `bson:"restriction,omitempty" json:"restriction,omitempty"`
	Input               []TaskInput       `bson:"input,omitempty" json:"input,omitempty"`
	Output              []TaskOutput      `bson:"output,omitempty" json:"output,omitempty"`
}
type TaskRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type TaskRestriction struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Repetitions       *int        `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Recipient         []Reference `bson:"recipient,omitempty" json:"recipient,omitempty"`
}
type TaskInput struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	ValueBase64Binary    *string          `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            *string          `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *string          `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueId              *string          `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueInstant         *string          `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        *string          `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueOid             *string          `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt     *int             `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueString          *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *string          `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt     *int             `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             *string          `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAge             *Age             `bson:"valueAge,omitempty" json:"valueAge,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueCount           *Count           `bson:"valueCount,omitempty" json:"valueCount,omitempty"`
	ValueDistance        *Distance        `bson:"valueDistance,omitempty" json:"valueDistance,omitempty"`
	ValueDuration        *Duration        `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney           *Money           `bson:"valueMoney,omitempty" json:"valueMoney,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
}
type TaskOutput struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	ValueBase64Binary    *string          `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            *string          `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *string          `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueId              *string          `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueInstant         *string          `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int             `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        *string          `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueOid             *string          `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt     *int             `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueString          *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *string          `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt     *int             `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             *string          `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAge             *Age             `bson:"valueAge,omitempty" json:"valueAge,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueCount           *Count           `bson:"valueCount,omitempty" json:"valueCount,omitempty"`
	ValueDistance        *Distance        `bson:"valueDistance,omitempty" json:"valueDistance,omitempty"`
	ValueDuration        *Duration        `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney           *Money           `bson:"valueMoney,omitempty" json:"valueMoney,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
}

// OtherTask is a helper type to use the default implementations of Marshall and Unmarshal
type OtherTask Task

// MarshalJSON marshals the given Task as JSON into a byte slice
func (r Task) MarshalJSON() ([]byte, error) {
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
		OtherTask
		ResourceType string `json:"resourceType"`
	}{
		OtherTask:    OtherTask(r),
		ResourceType: "Task",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Task
func (r *Task) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherTask)(r)); err != nil {
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
func (r Task) GetResourceType() ResourceType {
	return ResourceTypeTask
}
