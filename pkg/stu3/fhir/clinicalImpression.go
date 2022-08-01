package fhir

import "encoding/json"

// ClinicalImpression is documented here http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                     *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules            *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                 *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                     *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained             []json.RawMessage                 `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained                []IResource                       `bson:"-,omitempty" json:"-,omitempty"`
	Extension                []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier               []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                   ClinicalImpressionStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Code                     *CodeableConcept                  `bson:"code,omitempty" json:"code,omitempty"`
	Description              *string                           `bson:"description,omitempty" json:"description,omitempty"`
	Subject                  Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                  *Reference                        `bson:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime        *string                           `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                           `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Date                     *string                           `bson:"date,omitempty" json:"date,omitempty"`
	Assessor                 *Reference                        `bson:"assessor,omitempty" json:"assessor,omitempty"`
	Previous                 *Reference                        `bson:"previous,omitempty" json:"previous,omitempty"`
	Problem                  []Reference                       `bson:"problem,omitempty" json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `bson:"investigation,omitempty" json:"investigation,omitempty"`
	Protocol                 []string                          `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                  *string                           `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `bson:"finding,omitempty" json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	Action                   []Reference                       `bson:"action,omitempty" json:"action,omitempty"`
	Note                     []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
}
type ClinicalImpressionInvestigation struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Item              []Reference     `bson:"item,omitempty" json:"item,omitempty"`
}
type ClinicalImpressionFinding struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Basis               *string          `bson:"basis,omitempty" json:"basis,omitempty"`
}

// OtherClinicalImpression is a helper type to use the default implementations of Marshall and Unmarshal
type OtherClinicalImpression ClinicalImpression

// MarshalJSON marshals the given ClinicalImpression as JSON into a byte slice
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
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
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ClinicalImpression
func (r *ClinicalImpression) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherClinicalImpression)(r)); err != nil {
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
func (r ClinicalImpression) GetResourceType() ResourceType {
	return ResourceTypeClinicalImpression
}
