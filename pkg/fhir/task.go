package fhir

import "encoding/json"

// Task is documented here http://hl7.org/fhir/StructureDefinition/Task
type Task struct {
	Id                  *string           `bson:"id" json:"id"`
	Meta                *Meta             `bson:"meta" json:"meta"`
	ImplicitRules       *string           `bson:"implicitRules" json:"implicitRules"`
	Language            *string           `bson:"language" json:"language"`
	Text                *Narrative        `bson:"text" json:"text"`
	Contained           []json.RawMessage `bson:"contained" json:"contained"`
	Extension           []Extension       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier      `bson:"identifier" json:"identifier"`
	DefinitionUri       *string           `bson:"definitionUri,omitempty" json:"definitionUri,omitempty"`
	DefinitionReference *Reference        `bson:"definitionReference,omitempty" json:"definitionReference,omitempty"`
	BasedOn             []Reference       `bson:"basedOn" json:"basedOn"`
	GroupIdentifier     *Identifier       `bson:"groupIdentifier" json:"groupIdentifier"`
	PartOf              []Reference       `bson:"partOf" json:"partOf"`
	Status              TaskStatus        `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason        *CodeableConcept  `bson:"statusReason" json:"statusReason"`
	BusinessStatus      *CodeableConcept  `bson:"businessStatus" json:"businessStatus"`
	Intent              RequestIntent     `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority            *RequestPriority  `bson:"priority" json:"priority"`
	Code                *CodeableConcept  `bson:"code" json:"code"`
	Description         *string           `bson:"description" json:"description"`
	Focus               *Reference        `bson:"focus" json:"focus"`
	For                 *Reference        `bson:"for" json:"for"`
	Context             *Reference        `bson:"context" json:"context"`
	ExecutionPeriod     *Period           `bson:"executionPeriod" json:"executionPeriod"`
	AuthoredOn          *string           `bson:"authoredOn" json:"authoredOn"`
	LastModified        *string           `bson:"lastModified" json:"lastModified"`
	Requester           *TaskRequester    `bson:"requester" json:"requester"`
	PerformerType       []CodeableConcept `bson:"performerType" json:"performerType"`
	Owner               *Reference        `bson:"owner" json:"owner"`
	Reason              *CodeableConcept  `bson:"reason" json:"reason"`
	Note                []Annotation      `bson:"note" json:"note"`
	RelevantHistory     []Reference       `bson:"relevantHistory" json:"relevantHistory"`
	Restriction         *TaskRestriction  `bson:"restriction" json:"restriction"`
	Input               []TaskInput       `bson:"input" json:"input"`
	Output              []TaskOutput      `bson:"output" json:"output"`
}
type TaskRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type TaskRestriction struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Repetitions       *int        `bson:"repetitions" json:"repetitions"`
	Period            *Period     `bson:"period" json:"period"`
	Recipient         []Reference `bson:"recipient" json:"recipient"`
}
type TaskInput struct {
	Id                   *string          `bson:"id" json:"id"`
	Extension            []Extension      `bson:"extension" json:"extension"`
	ModifierExtension    []Extension      `bson:"modifierExtension" json:"modifierExtension"`
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
	Id                   *string          `bson:"id" json:"id"`
	Extension            []Extension      `bson:"extension" json:"extension"`
	ModifierExtension    []Extension      `bson:"modifierExtension" json:"modifierExtension"`
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
type OtherTask Task

// MarshalJSON marshals the given Task as JSON into a byte slice
func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTask
		ResourceType string `json:"resourceType"`
	}{
		OtherTask:    OtherTask(r),
		ResourceType: "Task",
	})
}

// UnmarshalTask unmarshalls a Task.
func UnmarshalTask(b []byte) (Task, error) {
	var task Task
	if err := json.Unmarshal(b, &task); err != nil {
		return task, err
	}
	return task, nil
}
