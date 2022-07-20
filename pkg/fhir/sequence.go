package fhir

import "encoding/json"

// Sequence is documented here http://hl7.org/fhir/StructureDefinition/Sequence
type Sequence struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
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
type OtherSequence Sequence

// MarshalJSON marshals the given Sequence as JSON into a byte slice
func (r Sequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSequence
		ResourceType string `json:"resourceType"`
	}{
		OtherSequence: OtherSequence(r),
		ResourceType:  "Sequence",
	})
}

// UnmarshalSequence unmarshalls a Sequence.
func UnmarshalSequence(b []byte) (Sequence, error) {
	var sequence Sequence
	if err := json.Unmarshal(b, &sequence); err != nil {
		return sequence, err
	}
	return sequence, nil
}
