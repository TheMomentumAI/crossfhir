package internal


var FhirSQLCommands = []string{
	`CREATE EXTENSION IF NOT EXISTS pgcrypto;`,
	`CREATE TABLE IF NOT EXISTS devicerequest_history (id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS servicerequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ServiceRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS servicerequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ServiceRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS devicemetric (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceMetric',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS devicemetric_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceMetric',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS careplan (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CarePlan',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS careplan_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CarePlan',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS observation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Observation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS observation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Observation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS enrollmentrequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EnrollmentRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS enrollmentrequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EnrollmentRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS "group" (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Group',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS group_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Group',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS messagedefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MessageDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS messagedefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MessageDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS appointment (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Appointment',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS appointment_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Appointment',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS biologicallyderivedproduct (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'BiologicallyDerivedProduct',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS biologicallyderivedproduct_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'BiologicallyDerivedProduct',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS questionnaireresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'QuestionnaireResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS questionnaireresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'QuestionnaireResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS effectevidencesynthesis (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EffectEvidenceSynthesis',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS effectevidencesynthesis_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EffectEvidenceSynthesis',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductcontraindication (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductContraindication',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductcontraindication_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductContraindication',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS episodeofcare (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EpisodeOfCare',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS episodeofcare_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EpisodeOfCare',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS evidence (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Evidence',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS evidence_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Evidence',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substancepolymer (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstancePolymer',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substancepolymer_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstancePolymer',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS supplydelivery (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SupplyDelivery',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS supplydelivery_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SupplyDelivery',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substancenucleicacid (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceNucleicAcid',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substancenucleicacid_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceNucleicAcid',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS adverseevent (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AdverseEvent',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS adverseevent_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AdverseEvent',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS endpoint (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Endpoint',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS endpoint_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Endpoint',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substancereferenceinformation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceReferenceInformation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substancereferenceinformation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceReferenceInformation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substancesourcematerial (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceSourceMaterial',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substancesourcematerial_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceSourceMaterial',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS compartmentdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CompartmentDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS compartmentdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CompartmentDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS detectedissue (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DetectedIssue',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS detectedissue_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DetectedIssue',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicationadministration (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationAdministration',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicationadministration_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationAdministration',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS evidencevariable (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EvidenceVariable',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS evidencevariable_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EvidenceVariable',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS implementationguide (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImplementationGuide',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS implementationguide_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImplementationGuide',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS goal (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Goal',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS goal_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Goal',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS communication (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Communication',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS communication_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Communication',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS schedule (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Schedule',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS schedule_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Schedule',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS documentreference (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DocumentReference',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS documentreference_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DocumentReference',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS organizationaffiliation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OrganizationAffiliation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS organizationaffiliation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OrganizationAffiliation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS devicedefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS devicedefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS coverage (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Coverage',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS coverage_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Coverage',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS auditevent (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AuditEvent',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS auditevent_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AuditEvent',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS messageheader (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MessageHeader',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS messageheader_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MessageHeader',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS contract (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Contract',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS contract_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Contract',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS testreport (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'TestReport',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS testreport_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'TestReport',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS codesystem (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CodeSystem',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS codesystem_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CodeSystem',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS plandefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PlanDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS plandefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PlanDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS invoice (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Invoice',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS invoice_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Invoice',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS claimresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ClaimResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS claimresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ClaimResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS chargeitem (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ChargeItem',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS chargeitem_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ChargeItem',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS coverageeligibilityresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CoverageEligibilityResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS coverageeligibilityresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CoverageEligibilityResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS bodystructure (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'BodyStructure',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS bodystructure_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'BodyStructure',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS parameters (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Parameters',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS parameters_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Parameters',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS clinicalimpression (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ClinicalImpression',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS clinicalimpression_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ClinicalImpression',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS familymemberhistory (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'FamilyMemberHistory',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS familymemberhistory_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'FamilyMemberHistory',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductauthorization (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductAuthorization',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductauthorization_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductAuthorization',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS "binary" (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Binary',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS binary_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Binary',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS composition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Composition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS composition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Composition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS practitionerrole (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PractitionerRole',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS practitionerrole_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PractitionerRole',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS healthcareservice (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'HealthcareService',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS healthcareservice_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'HealthcareService',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS patient (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Patient',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS patient_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Patient',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicationdispense (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationDispense',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicationdispense_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationDispense',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS deviceusestatement (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceUseStatement',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS deviceusestatement_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DeviceUseStatement',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS structuremap (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'StructureMap',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS structuremap_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'StructureMap',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS immunizationevaluation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImmunizationEvaluation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS immunizationevaluation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImmunizationEvaluation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS library (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Library',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS library_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Library',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS basic (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Basic',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS basic_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Basic',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS slot (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Slot',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS slot_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Slot',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS activitydefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ActivityDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS activitydefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ActivityDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductinteraction (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductInteraction',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductinteraction_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductInteraction',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS molecularsequence (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MolecularSequence',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS molecularsequence_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MolecularSequence',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS specimen (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Specimen',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS specimen_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Specimen',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS diagnosticreport (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DiagnosticReport',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS diagnosticreport_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DiagnosticReport',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS subscription (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Subscription',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS subscription_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Subscription',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS requestgroup (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RequestGroup',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS requestgroup_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RequestGroup',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS provenance (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Provenance',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS provenance_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Provenance',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproduct (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProduct',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproduct_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProduct',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS chargeitemdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ChargeItemDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS chargeitemdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ChargeItemDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS practitioner (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Practitioner',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS practitioner_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Practitioner',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductpackaged (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductPackaged',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductpackaged_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductPackaged',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS flag (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Flag',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS flag_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Flag',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS explanationofbenefit (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ExplanationOfBenefit',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS explanationofbenefit_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ExplanationOfBenefit',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS linkage (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Linkage',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS linkage_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Linkage',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS operationoutcome (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OperationOutcome',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS operationoutcome_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OperationOutcome',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductpharmaceutical (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductPharmaceutical',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductpharmaceutical_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductPharmaceutical',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS immunization (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Immunization',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS immunization_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Immunization',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicationknowledge (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationKnowledge',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicationknowledge_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationKnowledge',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS researchsubject (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchSubject',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS researchsubject_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchSubject',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductindication (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductIndication',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductindication_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductIndication',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS paymentnotice (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PaymentNotice',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS paymentnotice_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PaymentNotice',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS namingsystem (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'NamingSystem',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS namingsystem_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'NamingSystem',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicationstatement (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationStatement',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicationstatement_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationStatement',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS enrollmentresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EnrollmentResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS enrollmentresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EnrollmentResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS nutritionorder (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'NutritionOrder',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS nutritionorder_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'NutritionOrder',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS questionnaire (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Questionnaire',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS questionnaire_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Questionnaire',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS account (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Account',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS account_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Account',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS eventdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EventDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS eventdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'EventDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductundesirableeffect (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductUndesirableEffect',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductundesirableeffect_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductUndesirableEffect',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substancespecification (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceSpecification',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substancespecification_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceSpecification',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS communicationrequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CommunicationRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS communicationrequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CommunicationRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS specimendefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SpecimenDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS specimendefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SpecimenDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS verificationresult (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'VerificationResult',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS verificationresult_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'VerificationResult',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS documentmanifest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DocumentManifest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS documentmanifest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'DocumentManifest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS task (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Task',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS task_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Task',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS riskevidencesynthesis (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RiskEvidenceSynthesis',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS riskevidencesynthesis_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RiskEvidenceSynthesis',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS valueset (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ValueSet',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS valueset_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ValueSet',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS claim (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Claim',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS claim_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Claim',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS insuranceplan (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'InsurancePlan',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS insuranceplan_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'InsurancePlan',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS examplescenario (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ExampleScenario',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS examplescenario_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ExampleScenario',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS researchstudy (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchStudy',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS researchstudy_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchStudy',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicationrequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicationrequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicationRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS measure (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Measure',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS measure_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Measure',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS list (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'List',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS list_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'List',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS encounter (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Encounter',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS encounter_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Encounter',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS capabilitystatement (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CapabilityStatement',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS capabilitystatement_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CapabilityStatement',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS visionprescription (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'VisionPrescription',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS visionprescription_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'VisionPrescription',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS riskassessment (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RiskAssessment',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS riskassessment_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RiskAssessment',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substanceprotein (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceProtein',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substanceprotein_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SubstanceProtein',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS immunizationrecommendation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImmunizationRecommendation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS immunizationrecommendation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImmunizationRecommendation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS relatedperson (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RelatedPerson',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS relatedperson_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'RelatedPerson',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medication (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Medication',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medication_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Medication',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS appointmentresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AppointmentResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS appointmentresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AppointmentResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS researchelementdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchElementDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS researchelementdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchElementDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS substance (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Substance',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS substance_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Substance',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS paymentreconciliation (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PaymentReconciliation',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS paymentreconciliation_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'PaymentReconciliation',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS conceptmap (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ConceptMap',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS conceptmap_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ConceptMap',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS person (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Person',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS person_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Person',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS condition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Condition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS condition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Condition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS careteam (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CareTeam',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS careteam_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CareTeam',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS catalogentry (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CatalogEntry',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS catalogentry_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CatalogEntry',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS structuredefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'StructureDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS structuredefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'StructureDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS procedure (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Procedure',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS procedure_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Procedure',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS consent (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Consent',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS consent_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Consent',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS observationdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ObservationDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS observationdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ObservationDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS attribute (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Attribute',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS attribute_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Attribute',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS location (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Location',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS location_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Location',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS organization (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Organization',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS organization_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Organization',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS device (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Device',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS device_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Device',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS supplyrequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SupplyRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS supplyrequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'SupplyRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS allergyintolerance (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AllergyIntolerance',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS allergyintolerance_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'AllergyIntolerance',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS researchdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS researchdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ResearchDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS operationdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OperationDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS operationdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'OperationDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductmanufactured (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductManufactured',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductmanufactured_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductManufactured',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS imagingstudy (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImagingStudy',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS imagingstudy_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'ImagingStudy',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS coverageeligibilityrequest (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CoverageEligibilityRequest',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS coverageeligibilityrequest_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'CoverageEligibilityRequest',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductingredient (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductIngredient',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS medicinalproductingredient_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MedicinalProductIngredient',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS guidanceresponse (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'GuidanceResponse',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS guidanceresponse_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'GuidanceResponse',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS media (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Media',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS media_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'Media',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS measurereport (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MeasureReport',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS measurereport_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MeasureReport',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS graphdefinition (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'GraphDefinition',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS graphdefinition_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'GraphDefinition',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS terminologycapabilities (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'TerminologyCapabilities',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS terminologycapabilities_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'TerminologyCapabilities',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
	`CREATE TABLE IF NOT EXISTS metadataresource (
	  id text primary key,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MetadataResource',
	  resource jsonb not null
	);`,
	`CREATE TABLE IF NOT EXISTS metadataresource_history (
	  id text,
	  txid bigint not null,
	  ts timestamptz DEFAULT current_timestamp,
	  resource_type text default 'MetadataResource',
	  resource jsonb not null,
	  PRIMARY KEY (id, txid)
	);`,
}
