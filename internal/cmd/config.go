package cmd

type Config struct {
	Aws   AwsConfig   `toml:"aws"`
	Db    DbConfig    `toml:"db"`
	Smart SmartConfig `toml:"smart"`
}

type AwsConfig struct {
	AccessKey        string `toml:"aws_access_key"`         // e.g. "AKIA..."
	SecretKey        string `toml:"aws_secret_key"`         // e.g. "Tdaz4e..."
	Region           string `toml:"aws_region"`             // e.g. "us-east-1"
	S3Bucket         string `toml:"aws_s3_bucket"`          // e.g. "s3://my-bucket"
	IAMExportRole    string `toml:"aws_iam_export_role"`    // e.g. "arn:aws:iam::123123123:role/IAMRole"
	DatastoreId      string `toml:"aws_datastore_id"`       // e.g. "8699acc...c49744168"
	KmsKeyId         string `toml:"aws_kms_key_id"`         // e.g. "arn:aws:kms:us-east-1:123123123:key/749b1e97-85db-49af5"
	ExportJobName    string `toml:"aws_export_job_name"`    // e.g. "my-export-job"
	DatastoreFHIRUrl string `toml:"aws_datastore_fhir_url"` // e.g. "https://healthlake.us-east-1.amazonaws.com"
	// from code
	ExportJobId       string `toml:"-"`
	ExportJobStatus   string `toml:"-"`
	ExportJobS3Output string `toml:"-"`
}

type DbConfig struct {
	Host     string `toml:"db_host"`     // e.g. "localhost"
	Port     string `toml:"db_port"`     // e.g. "5432"
	Username string `toml:"db_username"` // e.g. "postgres"
	Password string `toml:"db_password"` // e.g. "password"
	Database string `toml:"db_database"` // e.g. "postgres"
}

type SmartConfig struct {
	ClientID          string `toml:"smart_client_id"`          // OAuth2 client ID from your SMART app registration
	ClientSecret      string `toml:"smart_client_secret"`      // OAuth2 client secret from your SMART app registration
	AuthURL           string `toml:"smart_auth_url"`           // Authorization endpoint URL from your IdP
	TokenURL          string `toml:"smart_token_url"`          // Token endpoint URL from your IdP
	CallbackURL       string `toml:"smart_callback_url"`       // OAuth2 callback URL (e.g. "https://localhost")
	DatastoreEndpoint string `toml:"smart_datastore_endpoint"` // HealthLake FHIR endpoint URL
	Scope             string `toml:"smart_scope"`              // OAuth2 scopes (e.g. "launch/patient patient/*.read")
	GrantType         string `toml:"smart_grant_type"`         // OAuth2 grant type (e.g. "client_credentials")
	// Runtime fields (not in TOML)
	ExportJobId string `toml:"-"` // Tracks the export job ID
}
