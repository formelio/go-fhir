package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ResourceType is documented here http://hl7.org/fhir/ValueSet/resource-types
type ResourceType int

const (
	ResourceTypeAccount ResourceType = iota
	ResourceTypeActivityDefinition
	ResourceTypeAdverseEvent
	ResourceTypeAllergyIntolerance
	ResourceTypeAppointment
	ResourceTypeAppointmentResponse
	ResourceTypeAuditEvent
	ResourceTypeBasic
	ResourceTypeBinary
	ResourceTypeBodySite
	ResourceTypeBundle
	ResourceTypeCapabilityStatement
	ResourceTypeCarePlan
	ResourceTypeCareTeam
	ResourceTypeChargeItem
	ResourceTypeClaim
	ResourceTypeClaimResponse
	ResourceTypeClinicalImpression
	ResourceTypeCodeSystem
	ResourceTypeCommunication
	ResourceTypeCommunicationRequest
	ResourceTypeCompartmentDefinition
	ResourceTypeComposition
	ResourceTypeConceptMap
	ResourceTypeCondition
	ResourceTypeConsent
	ResourceTypeContract
	ResourceTypeCoverage
	ResourceTypeDataElement
	ResourceTypeDetectedIssue
	ResourceTypeDevice
	ResourceTypeDeviceComponent
	ResourceTypeDeviceMetric
	ResourceTypeDeviceRequest
	ResourceTypeDeviceUseStatement
	ResourceTypeDiagnosticReport
	ResourceTypeDocumentManifest
	ResourceTypeDocumentReference
	ResourceTypeDomainResource
	ResourceTypeEligibilityRequest
	ResourceTypeEligibilityResponse
	ResourceTypeEncounter
	ResourceTypeEndpoint
	ResourceTypeEnrollmentRequest
	ResourceTypeEnrollmentResponse
	ResourceTypeEpisodeOfCare
	ResourceTypeExpansionProfile
	ResourceTypeExplanationOfBenefit
	ResourceTypeFamilyMemberHistory
	ResourceTypeFlag
	ResourceTypeGoal
	ResourceTypeGraphDefinition
	ResourceTypeGroup
	ResourceTypeGuidanceResponse
	ResourceTypeHealthcareService
	ResourceTypeImagingManifest
	ResourceTypeImagingStudy
	ResourceTypeImmunization
	ResourceTypeImmunizationRecommendation
	ResourceTypeImplementationGuide
	ResourceTypeLibrary
	ResourceTypeLinkage
	ResourceTypeList
	ResourceTypeLocation
	ResourceTypeMeasure
	ResourceTypeMeasureReport
	ResourceTypeMedia
	ResourceTypeMedication
	ResourceTypeMedicationAdministration
	ResourceTypeMedicationDispense
	ResourceTypeMedicationRequest
	ResourceTypeMedicationStatement
	ResourceTypeMessageDefinition
	ResourceTypeMessageHeader
	ResourceTypeNamingSystem
	ResourceTypeNutritionOrder
	ResourceTypeObservation
	ResourceTypeOperationDefinition
	ResourceTypeOperationOutcome
	ResourceTypeOrganization
	ResourceTypeParameters
	ResourceTypePatient
	ResourceTypePaymentNotice
	ResourceTypePaymentReconciliation
	ResourceTypePerson
	ResourceTypePlanDefinition
	ResourceTypePractitioner
	ResourceTypePractitionerRole
	ResourceTypeProcedure
	ResourceTypeProcedureRequest
	ResourceTypeProcessRequest
	ResourceTypeProcessResponse
	ResourceTypeProvenance
	ResourceTypeQuestionnaire
	ResourceTypeQuestionnaireResponse
	ResourceTypeReferralRequest
	ResourceTypeRelatedPerson
	ResourceTypeRequestGroup
	ResourceTypeResearchStudy
	ResourceTypeResearchSubject
	ResourceTypeResource
	ResourceTypeRiskAssessment
	ResourceTypeSchedule
	ResourceTypeSearchParameter
	ResourceTypeSequence
	ResourceTypeServiceDefinition
	ResourceTypeSlot
	ResourceTypeSpecimen
	ResourceTypeStructureDefinition
	ResourceTypeStructureMap
	ResourceTypeSubscription
	ResourceTypeSubstance
	ResourceTypeSupplyDelivery
	ResourceTypeSupplyRequest
	ResourceTypeTask
	ResourceTypeTestReport
	ResourceTypeTestScript
	ResourceTypeValueSet
	ResourceTypeVisionPrescription
)

func (code ResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResourceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "Account":
		*code = ResourceTypeAccount
	case "ActivityDefinition":
		*code = ResourceTypeActivityDefinition
	case "AdverseEvent":
		*code = ResourceTypeAdverseEvent
	case "AllergyIntolerance":
		*code = ResourceTypeAllergyIntolerance
	case "Appointment":
		*code = ResourceTypeAppointment
	case "AppointmentResponse":
		*code = ResourceTypeAppointmentResponse
	case "AuditEvent":
		*code = ResourceTypeAuditEvent
	case "Basic":
		*code = ResourceTypeBasic
	case "Binary":
		*code = ResourceTypeBinary
	case "BodySite":
		*code = ResourceTypeBodySite
	case "Bundle":
		*code = ResourceTypeBundle
	case "CapabilityStatement":
		*code = ResourceTypeCapabilityStatement
	case "CarePlan":
		*code = ResourceTypeCarePlan
	case "CareTeam":
		*code = ResourceTypeCareTeam
	case "ChargeItem":
		*code = ResourceTypeChargeItem
	case "Claim":
		*code = ResourceTypeClaim
	case "ClaimResponse":
		*code = ResourceTypeClaimResponse
	case "ClinicalImpression":
		*code = ResourceTypeClinicalImpression
	case "CodeSystem":
		*code = ResourceTypeCodeSystem
	case "Communication":
		*code = ResourceTypeCommunication
	case "CommunicationRequest":
		*code = ResourceTypeCommunicationRequest
	case "CompartmentDefinition":
		*code = ResourceTypeCompartmentDefinition
	case "Composition":
		*code = ResourceTypeComposition
	case "ConceptMap":
		*code = ResourceTypeConceptMap
	case "Condition":
		*code = ResourceTypeCondition
	case "Consent":
		*code = ResourceTypeConsent
	case "Contract":
		*code = ResourceTypeContract
	case "Coverage":
		*code = ResourceTypeCoverage
	case "DataElement":
		*code = ResourceTypeDataElement
	case "DetectedIssue":
		*code = ResourceTypeDetectedIssue
	case "Device":
		*code = ResourceTypeDevice
	case "DeviceComponent":
		*code = ResourceTypeDeviceComponent
	case "DeviceMetric":
		*code = ResourceTypeDeviceMetric
	case "DeviceRequest":
		*code = ResourceTypeDeviceRequest
	case "DeviceUseStatement":
		*code = ResourceTypeDeviceUseStatement
	case "DiagnosticReport":
		*code = ResourceTypeDiagnosticReport
	case "DocumentManifest":
		*code = ResourceTypeDocumentManifest
	case "DocumentReference":
		*code = ResourceTypeDocumentReference
	case "DomainResource":
		*code = ResourceTypeDomainResource
	case "EligibilityRequest":
		*code = ResourceTypeEligibilityRequest
	case "EligibilityResponse":
		*code = ResourceTypeEligibilityResponse
	case "Encounter":
		*code = ResourceTypeEncounter
	case "Endpoint":
		*code = ResourceTypeEndpoint
	case "EnrollmentRequest":
		*code = ResourceTypeEnrollmentRequest
	case "EnrollmentResponse":
		*code = ResourceTypeEnrollmentResponse
	case "EpisodeOfCare":
		*code = ResourceTypeEpisodeOfCare
	case "ExpansionProfile":
		*code = ResourceTypeExpansionProfile
	case "ExplanationOfBenefit":
		*code = ResourceTypeExplanationOfBenefit
	case "FamilyMemberHistory":
		*code = ResourceTypeFamilyMemberHistory
	case "Flag":
		*code = ResourceTypeFlag
	case "Goal":
		*code = ResourceTypeGoal
	case "GraphDefinition":
		*code = ResourceTypeGraphDefinition
	case "Group":
		*code = ResourceTypeGroup
	case "GuidanceResponse":
		*code = ResourceTypeGuidanceResponse
	case "HealthcareService":
		*code = ResourceTypeHealthcareService
	case "ImagingManifest":
		*code = ResourceTypeImagingManifest
	case "ImagingStudy":
		*code = ResourceTypeImagingStudy
	case "Immunization":
		*code = ResourceTypeImmunization
	case "ImmunizationRecommendation":
		*code = ResourceTypeImmunizationRecommendation
	case "ImplementationGuide":
		*code = ResourceTypeImplementationGuide
	case "Library":
		*code = ResourceTypeLibrary
	case "Linkage":
		*code = ResourceTypeLinkage
	case "List":
		*code = ResourceTypeList
	case "Location":
		*code = ResourceTypeLocation
	case "Measure":
		*code = ResourceTypeMeasure
	case "MeasureReport":
		*code = ResourceTypeMeasureReport
	case "Media":
		*code = ResourceTypeMedia
	case "Medication":
		*code = ResourceTypeMedication
	case "MedicationAdministration":
		*code = ResourceTypeMedicationAdministration
	case "MedicationDispense":
		*code = ResourceTypeMedicationDispense
	case "MedicationRequest":
		*code = ResourceTypeMedicationRequest
	case "MedicationStatement":
		*code = ResourceTypeMedicationStatement
	case "MessageDefinition":
		*code = ResourceTypeMessageDefinition
	case "MessageHeader":
		*code = ResourceTypeMessageHeader
	case "NamingSystem":
		*code = ResourceTypeNamingSystem
	case "NutritionOrder":
		*code = ResourceTypeNutritionOrder
	case "Observation":
		*code = ResourceTypeObservation
	case "OperationDefinition":
		*code = ResourceTypeOperationDefinition
	case "OperationOutcome":
		*code = ResourceTypeOperationOutcome
	case "Organization":
		*code = ResourceTypeOrganization
	case "Parameters":
		*code = ResourceTypeParameters
	case "Patient":
		*code = ResourceTypePatient
	case "PaymentNotice":
		*code = ResourceTypePaymentNotice
	case "PaymentReconciliation":
		*code = ResourceTypePaymentReconciliation
	case "Person":
		*code = ResourceTypePerson
	case "PlanDefinition":
		*code = ResourceTypePlanDefinition
	case "Practitioner":
		*code = ResourceTypePractitioner
	case "PractitionerRole":
		*code = ResourceTypePractitionerRole
	case "Procedure":
		*code = ResourceTypeProcedure
	case "ProcedureRequest":
		*code = ResourceTypeProcedureRequest
	case "ProcessRequest":
		*code = ResourceTypeProcessRequest
	case "ProcessResponse":
		*code = ResourceTypeProcessResponse
	case "Provenance":
		*code = ResourceTypeProvenance
	case "Questionnaire":
		*code = ResourceTypeQuestionnaire
	case "QuestionnaireResponse":
		*code = ResourceTypeQuestionnaireResponse
	case "ReferralRequest":
		*code = ResourceTypeReferralRequest
	case "RelatedPerson":
		*code = ResourceTypeRelatedPerson
	case "RequestGroup":
		*code = ResourceTypeRequestGroup
	case "ResearchStudy":
		*code = ResourceTypeResearchStudy
	case "ResearchSubject":
		*code = ResourceTypeResearchSubject
	case "Resource":
		*code = ResourceTypeResource
	case "RiskAssessment":
		*code = ResourceTypeRiskAssessment
	case "Schedule":
		*code = ResourceTypeSchedule
	case "SearchParameter":
		*code = ResourceTypeSearchParameter
	case "Sequence":
		*code = ResourceTypeSequence
	case "ServiceDefinition":
		*code = ResourceTypeServiceDefinition
	case "Slot":
		*code = ResourceTypeSlot
	case "Specimen":
		*code = ResourceTypeSpecimen
	case "StructureDefinition":
		*code = ResourceTypeStructureDefinition
	case "StructureMap":
		*code = ResourceTypeStructureMap
	case "Subscription":
		*code = ResourceTypeSubscription
	case "Substance":
		*code = ResourceTypeSubstance
	case "SupplyDelivery":
		*code = ResourceTypeSupplyDelivery
	case "SupplyRequest":
		*code = ResourceTypeSupplyRequest
	case "Task":
		*code = ResourceTypeTask
	case "TestReport":
		*code = ResourceTypeTestReport
	case "TestScript":
		*code = ResourceTypeTestScript
	case "ValueSet":
		*code = ResourceTypeValueSet
	case "VisionPrescription":
		*code = ResourceTypeVisionPrescription
	default:
		return fmt.Errorf("unknown ResourceType code `%s`", s)
	}
	return nil
}
func (code ResourceType) String() string {
	return code.Code()
}
func (code ResourceType) Code() string {
	switch code {
	case ResourceTypeAccount:
		return "Account"
	case ResourceTypeActivityDefinition:
		return "ActivityDefinition"
	case ResourceTypeAdverseEvent:
		return "AdverseEvent"
	case ResourceTypeAllergyIntolerance:
		return "AllergyIntolerance"
	case ResourceTypeAppointment:
		return "Appointment"
	case ResourceTypeAppointmentResponse:
		return "AppointmentResponse"
	case ResourceTypeAuditEvent:
		return "AuditEvent"
	case ResourceTypeBasic:
		return "Basic"
	case ResourceTypeBinary:
		return "Binary"
	case ResourceTypeBodySite:
		return "BodySite"
	case ResourceTypeBundle:
		return "Bundle"
	case ResourceTypeCapabilityStatement:
		return "CapabilityStatement"
	case ResourceTypeCarePlan:
		return "CarePlan"
	case ResourceTypeCareTeam:
		return "CareTeam"
	case ResourceTypeChargeItem:
		return "ChargeItem"
	case ResourceTypeClaim:
		return "Claim"
	case ResourceTypeClaimResponse:
		return "ClaimResponse"
	case ResourceTypeClinicalImpression:
		return "ClinicalImpression"
	case ResourceTypeCodeSystem:
		return "CodeSystem"
	case ResourceTypeCommunication:
		return "Communication"
	case ResourceTypeCommunicationRequest:
		return "CommunicationRequest"
	case ResourceTypeCompartmentDefinition:
		return "CompartmentDefinition"
	case ResourceTypeComposition:
		return "Composition"
	case ResourceTypeConceptMap:
		return "ConceptMap"
	case ResourceTypeCondition:
		return "Condition"
	case ResourceTypeConsent:
		return "Consent"
	case ResourceTypeContract:
		return "Contract"
	case ResourceTypeCoverage:
		return "Coverage"
	case ResourceTypeDataElement:
		return "DataElement"
	case ResourceTypeDetectedIssue:
		return "DetectedIssue"
	case ResourceTypeDevice:
		return "Device"
	case ResourceTypeDeviceComponent:
		return "DeviceComponent"
	case ResourceTypeDeviceMetric:
		return "DeviceMetric"
	case ResourceTypeDeviceRequest:
		return "DeviceRequest"
	case ResourceTypeDeviceUseStatement:
		return "DeviceUseStatement"
	case ResourceTypeDiagnosticReport:
		return "DiagnosticReport"
	case ResourceTypeDocumentManifest:
		return "DocumentManifest"
	case ResourceTypeDocumentReference:
		return "DocumentReference"
	case ResourceTypeDomainResource:
		return "DomainResource"
	case ResourceTypeEligibilityRequest:
		return "EligibilityRequest"
	case ResourceTypeEligibilityResponse:
		return "EligibilityResponse"
	case ResourceTypeEncounter:
		return "Encounter"
	case ResourceTypeEndpoint:
		return "Endpoint"
	case ResourceTypeEnrollmentRequest:
		return "EnrollmentRequest"
	case ResourceTypeEnrollmentResponse:
		return "EnrollmentResponse"
	case ResourceTypeEpisodeOfCare:
		return "EpisodeOfCare"
	case ResourceTypeExpansionProfile:
		return "ExpansionProfile"
	case ResourceTypeExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case ResourceTypeFamilyMemberHistory:
		return "FamilyMemberHistory"
	case ResourceTypeFlag:
		return "Flag"
	case ResourceTypeGoal:
		return "Goal"
	case ResourceTypeGraphDefinition:
		return "GraphDefinition"
	case ResourceTypeGroup:
		return "Group"
	case ResourceTypeGuidanceResponse:
		return "GuidanceResponse"
	case ResourceTypeHealthcareService:
		return "HealthcareService"
	case ResourceTypeImagingManifest:
		return "ImagingManifest"
	case ResourceTypeImagingStudy:
		return "ImagingStudy"
	case ResourceTypeImmunization:
		return "Immunization"
	case ResourceTypeImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case ResourceTypeImplementationGuide:
		return "ImplementationGuide"
	case ResourceTypeLibrary:
		return "Library"
	case ResourceTypeLinkage:
		return "Linkage"
	case ResourceTypeList:
		return "List"
	case ResourceTypeLocation:
		return "Location"
	case ResourceTypeMeasure:
		return "Measure"
	case ResourceTypeMeasureReport:
		return "MeasureReport"
	case ResourceTypeMedia:
		return "Media"
	case ResourceTypeMedication:
		return "Medication"
	case ResourceTypeMedicationAdministration:
		return "MedicationAdministration"
	case ResourceTypeMedicationDispense:
		return "MedicationDispense"
	case ResourceTypeMedicationRequest:
		return "MedicationRequest"
	case ResourceTypeMedicationStatement:
		return "MedicationStatement"
	case ResourceTypeMessageDefinition:
		return "MessageDefinition"
	case ResourceTypeMessageHeader:
		return "MessageHeader"
	case ResourceTypeNamingSystem:
		return "NamingSystem"
	case ResourceTypeNutritionOrder:
		return "NutritionOrder"
	case ResourceTypeObservation:
		return "Observation"
	case ResourceTypeOperationDefinition:
		return "OperationDefinition"
	case ResourceTypeOperationOutcome:
		return "OperationOutcome"
	case ResourceTypeOrganization:
		return "Organization"
	case ResourceTypeParameters:
		return "Parameters"
	case ResourceTypePatient:
		return "Patient"
	case ResourceTypePaymentNotice:
		return "PaymentNotice"
	case ResourceTypePaymentReconciliation:
		return "PaymentReconciliation"
	case ResourceTypePerson:
		return "Person"
	case ResourceTypePlanDefinition:
		return "PlanDefinition"
	case ResourceTypePractitioner:
		return "Practitioner"
	case ResourceTypePractitionerRole:
		return "PractitionerRole"
	case ResourceTypeProcedure:
		return "Procedure"
	case ResourceTypeProcedureRequest:
		return "ProcedureRequest"
	case ResourceTypeProcessRequest:
		return "ProcessRequest"
	case ResourceTypeProcessResponse:
		return "ProcessResponse"
	case ResourceTypeProvenance:
		return "Provenance"
	case ResourceTypeQuestionnaire:
		return "Questionnaire"
	case ResourceTypeQuestionnaireResponse:
		return "QuestionnaireResponse"
	case ResourceTypeReferralRequest:
		return "ReferralRequest"
	case ResourceTypeRelatedPerson:
		return "RelatedPerson"
	case ResourceTypeRequestGroup:
		return "RequestGroup"
	case ResourceTypeResearchStudy:
		return "ResearchStudy"
	case ResourceTypeResearchSubject:
		return "ResearchSubject"
	case ResourceTypeResource:
		return "Resource"
	case ResourceTypeRiskAssessment:
		return "RiskAssessment"
	case ResourceTypeSchedule:
		return "Schedule"
	case ResourceTypeSearchParameter:
		return "SearchParameter"
	case ResourceTypeSequence:
		return "Sequence"
	case ResourceTypeServiceDefinition:
		return "ServiceDefinition"
	case ResourceTypeSlot:
		return "Slot"
	case ResourceTypeSpecimen:
		return "Specimen"
	case ResourceTypeStructureDefinition:
		return "StructureDefinition"
	case ResourceTypeStructureMap:
		return "StructureMap"
	case ResourceTypeSubscription:
		return "Subscription"
	case ResourceTypeSubstance:
		return "Substance"
	case ResourceTypeSupplyDelivery:
		return "SupplyDelivery"
	case ResourceTypeSupplyRequest:
		return "SupplyRequest"
	case ResourceTypeTask:
		return "Task"
	case ResourceTypeTestReport:
		return "TestReport"
	case ResourceTypeTestScript:
		return "TestScript"
	case ResourceTypeValueSet:
		return "ValueSet"
	case ResourceTypeVisionPrescription:
		return "VisionPrescription"
	}
	return "<unknown>"
}
func (code ResourceType) Display() string {
	switch code {
	case ResourceTypeAccount:
		return "Account"
	case ResourceTypeActivityDefinition:
		return "ActivityDefinition"
	case ResourceTypeAdverseEvent:
		return "AdverseEvent"
	case ResourceTypeAllergyIntolerance:
		return "AllergyIntolerance"
	case ResourceTypeAppointment:
		return "Appointment"
	case ResourceTypeAppointmentResponse:
		return "AppointmentResponse"
	case ResourceTypeAuditEvent:
		return "AuditEvent"
	case ResourceTypeBasic:
		return "Basic"
	case ResourceTypeBinary:
		return "Binary"
	case ResourceTypeBodySite:
		return "BodySite"
	case ResourceTypeBundle:
		return "Bundle"
	case ResourceTypeCapabilityStatement:
		return "CapabilityStatement"
	case ResourceTypeCarePlan:
		return "CarePlan"
	case ResourceTypeCareTeam:
		return "CareTeam"
	case ResourceTypeChargeItem:
		return "ChargeItem"
	case ResourceTypeClaim:
		return "Claim"
	case ResourceTypeClaimResponse:
		return "ClaimResponse"
	case ResourceTypeClinicalImpression:
		return "ClinicalImpression"
	case ResourceTypeCodeSystem:
		return "CodeSystem"
	case ResourceTypeCommunication:
		return "Communication"
	case ResourceTypeCommunicationRequest:
		return "CommunicationRequest"
	case ResourceTypeCompartmentDefinition:
		return "CompartmentDefinition"
	case ResourceTypeComposition:
		return "Composition"
	case ResourceTypeConceptMap:
		return "ConceptMap"
	case ResourceTypeCondition:
		return "Condition"
	case ResourceTypeConsent:
		return "Consent"
	case ResourceTypeContract:
		return "Contract"
	case ResourceTypeCoverage:
		return "Coverage"
	case ResourceTypeDataElement:
		return "DataElement"
	case ResourceTypeDetectedIssue:
		return "DetectedIssue"
	case ResourceTypeDevice:
		return "Device"
	case ResourceTypeDeviceComponent:
		return "DeviceComponent"
	case ResourceTypeDeviceMetric:
		return "DeviceMetric"
	case ResourceTypeDeviceRequest:
		return "DeviceRequest"
	case ResourceTypeDeviceUseStatement:
		return "DeviceUseStatement"
	case ResourceTypeDiagnosticReport:
		return "DiagnosticReport"
	case ResourceTypeDocumentManifest:
		return "DocumentManifest"
	case ResourceTypeDocumentReference:
		return "DocumentReference"
	case ResourceTypeDomainResource:
		return "DomainResource"
	case ResourceTypeEligibilityRequest:
		return "EligibilityRequest"
	case ResourceTypeEligibilityResponse:
		return "EligibilityResponse"
	case ResourceTypeEncounter:
		return "Encounter"
	case ResourceTypeEndpoint:
		return "Endpoint"
	case ResourceTypeEnrollmentRequest:
		return "EnrollmentRequest"
	case ResourceTypeEnrollmentResponse:
		return "EnrollmentResponse"
	case ResourceTypeEpisodeOfCare:
		return "EpisodeOfCare"
	case ResourceTypeExpansionProfile:
		return "ExpansionProfile"
	case ResourceTypeExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case ResourceTypeFamilyMemberHistory:
		return "FamilyMemberHistory"
	case ResourceTypeFlag:
		return "Flag"
	case ResourceTypeGoal:
		return "Goal"
	case ResourceTypeGraphDefinition:
		return "GraphDefinition"
	case ResourceTypeGroup:
		return "Group"
	case ResourceTypeGuidanceResponse:
		return "GuidanceResponse"
	case ResourceTypeHealthcareService:
		return "HealthcareService"
	case ResourceTypeImagingManifest:
		return "ImagingManifest"
	case ResourceTypeImagingStudy:
		return "ImagingStudy"
	case ResourceTypeImmunization:
		return "Immunization"
	case ResourceTypeImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case ResourceTypeImplementationGuide:
		return "ImplementationGuide"
	case ResourceTypeLibrary:
		return "Library"
	case ResourceTypeLinkage:
		return "Linkage"
	case ResourceTypeList:
		return "List"
	case ResourceTypeLocation:
		return "Location"
	case ResourceTypeMeasure:
		return "Measure"
	case ResourceTypeMeasureReport:
		return "MeasureReport"
	case ResourceTypeMedia:
		return "Media"
	case ResourceTypeMedication:
		return "Medication"
	case ResourceTypeMedicationAdministration:
		return "MedicationAdministration"
	case ResourceTypeMedicationDispense:
		return "MedicationDispense"
	case ResourceTypeMedicationRequest:
		return "MedicationRequest"
	case ResourceTypeMedicationStatement:
		return "MedicationStatement"
	case ResourceTypeMessageDefinition:
		return "MessageDefinition"
	case ResourceTypeMessageHeader:
		return "MessageHeader"
	case ResourceTypeNamingSystem:
		return "NamingSystem"
	case ResourceTypeNutritionOrder:
		return "NutritionOrder"
	case ResourceTypeObservation:
		return "Observation"
	case ResourceTypeOperationDefinition:
		return "OperationDefinition"
	case ResourceTypeOperationOutcome:
		return "OperationOutcome"
	case ResourceTypeOrganization:
		return "Organization"
	case ResourceTypeParameters:
		return "Parameters"
	case ResourceTypePatient:
		return "Patient"
	case ResourceTypePaymentNotice:
		return "PaymentNotice"
	case ResourceTypePaymentReconciliation:
		return "PaymentReconciliation"
	case ResourceTypePerson:
		return "Person"
	case ResourceTypePlanDefinition:
		return "PlanDefinition"
	case ResourceTypePractitioner:
		return "Practitioner"
	case ResourceTypePractitionerRole:
		return "PractitionerRole"
	case ResourceTypeProcedure:
		return "Procedure"
	case ResourceTypeProcedureRequest:
		return "ProcedureRequest"
	case ResourceTypeProcessRequest:
		return "ProcessRequest"
	case ResourceTypeProcessResponse:
		return "ProcessResponse"
	case ResourceTypeProvenance:
		return "Provenance"
	case ResourceTypeQuestionnaire:
		return "Questionnaire"
	case ResourceTypeQuestionnaireResponse:
		return "QuestionnaireResponse"
	case ResourceTypeReferralRequest:
		return "ReferralRequest"
	case ResourceTypeRelatedPerson:
		return "RelatedPerson"
	case ResourceTypeRequestGroup:
		return "RequestGroup"
	case ResourceTypeResearchStudy:
		return "ResearchStudy"
	case ResourceTypeResearchSubject:
		return "ResearchSubject"
	case ResourceTypeResource:
		return "Resource"
	case ResourceTypeRiskAssessment:
		return "RiskAssessment"
	case ResourceTypeSchedule:
		return "Schedule"
	case ResourceTypeSearchParameter:
		return "SearchParameter"
	case ResourceTypeSequence:
		return "Sequence"
	case ResourceTypeServiceDefinition:
		return "ServiceDefinition"
	case ResourceTypeSlot:
		return "Slot"
	case ResourceTypeSpecimen:
		return "Specimen"
	case ResourceTypeStructureDefinition:
		return "StructureDefinition"
	case ResourceTypeStructureMap:
		return "StructureMap"
	case ResourceTypeSubscription:
		return "Subscription"
	case ResourceTypeSubstance:
		return "Substance"
	case ResourceTypeSupplyDelivery:
		return "SupplyDelivery"
	case ResourceTypeSupplyRequest:
		return "SupplyRequest"
	case ResourceTypeTask:
		return "Task"
	case ResourceTypeTestReport:
		return "TestReport"
	case ResourceTypeTestScript:
		return "TestScript"
	case ResourceTypeValueSet:
		return "ValueSet"
	case ResourceTypeVisionPrescription:
		return "VisionPrescription"
	}
	return "<unknown>"
}
func (code ResourceType) Definition() string {
	switch code {
	case ResourceTypeAccount:
		return "A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc."
	case ResourceTypeActivityDefinition:
		return "This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context."
	case ResourceTypeAdverseEvent:
		return "Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death."
	case ResourceTypeAllergyIntolerance:
		return "Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance."
	case ResourceTypeAppointment:
		return "A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s)."
	case ResourceTypeAppointmentResponse:
		return "A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection."
	case ResourceTypeAuditEvent:
		return "A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage."
	case ResourceTypeBasic:
		return "Basic is used for handling concepts not yet defined in FHIR, narrative-only resources that don't map to an existing resource, and custom resources not appropriate for inclusion in the FHIR specification."
	case ResourceTypeBinary:
		return "A binary resource can contain any content, whether text, image, pdf, zip archive, etc."
	case ResourceTypeBodySite:
		return "Record details about the anatomical location of a specimen or body part.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case."
	case ResourceTypeBundle:
		return "A container for a collection of resources."
	case ResourceTypeCapabilityStatement:
		return "A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation."
	case ResourceTypeCarePlan:
		return "Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions."
	case ResourceTypeCareTeam:
		return "The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient."
	case ResourceTypeChargeItem:
		return "The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation."
	case ResourceTypeClaim:
		return "A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery."
	case ResourceTypeClaimResponse:
		return "This resource provides the adjudication details from the processing of a Claim resource."
	case ResourceTypeClinicalImpression:
		return "A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called \"ClinicalImpression\" rather than \"ClinicalAssessment\" to avoid confusion with the recording of assessment tools such as Apgar score."
	case ResourceTypeCodeSystem:
		return "A code system resource specifies a set of codes drawn from one or more code systems."
	case ResourceTypeCommunication:
		return "An occurrence of information being transmitted; e.g. an alert that was sent to a responsible provider, a public health agency was notified about a reportable condition."
	case ResourceTypeCommunicationRequest:
		return "A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition."
	case ResourceTypeCompartmentDefinition:
		return "A compartment definition that defines how resources are accessed on a server."
	case ResourceTypeComposition:
		return "A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained."
	case ResourceTypeConceptMap:
		return "A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models."
	case ResourceTypeCondition:
		return "A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern."
	case ResourceTypeConsent:
		return "A record of a healthcare consumer’s policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time."
	case ResourceTypeContract:
		return "A formal agreement between parties regarding the conduct of business, exchange of information or other matters."
	case ResourceTypeCoverage:
		return "Financial instrument which may be used to reimburse or pay for health care products and services."
	case ResourceTypeDataElement:
		return "The formal description of a single piece of information that can be gathered and reported."
	case ResourceTypeDetectedIssue:
		return "Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc."
	case ResourceTypeDevice:
		return "This resource identifies an instance or a type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.  Medical devices include durable (reusable) medical equipment, implantable devices, as well as disposable equipment used for diagnostic, treatment, and research for healthcare and public health.  Non-medical devices may include items such as a machine, cellphone, computer, application, etc."
	case ResourceTypeDeviceComponent:
		return "The characteristics, operational status and capabilities of a medical-related component of a medical device."
	case ResourceTypeDeviceMetric:
		return "Describes a measurement, calculation or setting capability of a medical device."
	case ResourceTypeDeviceRequest:
		return "Represents a request for a patient to employ a medical device. The device may be an implantable device, or an external assistive device, such as a walker."
	case ResourceTypeDeviceUseStatement:
		return "A record of a device being used by a patient where the record is the result of a report from the patient or another clinician."
	case ResourceTypeDiagnosticReport:
		return "The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports."
	case ResourceTypeDocumentManifest:
		return "A collection of documents compiled for a purpose together with metadata that applies to the collection."
	case ResourceTypeDocumentReference:
		return "A reference to a document."
	case ResourceTypeDomainResource:
		return "A resource that includes narrative, extensions, and contained resources."
	case ResourceTypeEligibilityRequest:
		return "The EligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an EligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy."
	case ResourceTypeEligibilityResponse:
		return "This resource provides eligibility and plan details from the processing of an Eligibility resource."
	case ResourceTypeEncounter:
		return "An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient."
	case ResourceTypeEndpoint:
		return "The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b or a REST endpoint for another FHIR server. This may include any security context information."
	case ResourceTypeEnrollmentRequest:
		return "This resource provides the insurance enrollment details to the insurer regarding a specified coverage."
	case ResourceTypeEnrollmentResponse:
		return "This resource provides enrollment and plan details from the processing of an Enrollment resource."
	case ResourceTypeEpisodeOfCare:
		return "An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time."
	case ResourceTypeExpansionProfile:
		return "Resource to define constraints on the Expansion of a FHIR ValueSet."
	case ResourceTypeExplanationOfBenefit:
		return "This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided."
	case ResourceTypeFamilyMemberHistory:
		return "Significant health events and conditions for a person related to the patient relevant in the context of care for the patient."
	case ResourceTypeFlag:
		return "Prospective warnings of potential issues when providing care to the patient."
	case ResourceTypeGoal:
		return "Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc."
	case ResourceTypeGraphDefinition:
		return "A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set."
	case ResourceTypeGroup:
		return "Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization."
	case ResourceTypeGuidanceResponse:
		return "A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken."
	case ResourceTypeHealthcareService:
		return "The details of a healthcare service available at a location."
	case ResourceTypeImagingManifest:
		return "A text description of the DICOM SOP instances selected in the ImagingManifest; or the reason for, or significance of, the selection."
	case ResourceTypeImagingStudy:
		return "Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities."
	case ResourceTypeImmunization:
		return "Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed."
	case ResourceTypeImmunizationRecommendation:
		return "A patient's point-in-time immunization and recommendation (i.e. forecasting a patient's immunization eligibility according to a published schedule) with optional supporting justification."
	case ResourceTypeImplementationGuide:
		return "A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts."
	case ResourceTypeLibrary:
		return "The Library resource is a general-purpose container for knowledge asset definitions. It can be used to describe and expose existing knowledge assets such as logic libraries and information model descriptions, as well as to describe a collection of knowledge assets."
	case ResourceTypeLinkage:
		return "Identifies two or more records (resource instances) that are referring to the same real-world \"occurrence\"."
	case ResourceTypeList:
		return "A set of information summarized from a list of other resources."
	case ResourceTypeLocation:
		return "Details and position information for a physical place where services are provided  and resources and participants may be stored, found, contained or accommodated."
	case ResourceTypeMeasure:
		return "The Measure resource provides the definition of a quality measure."
	case ResourceTypeMeasureReport:
		return "The MeasureReport resource contains the results of evaluating a measure."
	case ResourceTypeMedia:
		return "A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference."
	case ResourceTypeMedication:
		return "This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication."
	case ResourceTypeMedicationAdministration:
		return "Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner."
	case ResourceTypeMedicationDispense:
		return "Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order."
	case ResourceTypeMedicationRequest:
		return "An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called \"MedicationRequest\" rather than \"MedicationPrescription\" or \"MedicationOrder\" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns."
	case ResourceTypeMedicationStatement:
		return "A record of a medication that is being consumed by a patient.   A MedicationStatement may indicate that the patient may be taking the medication now, or has taken the medication in the past or will be taking the medication in the future.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay.   The medication information may come from sources such as the patient's memory, from a prescription bottle,  or from a list of medications the patient, clinician or other party maintains \r\rThe primary difference between a medication statement and a medication administration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication.  A medication statement is often, if not always, less specific.  There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise.  As stated earlier, the medication statement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party maintains.  Medication administration is more formal and is not missing detailed information."
	case ResourceTypeMessageDefinition:
		return "Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted."
	case ResourceTypeMessageHeader:
		return "The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle."
	case ResourceTypeNamingSystem:
		return "A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a \"System\" used within the Identifier and Coding data types."
	case ResourceTypeNutritionOrder:
		return "A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident."
	case ResourceTypeObservation:
		return "Measurements and simple assertions made about a patient, device or other subject."
	case ResourceTypeOperationDefinition:
		return "A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction)."
	case ResourceTypeOperationOutcome:
		return "A collection of error, warning or information messages that result from a system action."
	case ResourceTypeOrganization:
		return "A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.  Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, etc."
	case ResourceTypeParameters:
		return "This special resource type is used to represent an operation request and response (operations.html). It has no other use, and there is no RESTful endpoint associated with it."
	case ResourceTypePatient:
		return "Demographics and other administrative information about an individual or animal receiving care or other health-related services."
	case ResourceTypePaymentNotice:
		return "This resource provides the status of the payment for goods and services rendered, and the request and response resource references."
	case ResourceTypePaymentReconciliation:
		return "This resource provides payment details and claim references supporting a bulk payment."
	case ResourceTypePerson:
		return "Demographics and administrative information about a person independent of a specific health-related context."
	case ResourceTypePlanDefinition:
		return "This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols."
	case ResourceTypePractitioner:
		return "A person who is directly or indirectly involved in the provisioning of healthcare."
	case ResourceTypePractitionerRole:
		return "A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time."
	case ResourceTypeProcedure:
		return "An action that is or was performed on a patient. This can be a physical intervention like an operation, or less invasive like counseling or hypnotherapy."
	case ResourceTypeProcedureRequest:
		return "A record of a request for diagnostic investigations, treatments, or operations to be performed."
	case ResourceTypeProcessRequest:
		return "This resource provides the target, request and response, and action details for an action to be performed by the target on or about existing resources."
	case ResourceTypeProcessResponse:
		return "This resource provides processing status, errors and notes from the processing of a resource."
	case ResourceTypeProvenance:
		return "Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies."
	case ResourceTypeQuestionnaire:
		return "A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection."
	case ResourceTypeQuestionnaireResponse:
		return "A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to."
	case ResourceTypeReferralRequest:
		return "Used to record and send details about a request for referral service or transfer of a patient to the care of another provider or provider organization."
	case ResourceTypeRelatedPerson:
		return "Information about a person that is involved in the care for a patient, but who is not the target of healthcare, nor has a formal responsibility in the care process."
	case ResourceTypeRequestGroup:
		return "A group of related requests that can be used to capture intended activities that have inter-dependencies such as \"give this medication after that one\"."
	case ResourceTypeResearchStudy:
		return "A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects."
	case ResourceTypeResearchSubject:
		return "A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects."
	case ResourceTypeResource:
		return "This is the base resource type for everything."
	case ResourceTypeRiskAssessment:
		return "An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome."
	case ResourceTypeSchedule:
		return "A container for slots of time that may be available for booking appointments."
	case ResourceTypeSearchParameter:
		return "A search parameter that defines a named search item that can be used to search/filter on a resource."
	case ResourceTypeSequence:
		return "Raw data describing a biological sequence."
	case ResourceTypeServiceDefinition:
		return "The ServiceDefinition describes a unit of decision support functionality that is made available as a service, such as immunization modules or drug-drug interaction checking."
	case ResourceTypeSlot:
		return "A slot of time on a schedule that may be available for booking appointments."
	case ResourceTypeSpecimen:
		return "A sample to be used for analysis."
	case ResourceTypeStructureDefinition:
		return "A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types."
	case ResourceTypeStructureMap:
		return "A Map of relationships between 2 structures that can be used to transform data."
	case ResourceTypeSubscription:
		return "The subscription resource is used to define a push based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined \"channel\" so that another system is able to take an appropriate action."
	case ResourceTypeSubstance:
		return "A homogeneous material with a definite composition."
	case ResourceTypeSupplyDelivery:
		return "Record of delivery of what is supplied."
	case ResourceTypeSupplyRequest:
		return "A record of a request for a medication, substance or device used in the healthcare setting."
	case ResourceTypeTask:
		return "A task to be performed."
	case ResourceTypeTestReport:
		return "A summary of information based on the results of executing a TestScript."
	case ResourceTypeTestScript:
		return "A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification."
	case ResourceTypeValueSet:
		return "A value set specifies a set of codes drawn from one or more code systems."
	case ResourceTypeVisionPrescription:
		return "An authorization for the supply of glasses and/or contact lenses to a patient."
	}
	return "<unknown>"
}
