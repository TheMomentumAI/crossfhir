package cmd

import (
	"context"
	"crossfhir/internal"
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

var healthlakeClient *healthlake.Client
var s3Client *s3.Client
var pgConn *pgx.Conn

var rootCmd = &cobra.Command{
	Use:   "crossfhir",
	Short: "crossfhir is a CLI for converting AWS Health Lake FHIR data to PostgreSQL",
}

func Execute() {
	LoadEnv()
	ConfigAWSClient()

	pgConn, _ = internal.InitConnection()
	internal.RunFhirMigration(pgConn)

	// internal.ExecQuery(pgConn, "SELECT id FROM patient")

	rootCmd.AddCommand(ExportCmd())
	rootCmd.AddCommand(DescribeCmd())
	rootCmd.AddCommand(FetchCmd())
	rootCmd.AddCommand(ConvertCmd())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	defer pgConn.Close(context.Background())

	if err != nil {
		os.Exit(1)
	}
}

func LoadEnv() {
	envFile := ".env"
	rootCmd.PersistentFlags().StringVar(&envFile, "env-file", ".env", "environment file to load")

	// check whether AWS creds exists

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConfigAWSClient() {
	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")

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
