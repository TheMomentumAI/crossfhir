package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/spf13/cobra"
)

func ExportCmd() *cobra.Command {
	ExportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export FHIR data from AWS Health Lake",
		RunE:  Export,
	}

	return ExportCmd
}

func Export(cmd *cobra.Command, args []string) error {

	// Define the input for the StartFHIRExportJobInput
	input := &healthlake.StartFHIRExportJobInput{
		// ClientToken:       aws.String("my-client-token"),
		DataAccessRoleArn: aws.String("arn:aws:iam::938862131513:role/HealthLakeImportExportRole"),
		DatastoreId:       aws.String("8699accb152044514abe6bcc49744168"),
		OutputDataConfig: &types.OutputDataConfigMemberS3Configuration{
			Value: types.S3Configuration{
				S3Uri: aws.String("s3://test-fhir-sandbox-synthea-bucket20241011071523475200000001"),
				KmsKeyId: aws.String("arn:aws:kms:us-east-1:938862131513:key/749b1e97-85db-495d-bd1e-dd05af8adaf5"),
			},
		},
		JobName: aws.String("my-export-job"),
	}

	out, err := healthlakeClient.StartFHIRExportJob(context.TODO(), input)

	if err != nil {
		return err
	}

	fmt.Printf("Export id job : %v\n", *out.JobId)


	return nil
}
