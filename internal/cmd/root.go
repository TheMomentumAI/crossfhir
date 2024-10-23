package cmd

import (
	"context"
	// "crossfhir/internal"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	healthlakeClient *healthlake.Client
	s3Client         *s3.Client
	pgConn           *pgx.Conn
	cfg              Config
	dir              string
	// verbose          bool
)

type Config struct {
	// from Env
	AwsAccessKey     string // e.g. "AKIA..."
	AwsSecretKey     string // e.g. "Tdaz4e..."
	AwsRegion        string // e.g. "us-east-1"
	AwsS3Bucket      string // e.g. "s3://my-bucket"
	AwsIAMExportRole string // e.g. "arn:aws:iam::123123123:role/IAMRole"
	AwsDatastoreId   string // e.g. "8699acc...c49744168"
	AwsKmsKeyId      string // e.g. "arn:aws:kms:us-east-1:123123123:key/749b1e97-85db-49af5"
	AwsExportJobName string // e.g. "my-export-job"

	// from code
	AwsExportJobId       string
	AwsExportJobStatus   string
	AwsExportJobS3Output string
}

var rootCmd = &cobra.Command{
	Use:   "crossfhir",
	Short: "crossfhir is a CLI for converting AWS Health Lake FHIR data to PostgreSQL",
}

func Execute() {
	LoadEnv()
	ConfigAWSClient()

	// pgConn, _ = internal.InitConnection()
	// internal.RunFhirMigration(pgConn)
	// internal.ExecQuery(pgConn, "SELECT id FROM patient")

	rootCmd.AddCommand(ExportCmd())
	rootCmd.AddCommand(PullCmd())
	rootCmd.AddCommand(ConvertCmd())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	// defer pgConn.Close(context.Background())

	if err != nil {
		os.Exit(1)
	}
}

func LoadEnv() {
	envFile := ".env"
	rootCmd.PersistentFlags().StringVar(&envFile, "env-file", ".env", "environment file to load")

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	missingEnvVars := []string{}

	missingEnvVars = ValidateEnvs(missingEnvVars)

	if len(missingEnvVars) > 0 {
		log.Println("Missing required environment variables:")
		for _, envVar := range missingEnvVars {
			fmt.Printf("%s\n", envVar)
		}

		os.Exit(1)
	}
}

func ValidateEnvs(missingEnvVars []string) []string {
	cfg.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	if cfg.AwsAccessKey == "" {
		missingEnvVars = append(missingEnvVars, "AWS_ACCESS_KEY")
	}

	cfg.AwsSecretKey = os.Getenv("AWS_SECRET_KEY")
	if cfg.AwsSecretKey == "" {
		missingEnvVars = append(missingEnvVars, "AWS_SECRET_KEY")
	}

	cfg.AwsRegion = os.Getenv("AWS_REGION")
	if cfg.AwsRegion == "" {
		missingEnvVars = append(missingEnvVars, "AWS_REGION")
	}

	cfg.AwsS3Bucket = os.Getenv("AWS_S3_BUCKET")
	if cfg.AwsS3Bucket == "" {
		missingEnvVars = append(missingEnvVars, "AWS_S3_BUCKET")
	}

	cfg.AwsIAMExportRole = os.Getenv("AWS_IAM_EXPORT_ROLE")
	if cfg.AwsIAMExportRole == "" {
		missingEnvVars = append(missingEnvVars, "AWS_IAM_EXPORT_ROLE")
	}

	cfg.AwsDatastoreId = os.Getenv("AWS_DATASTORE_ID")
	if cfg.AwsDatastoreId == "" {
		missingEnvVars = append(missingEnvVars, "AWS_DATASTORE_ID")
	}

	cfg.AwsKmsKeyId = os.Getenv("AWS_KMS_KEY_ID_ARN")
	if cfg.AwsKmsKeyId == "" {
		missingEnvVars = append(missingEnvVars, "AWS_KMS_KEY_ID_ARN")
	}

	cfg.AwsExportJobName = os.Getenv("AWS_EXPORT_JOB_NAME")
	if cfg.AwsExportJobName == "" {
		missingEnvVars = append(missingEnvVars, "AWS_EXPORT_JOB_NAME")
	}
	return missingEnvVars
}

func ConfigAWSClient() {
	creds := credentials.NewStaticCredentialsProvider(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_SECRET_KEY"),
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	healthlakeClient = healthlake.NewFromConfig(cfg)
	s3Client = s3.NewFromConfig(cfg)
}
