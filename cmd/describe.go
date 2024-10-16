package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"

	"github.com/spf13/cobra"
)

func DescribeCmd() *cobra.Command {
	DescribeCmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe FHIR data export job",
		RunE:  DescribeExport,
	}

	return DescribeCmd
}

func DescribeExport(cmd *cobra.Command, args []string) error {

	// Define the input for the StartFHIRExportJobInput
	input := &healthlake.DescribeFHIRExportJobInput{
		JobId:       aws.String("fbefe69179cc662b675a49a190089b83"),
		DatastoreId: aws.String("8699accb152044514abe6bcc49744168"),
	}

	out, err := healthlakeClient.DescribeFHIRExportJob(context.TODO(), input)

	if err != nil {
		return err
	}
	
	fmt.Printf("Job Status: %v\n", out.ExportJobProperties.JobStatus)

	return nil
}
