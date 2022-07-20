package fhir

import "encoding/json"

// Sequence is documented here http://hl7.org/fhir/StructureDefinition/Sequence
type Sequence struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *string               `bson:"type,omitempty" json:"type,omitempty"`
	CoordinateSystem  int                   `bson:"coordinateSystem" json:"coordinateSystem"`
	Patient           *Reference            `bson:"patient,omitempty" json:"patient,omitempty"`
	Specimen          *Reference            `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device            *Reference            `bson:"device,omitempty" json:"device,omitempty"`
	Performer         *Reference            `bson:"performer,omitempty" json:"performer,omitempty"`
	Quantity          *Quantity             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ReferenceSeq      *SequenceReferenceSeq `bson:"referenceSeq,omitempty" json:"referenceSeq,omitempty"`
	Variant           []SequenceVariant     `bson:"variant,omitempty" json:"variant,omitempty"`
	ObservedSeq       *string               `bson:"observedSeq,omitempty" json:"observedSeq,omitempty"`
	Quality           []SequenceQuality     `bson:"quality,omitempty" json:"quality,omitempty"`
	ReadCoverage      *int                  `bson:"readCoverage,omitempty" json:"readCoverage,omitempty"`
	Repository        []SequenceRepository  `bson:"repository,omitempty" json:"repository,omitempty"`
	Pointer           []Reference           `bson:"pointer,omitempty" json:"pointer,omitempty"`
}
type SequenceReferenceSeq struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Chromosome          *CodeableConcept `bson:"chromosome,omitempty" json:"chromosome,omitempty"`
	GenomeBuild         *string          `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty"`
	ReferenceSeqId      *CodeableConcept `bson:"referenceSeqId,omitempty" json:"referenceSeqId,omitempty"`
	ReferenceSeqPointer *Reference       `bson:"referenceSeqPointer,omitempty" json:"referenceSeqPointer,omitempty"`
	ReferenceSeqString  *string          `bson:"referenceSeqString,omitempty" json:"referenceSeqString,omitempty"`
	Strand              *int             `bson:"strand,omitempty" json:"strand,omitempty"`
	WindowStart         int              `bson:"windowStart" json:"windowStart"`
	WindowEnd           int              `bson:"windowEnd" json:"windowEnd"`
}
type SequenceVariant struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Start             *int        `bson:"start,omitempty" json:"start,omitempty"`
	End               *int        `bson:"end,omitempty" json:"end,omitempty"`
	ObservedAllele    *string     `bson:"observedAllele,omitempty" json:"observedAllele,omitempty"`
	ReferenceAllele   *string     `bson:"referenceAllele,omitempty" json:"referenceAllele,omitempty"`
	Cigar             *string     `bson:"cigar,omitempty" json:"cigar,omitempty"`
	VariantPointer    *Reference  `bson:"variantPointer,omitempty" json:"variantPointer,omitempty"`
}
type SequenceQuality struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string           `bson:"type" json:"type"`
	StandardSequence  *CodeableConcept `bson:"standardSequence,omitempty" json:"standardSequence,omitempty"`
	Start             *int             `bson:"start,omitempty" json:"start,omitempty"`
	End               *int             `bson:"end,omitempty" json:"end,omitempty"`
	Score             *Quantity        `bson:"score,omitempty" json:"score,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	TruthTP           *string          `bson:"truthTP,omitempty" json:"truthTP,omitempty"`
	QueryTP           *string          `bson:"queryTP,omitempty" json:"queryTP,omitempty"`
	TruthFN           *string          `bson:"truthFN,omitempty" json:"truthFN,omitempty"`
	QueryFP           *string          `bson:"queryFP,omitempty" json:"queryFP,omitempty"`
	GtFP              *string          `bson:"gtFP,omitempty" json:"gtFP,omitempty"`
	Precision         *string          `bson:"precision,omitempty" json:"precision,omitempty"`
	Recall            *string          `bson:"recall,omitempty" json:"recall,omitempty"`
	FScore            *string          `bson:"fScore,omitempty" json:"fScore,omitempty"`
}
type SequenceRepository struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	DatasetId         *string     `bson:"datasetId,omitempty" json:"datasetId,omitempty"`
	VariantsetId      *string     `bson:"variantsetId,omitempty" json:"variantsetId,omitempty"`
	ReadsetId         *string     `bson:"readsetId,omitempty" json:"readsetId,omitempty"`
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

// UnmarshalSequence unmarshals a Sequence.
func UnmarshalSequence(b []byte) (Sequence, error) {
	var sequence Sequence
	if err := json.Unmarshal(b, &sequence); err != nil {
		return sequence, err
	}
	return sequence, nil
}
