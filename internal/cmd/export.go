package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/spf13/cobra"
)

var (
	pull bool
)

func ExportCmd() *cobra.Command {
	ExportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export FHIR data from AWS Health Lake to S3 bucket",
		RunE:  Export,
	}

	ExportCmd.Flags().BoolVarP(&pull, "pull", "p", false, "Pull data after export")
	ExportCmd.Flags().StringVarP(&dir, "dir", "d", "./fhir-data", "Directory to pull exported data")
	return ExportCmd
}

func Export(cmd *cobra.Command, args []string) error {
	validateExportConfig()

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

	err = DescribeProgress()

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

func DescribeProgress() error {
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

func validateExportConfig() {
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
