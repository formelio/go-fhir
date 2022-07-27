package fhir

import "encoding/json"

// Sequence is documented here http://hl7.org/fhir/StructureDefinition/Sequence
type Sequence struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Type              *string               `bson:"type" json:"type"`
	CoordinateSystem  int                   `bson:"coordinateSystem,omitempty" json:"coordinateSystem,omitempty"`
	Patient           *Reference            `bson:"patient" json:"patient"`
	Specimen          *Reference            `bson:"specimen" json:"specimen"`
	Device            *Reference            `bson:"device" json:"device"`
	Performer         *Reference            `bson:"performer" json:"performer"`
	Quantity          *Quantity             `bson:"quantity" json:"quantity"`
	ReferenceSeq      *SequenceReferenceSeq `bson:"referenceSeq" json:"referenceSeq"`
	Variant           []SequenceVariant     `bson:"variant" json:"variant"`
	ObservedSeq       *string               `bson:"observedSeq" json:"observedSeq"`
	Quality           []SequenceQuality     `bson:"quality" json:"quality"`
	ReadCoverage      *int                  `bson:"readCoverage" json:"readCoverage"`
	Repository        []SequenceRepository  `bson:"repository" json:"repository"`
	Pointer           []Reference           `bson:"pointer" json:"pointer"`
}
type SequenceReferenceSeq struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Chromosome          *CodeableConcept `bson:"chromosome" json:"chromosome"`
	GenomeBuild         *string          `bson:"genomeBuild" json:"genomeBuild"`
	ReferenceSeqId      *CodeableConcept `bson:"referenceSeqId" json:"referenceSeqId"`
	ReferenceSeqPointer *Reference       `bson:"referenceSeqPointer" json:"referenceSeqPointer"`
	ReferenceSeqString  *string          `bson:"referenceSeqString" json:"referenceSeqString"`
	Strand              *int             `bson:"strand" json:"strand"`
	WindowStart         int              `bson:"windowStart,omitempty" json:"windowStart,omitempty"`
	WindowEnd           int              `bson:"windowEnd,omitempty" json:"windowEnd,omitempty"`
}
type SequenceVariant struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Start             *int        `bson:"start" json:"start"`
	End               *int        `bson:"end" json:"end"`
	ObservedAllele    *string     `bson:"observedAllele" json:"observedAllele"`
	ReferenceAllele   *string     `bson:"referenceAllele" json:"referenceAllele"`
	Cigar             *string     `bson:"cigar" json:"cigar"`
	VariantPointer    *Reference  `bson:"variantPointer" json:"variantPointer"`
}
type SequenceQuality struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              string           `bson:"type,omitempty" json:"type,omitempty"`
	StandardSequence  *CodeableConcept `bson:"standardSequence" json:"standardSequence"`
	Start             *int             `bson:"start" json:"start"`
	End               *int             `bson:"end" json:"end"`
	Score             *Quantity        `bson:"score" json:"score"`
	Method            *CodeableConcept `bson:"method" json:"method"`
	TruthTP           *float64         `bson:"truthTP" json:"truthTP"`
	QueryTP           *float64         `bson:"queryTP" json:"queryTP"`
	TruthFN           *float64         `bson:"truthFN" json:"truthFN"`
	QueryFP           *float64         `bson:"queryFP" json:"queryFP"`
	GtFP              *float64         `bson:"gtFP" json:"gtFP"`
	Precision         *float64         `bson:"precision" json:"precision"`
	Recall            *float64         `bson:"recall" json:"recall"`
	FScore            *float64         `bson:"fScore" json:"fScore"`
}
type SequenceRepository struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Type              string      `bson:"type,omitempty" json:"type,omitempty"`
	Url               *string     `bson:"url" json:"url"`
	Name              *string     `bson:"name" json:"name"`
	DatasetId         *string     `bson:"datasetId" json:"datasetId"`
	VariantsetId      *string     `bson:"variantsetId" json:"variantsetId"`
	ReadsetId         *string     `bson:"readsetId" json:"readsetId"`
}

// OtherSequence is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSequence Sequence

// MarshalJSON marshals the given Sequence as JSON into a byte slice
func (r Sequence) MarshalJSON() ([]byte, error) {
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
		OtherSequence
		ResourceType string `json:"resourceType"`
	}{
		OtherSequence: OtherSequence(r),
		ResourceType:  "Sequence",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Sequence
func (r *Sequence) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSequence)(r)); err != nil {
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
func (r Sequence) GetResourceType() ResourceType {
	return ResourceTypeSequence
}
