package cmd

// Config represents the main configuration structure for crossfhir.
// It contains all necessary settings for AWS, database, and SMART on FHIR connectivity.
type Config struct {
	Aws   AwsConfig   `toml:"aws"`   // AWS-related configuration
	Db    DbConfig    `toml:"db"`    // Database-related configuration
	Smart SmartConfig `toml:"smart"` // SMART on FHIR-related configuration
}

// AwsConfig holds AWS-specific configuration parameters including credentials
// and service endpoints.
type AwsConfig struct {
	// AccessKey is the AWS access key ID
	AccessKey string `toml:"aws_access_key"`

	// SecretKey is the AWS secret access key
	SecretKey string `toml:"aws_secret_key"`

	// Region specifies the AWS region (e.g., "us-east-1")
	Region string `toml:"aws_region"`

	// S3Bucket is the S3 bucket URL (e.g., "s3://my-bucket")
	S3Bucket string `toml:"aws_s3_bucket"`

	// IAMExportRole is the ARN of the IAM role used for exports
	IAMExportRole string `toml:"aws_iam_export_role"`

	// DatastoreId is the AWS HealthLake FHIR datastore ID
	DatastoreId string `toml:"aws_datastore_id"`

	// KmsKeyId is the ARN of the KMS key used for S3 bucket encryption
	KmsKeyId string `toml:"aws_kms_key_id"`

	// ExportJobName is a user-defined name for the export job
	ExportJobName string `toml:"aws_export_job_name"`

	// DatastoreFHIRUrl is the base URL for the FHIR datastore
	DatastoreFHIRUrl string `toml:"aws_datastore_fhir_url"`

	// Runtime fields (not in TOML configuration)
	ExportJobId       string `toml:"-"` // ID of the current export job
	ExportJobStatus   string `toml:"-"` // Status of the current export job
	ExportJobS3Output string `toml:"-"` // S3 output location of the export
}

// DbConfig holds PostgreSQL database connection parameters.
type DbConfig struct {
	// Host is the database server hostname (default: "localhost")
	Host string `toml:"db_host"`

	// Port is the database server port (default: "5432")
	Port string `toml:"db_port"`

	// Username for database authentication
	Username string `toml:"db_username"`

	// Password for database authentication
	Password string `toml:"db_password"`

	// Database name to connect to
	Database string `toml:"db_database"`
}

type SmartConfig struct {
	// ClientID from SMART app registration
	ClientID string `toml:"smart_client_id"`

	// ClientSecret from SMART app registration
	ClientSecret string `toml:"smart_client_secret"`

	// TokenURL for OAuth2 token endpoint
	TokenURL string `toml:"smart_token_url"`

	// DatastoreEndpoint for the FHIR server
	DatastoreEndpoint string `toml:"smart_datastore_endpoint"`

	// Scope for OAuth2 authentication
	Scope string `toml:"smart_scope"`

	// GrantType for OAuth2 (e.g., "client_credentials")
	GrantType string `toml:"smart_grant_type"`

	// Runtime fields
	ExportJobId string `toml:"-"` // Tracks the export job ID
}
