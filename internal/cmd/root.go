package cmd

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	healthlakeClient *healthlake.Client
	s3Client         *s3.Client
	cfg              Config
	dir              string
	// verbose          bool
)

type Config struct {
	// from Env
	AwsAccessKey        string // e.g. "AKIA..."
	AwsSecretKey        string // e.g. "Tdaz4e..."
	AwsRegion           string // e.g. "us-east-1"
	AwsS3Bucket         string // e.g. "s3://my-bucket"
	AwsIAMExportRole    string // e.g. "arn:aws:iam::123123123:role/IAMRole"
	AwsDatastoreId      string // e.g. "8699acc...c49744168"
	AwsKmsKeyId         string // e.g. "arn:aws:kms:us-east-1:123123123:key/749b1e97-85db-49af5"
	AwsExportJobName    string // e.g. "my-export-job"
	AwsDatastoreFHIRUrl string // e.g. "https://healthlake.us-east-1.amazonaws.com"
	DbHost              string // e.g. "localhost"
	DbPort              string // e.g. "5432"
	DbUsername          string // e.g. "postgres"
	DbPassword          string // e.g. "password"
	DbDatabase          string // e.g. "postgres"

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
	loadEnv()
	configAWSClient()

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

func loadEnv() {
	envFile := ".env"
	rootCmd.PersistentFlags().StringVar(&envFile, "env-file", ".env", "environment file to load")

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	missingEnvs := []string{}
	missingEnvs = validateRootEnvs(missingEnvs)

	if len(missingEnvs) > 0 {
		log.Println("Missing required environment variables:")
		for _, envVar := range missingEnvs {
			log.Printf("%s\n", envVar)
		}

		os.Exit(1)
	}
}

func validateRootEnvs(missingEnvs []string) []string {
	cfg.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	if cfg.AwsAccessKey == "" {
		missingEnvs = append(missingEnvs, "AWS_ACCESS_KEY")
	}

	cfg.AwsSecretKey = os.Getenv("AWS_SECRET_KEY")
	if cfg.AwsSecretKey == "" {
		missingEnvs = append(missingEnvs, "AWS_SECRET_KEY")
	}

	cfg.AwsRegion = os.Getenv("AWS_REGION")
	if cfg.AwsRegion == "" {
		missingEnvs = append(missingEnvs, "AWS_REGION")
	}

	return missingEnvs
}

func configAWSClient() {
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
