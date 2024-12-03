package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"crossfhir/internal/helpers"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	pull      bool
	awsExport bool
)

type ExportRequestBody struct {
	DataAccessRoleArn string           `json:"DataAccessRoleArn"`
	OutputDataConfig  OutputDataConfig `json:"OutputDataConfig"`
	JobName           string           `json:"JobName"`
}

type OutputDataConfig struct {
	S3Configuration S3Configuration `json:"S3Configuration"`
}

type S3Configuration struct {
	S3Uri    string `json:"S3Uri"`
	KmsKeyId string `json:"KmsKeyId"`
}

// ExportCmd is responsible for handling FHIR data exports.
//
// Usage:
//
//	crossfhir export [flags]
//
// Flags:
//
//	-p, --pull        Pull data after export
//	-d, --dir         Directory to save exported data (default: "./fhir-data")
//	    --aws         Use AWS credentials for export
//
// Example:
//
//	crossfhir export --pull --dir ./mydirectory
func ExportCmd() *cobra.Command {
	ExportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export FHIR data from AWS Health Lake to S3 bucket",
		RunE:  Export,
	}

	ExportCmd.Flags().BoolVarP(&pull, "pull", "p", false, "Pull data after export")
	ExportCmd.Flags().StringVarP(&dir, "dir", "d", "./fhir-data", "Directory to pull exported data")
	ExportCmd.Flags().BoolVarP(&awsExport, "aws", "", false, "Export data from AWS Health Lake using AWS credentials")

	return ExportCmd
}

// Export handles the main export logic and determines whether to use
// AWS credentials or SMART on FHIR authentication based on the flags.
//
// The function supports two authentication methods:
//   - AWS credentials using static credentials
//   - SMART on FHIR using OAuth2 client credentials
func Export(cmd *cobra.Command, args []string) error {
	if awsExport {
		return ExportAws(cmd, args)
	} else {
		return ExportSmart(cmd, args)
	}
}

// ExportSmart initiates a FHIR data export using SMART on FHIR authentication and REST API.
//
// The function performs the following steps:
//  1. Validates the SMART configuration variables
//  2. Obtains an OAuth2 token
//  3. Initiates the export job
//  4. Monitors the export progress
//  5. Optionally pulls the exported data locally (if the --pull flag is set)
func ExportSmart(cmd *cobra.Command, args []string) error {
	validateSmartConfig()

	oauth2Config := &oauth2.Config{
		ClientID:     cfg.Smart.ClientID,
		ClientSecret: cfg.Smart.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: cfg.Smart.TokenURL,
		},
		Scopes: []string{cfg.Smart.Scope},
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	token, err := helpers.GetAuthToken(oauth2Config, client, cfg.Smart.GrantType)

	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	exportURL := fmt.Sprintf("%s/$export", cfg.Smart.DatastoreEndpoint)
	exportReq, err := http.NewRequest("POST", exportURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create export request: %w", err)
	}

	exportReq.Header.Set("Authorization", "Bearer "+token.AccessToken)
	exportReq.Header.Set("Accept", "application/fhir+json")
	exportReq.Header.Set("Prefer", "respond-async")

	exportBody := ExportRequestBody{
		DataAccessRoleArn: cfg.Aws.IAMExportRole,
		OutputDataConfig: OutputDataConfig{
			S3Configuration: S3Configuration{
				S3Uri:    cfg.Aws.S3Bucket,
				KmsKeyId: cfg.Aws.KmsKeyId,
			},
		},
		JobName: cfg.Aws.ExportJobName,
	}

	bodyBytes, err := json.Marshal(exportBody)
	if err != nil {
		return fmt.Errorf("failed to marshal export body: %w", err)
	}

	exportReq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	exportReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(exportReq)
	if err != nil {
		return fmt.Errorf("failed to execute export request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("export request failed with status %d: %s", resp.StatusCode, string(body))
	}

	jobLocation := resp.Header.Get("Content-Location")
	cfg.Smart.ExportJobId = extractJobId(jobLocation)

	fmt.Printf("Export job ID: %s\n", cfg.Smart.ExportJobId)

	err = monitorExportProgress(client, token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to monitor export progress: %w", err)
	}

	if pull {
		log.Println("Pulling FHIR data")
		err = PullFhirData()
		if err != nil {
			return err
		}
	}

	return nil
}

// ExportAws initiates a FHIR data export using AWS credentials.
//
// The function performs the following steps:
//  1. Validates the AWS configuration variables
//  2. Creates an export job using the AWS HealthLake SDK
//  3. Monitors the export progress
//  5. Optionally pulls the exported data locally (if the --pull flag is set)
func ExportAws(cmd *cobra.Command, args []string) error {
	validateAwsConfig()

	input := &healthlake.StartFHIRExportJobInput{
		DataAccessRoleArn: aws.String(cfg.Aws.IAMExportRole),
		DatastoreId:       aws.String(cfg.Aws.DatastoreId),
		OutputDataConfig: &types.OutputDataConfigMemberS3Configuration{
			Value: types.S3Configuration{
				S3Uri:    aws.String(cfg.Aws.S3Bucket),
				KmsKeyId: aws.String(cfg.Aws.KmsKeyId),
			},
		},
		JobName: aws.String(cfg.Aws.ExportJobName),
	}

	out, err := healthlakeClient.StartFHIRExportJob(context.TODO(), input)

	if err != nil {
		return err
	}

	cfg.Aws.ExportJobId = *out.JobId
	cfg.Aws.ExportJobStatus = string(out.JobStatus)

	err = monitorExportProgressAws()

	if pull {
		log.Println("Pulling FHIR data")
		err = PullFhirData()
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	return nil
}

// HELPER FUNCTIONS //

func extractJobId(locationUrl string) string {
	if locationUrl == "" {
		return ""
	}

	// Split URL by '/' and get the last part which should be the job ID
	parts := strings.Split(locationUrl, "/")
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}

func monitorExportProgress(client *http.Client, accessToken string) error {
	statusURL := fmt.Sprintf("%s/export/%s", cfg.Smart.DatastoreEndpoint, cfg.Smart.ExportJobId)

	for {
		req, err := http.NewRequest("GET", statusURL, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			return err
		}
		resp.Body.Close()

		var status struct {
			ExportJobProperties struct {
				JobStatus string `json:"jobStatus"`
			} `json:"exportJobProperties"`
		}

		if err := json.Unmarshal(body, &status); err != nil {
			return err
		}

		log.Printf("Export job status: %s", status.ExportJobProperties.JobStatus)

		if status.ExportJobProperties.JobStatus == "COMPLETED" {
			break
		} else if status.ExportJobProperties.JobStatus == "FAILED" {
			return fmt.Errorf("export job failed")
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

func monitorExportProgressAws() error {
	input := &healthlake.DescribeFHIRExportJobInput{
		JobId:       aws.String(cfg.Aws.ExportJobId),
		DatastoreId: aws.String(cfg.Aws.DatastoreId),
	}

	for cfg.Aws.ExportJobStatus != "COMPLETED" {
		out, err := healthlakeClient.DescribeFHIRExportJob(context.TODO(), input)

		if err != nil {
			return err
		}

		if cfg.Aws.ExportJobS3Output == "" {
			conf := out.ExportJobProperties.OutputDataConfig
			cfg.Aws.ExportJobS3Output = *conf.(*types.OutputDataConfigMemberS3Configuration).Value.S3Uri
			log.Printf("Job S3 Output: %v\n", cfg.Aws.ExportJobS3Output)
		}

		cfg.Aws.ExportJobStatus = string(out.ExportJobProperties.JobStatus)

		if cfg.Aws.ExportJobStatus == "FAILED" {
			log.Fatalf("Job - %v - Failed\n", cfg.Aws.ExportJobName)
		}

		log.Printf("Job Status: %v\n", cfg.Aws.ExportJobStatus)
		time.Sleep(5 * time.Second)
	}

	log.Printf("Job - %v - Completed\n", cfg.Aws.ExportJobName)

	return nil
}

// VALIDATION FUNCTIONS //

// Required configuration values:
//   - ClientID
//   - ClientSecret
//   - TokenURL
//   - DatastoreEndpoint
func validateSmartConfig() {
	missingEnvs := []string{}

	if cfg.Smart.ClientID == "" {
		missingEnvs = append(missingEnvs, "smart_client_id")
	}
	if cfg.Smart.ClientSecret == "" {
		missingEnvs = append(missingEnvs, "smart_client_secret")
	}
	if cfg.Smart.TokenURL == "" {
		missingEnvs = append(missingEnvs, "smart_token_url")
	}
	if cfg.Smart.DatastoreEndpoint == "" {
		missingEnvs = append(missingEnvs, "smart_datastore_endpoint")
	}

	if len(missingEnvs) > 0 {
		log.Println("Missing required SMART config variables:")
		for _, envVar := range missingEnvs {
			fmt.Printf("%s\n", envVar)
		}
		os.Exit(1)
	}
}

// Required configuration values for AWS export:
//   - AccessKey
//   - SecretKey
//   - Region
//   - S3Bucket
//   - IAMExportRole
//   - DatastoreId
//   - KmsKeyId
//   - ExportJobName
//   - DatastoreFHIRUrl
func validateAwsConfig() {
	missingEnvs := []string{}

	if cfg.Aws.AccessKey == "" {
		missingEnvs = append(missingEnvs, "aws_access_key")
	}

	if cfg.Aws.SecretKey == "" {
		missingEnvs = append(missingEnvs, "aws_secret_key")
	}

	if cfg.Aws.Region == "" {
		missingEnvs = append(missingEnvs, "aws_region")
	}

	if cfg.Aws.S3Bucket == "" {
		missingEnvs = append(missingEnvs, "aws_s3_bucket")
	}

	if cfg.Aws.IAMExportRole == "" {
		missingEnvs = append(missingEnvs, "aws_iam_export_role")
	}

	if cfg.Aws.DatastoreId == "" {
		missingEnvs = append(missingEnvs, "aws_datastore_id")
	}

	if cfg.Aws.KmsKeyId == "" {
		missingEnvs = append(missingEnvs, "aws_kms_key_id")
	}

	if cfg.Aws.ExportJobName == "" {
		missingEnvs = append(missingEnvs, "aws_export_job_name")
	}

	if cfg.Aws.DatastoreFHIRUrl == "" {
		missingEnvs = append(missingEnvs, "aws_datastore_fhir_url")
	}

	if len(missingEnvs) > 0 {
		log.Println("Missing required config variables for export action:")
		for _, envVar := range missingEnvs {
			fmt.Printf("%s\n", envVar)
		}
		os.Exit(1)
	}
}
