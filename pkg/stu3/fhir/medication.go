package fhir

import "encoding/json"

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Status            *MedicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	IsBrand           *bool                  `bson:"isBrand,omitempty" json:"isBrand,omitempty"`
	IsOverTheCounter  *bool                  `bson:"isOverTheCounter,omitempty" json:"isOverTheCounter,omitempty"`
	Manufacturer      *Reference             `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Form              *CodeableConcept       `bson:"form,omitempty" json:"form,omitempty"`
	Ingredient        []MedicationIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Package           *MedicationPackage     `bson:"package,omitempty" json:"package,omitempty"`
	Image             []Attachment           `bson:"image,omitempty" json:"image,omitempty"`
}
type MedicationIngredient struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	IsActive            *bool            `bson:"isActive,omitempty" json:"isActive,omitempty"`
	Amount              *Ratio           `bson:"amount,omitempty" json:"amount,omitempty"`
}
type MedicationPackage struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Container         *CodeableConcept           `bson:"container,omitempty" json:"container,omitempty"`
	Content           []MedicationPackageContent `bson:"content,omitempty" json:"content,omitempty"`
	Batch             []MedicationPackageBatch   `bson:"batch,omitempty" json:"batch,omitempty"`
}
type MedicationPackageContent struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Amount              *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}
type MedicationPackageBatch struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LotNumber         *string     `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate    *string     `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}

// OtherMedication is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice
func (r Medication) MarshalJSON() ([]byte, error) {
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
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Medication
func (r *Medication) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedication)(r)); err != nil {
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
func (r Medication) GetResourceType() ResourceType {
	return ResourceTypeMedication
}
