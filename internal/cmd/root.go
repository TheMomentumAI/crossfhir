package cmd

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

const (
	version          = "0.1-beta"
	shortDescription = `
crossfhir is a CLI tool for converting AWS HealthLake FHIR data to PostgreSQL
and interacting with the HealthLake FHIR REST API.
`
	examples = `
# Exporting data with pull into local directory
crossfhir export --pull --dir ./mydirectory

# Loading data from local directory to PostgreSQL with migration
crossfhir load -m --data ./mydirectory
`
)

var (
	healthlakeClient *healthlake.Client
	s3Client         *s3.Client
	cfg              Config
	dir              string
	configFile       string
	// verbose          bool
)

type Config struct {
	Aws AwsConfig `toml:"aws"`
	Db  DbConfig  `toml:"db"`
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

var rootCmd = &cobra.Command{
	Use:     "crossfhir",
	Short:   shortDescription,
	Example: examples,
	Version: version,
}

func init() {
	cobra.OnInitialize(loadConfig, configAWSClient)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.toml", "crossfhir config file path")
}

func Execute() {
	rootCmd.AddCommand(ExportCmd())
	rootCmd.AddCommand(PullCmd())
	rootCmd.AddCommand(LoadCmd())
	rootCmd.AddCommand(RestCmd())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	if err != nil {
		os.Exit(1)
	}
}

func loadConfig() {
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}

// TODO use smart on fhir
func configAWSClient() {
	creds := credentials.NewStaticCredentialsProvider(
		cfg.Aws.AccessKey,
		cfg.Aws.SecretKey,
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(cfg.Aws.Region),
	)

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	healthlakeClient = healthlake.NewFromConfig(cfg)
	s3Client = s3.NewFromConfig(cfg)
}
