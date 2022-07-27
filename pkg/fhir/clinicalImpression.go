package fhir

import "encoding/json"

// ClinicalImpression is documented here http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                           `bson:"id" json:"id"`
	Meta                     *Meta                             `bson:"meta" json:"meta"`
	ImplicitRules            *string                           `bson:"implicitRules" json:"implicitRules"`
	Language                 *string                           `bson:"language" json:"language"`
	Text                     *Narrative                        `bson:"text" json:"text"`
	RawContained             []json.RawMessage                 `bson:"contained" json:"contained"`
	Contained                []IResource                       `bson:"-" json:"-"`
	Extension                []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension        []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier               []Identifier                      `bson:"identifier" json:"identifier"`
	Status                   ClinicalImpressionStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Code                     *CodeableConcept                  `bson:"code" json:"code"`
	Description              *string                           `bson:"description" json:"description"`
	Subject                  Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                  *Reference                        `bson:"context" json:"context"`
	EffectiveDateTime        *string                           `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                           `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Date                     *string                           `bson:"date" json:"date"`
	Assessor                 *Reference                        `bson:"assessor" json:"assessor"`
	Previous                 *Reference                        `bson:"previous" json:"previous"`
	Problem                  []Reference                       `bson:"problem" json:"problem"`
	Investigation            []ClinicalImpressionInvestigation `bson:"investigation" json:"investigation"`
	Protocol                 []string                          `bson:"protocol" json:"protocol"`
	Summary                  *string                           `bson:"summary" json:"summary"`
	Finding                  []ClinicalImpressionFinding       `bson:"finding" json:"finding"`
	PrognosisCodeableConcept []CodeableConcept                 `bson:"prognosisCodeableConcept" json:"prognosisCodeableConcept"`
	PrognosisReference       []Reference                       `bson:"prognosisReference" json:"prognosisReference"`
	Action                   []Reference                       `bson:"action" json:"action"`
	Note                     []Annotation                      `bson:"note" json:"note"`
}
type ClinicalImpressionInvestigation struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Item              []Reference     `bson:"item" json:"item"`
}
type ClinicalImpressionFinding struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Basis               *string          `bson:"basis" json:"basis"`
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
