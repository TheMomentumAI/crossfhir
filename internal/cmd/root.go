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
