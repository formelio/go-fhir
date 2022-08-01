package fhir

import "encoding/json"

// Sequence is documented here http://hl7.org/fhir/StructureDefinition/Sequence
type Sequence struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *string               `bson:"type,omitempty" json:"type,omitempty"`
	CoordinateSystem  int                   `bson:"coordinateSystem,omitempty" json:"coordinateSystem,omitempty"`
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
	WindowStart         int              `bson:"windowStart,omitempty" json:"windowStart,omitempty"`
	WindowEnd           int              `bson:"windowEnd,omitempty" json:"windowEnd,omitempty"`
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
	Type              string           `bson:"type,omitempty" json:"type,omitempty"`
	StandardSequence  *CodeableConcept `bson:"standardSequence,omitempty" json:"standardSequence,omitempty"`
	Start             *int             `bson:"start,omitempty" json:"start,omitempty"`
	End               *int             `bson:"end,omitempty" json:"end,omitempty"`
	Score             *Quantity        `bson:"score,omitempty" json:"score,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	TruthTP           *float64         `bson:"truthTP,omitempty" json:"truthTP,omitempty"`
	QueryTP           *float64         `bson:"queryTP,omitempty" json:"queryTP,omitempty"`
	TruthFN           *float64         `bson:"truthFN,omitempty" json:"truthFN,omitempty"`
	QueryFP           *float64         `bson:"queryFP,omitempty" json:"queryFP,omitempty"`
	GtFP              *float64         `bson:"gtFP,omitempty" json:"gtFP,omitempty"`
	Precision         *float64         `bson:"precision,omitempty" json:"precision,omitempty"`
	Recall            *float64         `bson:"recall,omitempty" json:"recall,omitempty"`
	FScore            *float64         `bson:"fScore,omitempty" json:"fScore,omitempty"`
}
type SequenceRepository struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type,omitempty" json:"type,omitempty"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	DatasetId         *string     `bson:"datasetId,omitempty" json:"datasetId,omitempty"`
	VariantsetId      *string     `bson:"variantsetId,omitempty" json:"variantsetId,omitempty"`
	ReadsetId         *string     `bson:"readsetId,omitempty" json:"readsetId,omitempty"`
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
