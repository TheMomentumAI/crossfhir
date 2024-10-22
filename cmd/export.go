package cmd

import (
	"context"
	"log"
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
		Short: "Export FHIR data from AWS Health Lake",
		RunE:  Export,
	}

	ExportCmd.Flags().BoolVarP(&pull, "pull", "p", false, "Pull data after export")
	ExportCmd.Flags().StringVarP(&dir, "dir", "d", "./fhir-data", "Directory to save FHIR exported data")
	return ExportCmd
}

func Export(cmd *cobra.Command, args []string) error {
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
