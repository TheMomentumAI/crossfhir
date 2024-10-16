package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

func FetchCmd() *cobra.Command {
	FetchCmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch FHIR data from S3",
		RunE:  FetchData,
	}

	return FetchCmd
}

func FetchData(cmd *cobra.Command, args []string) error {

	bucket := "test-fhir-sandbox-synthea-bucket20241011071523475200000001"
	prefix := "8699accb152044514abe6bcc49744168-FHIR_EXPORT-28ee898ca886e56a0a28ef73ad75c370/"

	outputFile := "export_data"

	result, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("Error getting object from S3:", err)
	}
	defer result.Body.Close()

	// Create a file to write the S3 object data
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer outFile.Close()

	// Write the content to the file
	_, err = outFile.ReadFrom(result.Body)
	if err != nil {
		fmt.Println("Error writing object to file:", err)
	}

	fmt.Printf("Successfully downloaded %s from S3 bucket %s to local file %s\n", key, bucket, outputFile)

	return nil
}
