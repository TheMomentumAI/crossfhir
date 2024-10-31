package cmd

import (
	"context"
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
	validateExportEnvs()

	input := &healthlake.StartFHIRExportJobInput{
		DataAccessRoleArn: aws.String(cfg.AwsIAMExportRole),
		DatastoreId:       aws.String(cfg.AwsDatastoreId),
		OutputDataConfig: &types.OutputDataConfigMemberS3Configuration{
			Value: types.S3Configuration{
				S3Uri:    aws.String(cfg.AwsS3Bucket),
				KmsKeyId: aws.String(cfg.AwsKmsKeyId),
			},
		},
		JobName: aws.String(cfg.AwsExportJobName),
	}

	out, err := healthlakeClient.StartFHIRExportJob(context.TODO(), input)

	if err != nil {
		return err
	}

	cfg.AwsExportJobId = *out.JobId
	cfg.AwsExportJobStatus = string(out.JobStatus)

	err = DescribeProgress()

	if pull {
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
		JobId:       aws.String(cfg.AwsExportJobId),
		DatastoreId: aws.String(cfg.AwsDatastoreId),
	}

	for cfg.AwsExportJobStatus != "COMPLETED" {
		out, err := healthlakeClient.DescribeFHIRExportJob(context.TODO(), input)

		if err != nil {
			return err
		}

		if cfg.AwsExportJobS3Output == "" {
			conf := out.ExportJobProperties.OutputDataConfig
			cfg.AwsExportJobS3Output = *conf.(*types.OutputDataConfigMemberS3Configuration).Value.S3Uri
			log.Printf("Job S3 Output: %v\n", cfg.AwsExportJobS3Output)
		}

		cfg.AwsExportJobStatus = string(out.ExportJobProperties.JobStatus)

		if cfg.AwsExportJobStatus == "FAILED" {
			log.Fatalf("Job - %v - Failed\n", cfg.AwsExportJobName)
		}

		log.Printf("Job Status: %v\n", cfg.AwsExportJobStatus)
		time.Sleep(5 * time.Second)
	}

	log.Printf("Job - %v - Completed\n", cfg.AwsExportJobName)

	if pull {
		log.Println("Pulling FHIR data")

		err := PullFhirData()
		if err != nil {
			return err
		}
	}

	return nil
}

func validateExportEnvs() {
	missingEnvs := []string{}

	cfg.AwsS3Bucket = os.Getenv("AWS_S3_BUCKET")
	if cfg.AwsS3Bucket == "" {
		missingEnvs = append(missingEnvs, "AWS_S3_BUCKET")
	}

	cfg.AwsIAMExportRole = os.Getenv("AWS_IAM_EXPORT_ROLE")
	if cfg.AwsIAMExportRole == "" {
		missingEnvs = append(missingEnvs, "AWS_IAM_EXPORT_ROLE")
	}

	cfg.AwsDatastoreId = os.Getenv("AWS_DATASTORE_ID")
	if cfg.AwsDatastoreId == "" {
		missingEnvs = append(missingEnvs, "AWS_DATASTORE_ID")
	}

	cfg.AwsKmsKeyId = os.Getenv("AWS_KMS_KEY_ID_ARN")
	if cfg.AwsKmsKeyId == "" {
		missingEnvs = append(missingEnvs, "AWS_KMS_KEY_ID_ARN")
	}

	cfg.AwsExportJobName = os.Getenv("AWS_EXPORT_JOB_NAME")
	if cfg.AwsExportJobName == "" {
		missingEnvs = append(missingEnvs, "AWS_EXPORT_JOB_NAME")
	}

	cfg.AwsDatastoreFHIRUrl = os.Getenv("AWS_DATASTORE_FHIR_URL")

	if len(missingEnvs) > 0 {
		log.Println("Missing required environment variables:")
		for _, envVar := range missingEnvs {
			log.Printf("%s\n", envVar)
		}

		os.Exit(1)
	}
}
