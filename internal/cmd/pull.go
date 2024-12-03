package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"crossfhir/internal/helpers"

	"github.com/spf13/cobra"
)

var (
	s3Url string
)


// PullCmd is responsible for handling FHIR data downloads from S3.
//
// The command supports concurrent downloads with a default limit of 10 concurrent operations.
// It automatically creates the destination directory if it doesn't exist and handles
// both individual file downloads and bulk transfers.
//
// Usage:
//
//	crossfhir pull [flags]
//
// Flags:
//
//	-u, --url        URL of the S3 bucket to pull FHIR data from (required)
//	-d, --dir        Directory to save FHIR data (default: "./fhir-data")
func PullCmd() *cobra.Command {
	PullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull FHIR data from S3 to local",
		RunE:  Pull,
	}

	PullCmd.Flags().StringVarP(&s3Url, "url", "u", "", "URL of the S3 bucket to pull FHIR data from")
	PullCmd.Flags().StringVarP(&dir, "dir", "d", "./fhir-data", "Directory to save FHIR exported data")

	PullCmd.MarkFlagRequired("url")

	return PullCmd
}

func Pull(cmd *cobra.Command, args []string) error {
	PullFhirData()

	return nil
}

// PullFhirData downloads FHIR data from an S3 bucket to local storage.
//
// The function performs the following steps:
//  1. Validates the pull configuration
//  2. Lists all objects in the specified S3 location
//  3. Downloads files concurrently with rate limiting
//  4. Creates the local directory structure as needed
//
// The function uses a worker pool pattern to handle concurrent downloads
// while maintaining a maximum of 10 concurrent operations to prevent
// overwhelming system resources.
// This function can be also triggered automatically by the `export -p` command.
// In this case, the S3 URL is passed automatically from the export command.
// While running the `pull` command, the S3 URL must be passed explicitly.
func PullFhirData() error {
	validatePullConfig()

	if s3Url != "" {
		cfg.Aws.ExportJobS3Output = s3Url
	}

	bucket, prefix := helpers.ParseS3Url(cfg.Aws.ExportJobS3Output)
	objects, err := helpers.ListPrefixObjects(s3Client, bucket, prefix)

	log.Printf("Downloading %d FHIR data objects from S3 to local directory %s", len(objects), dir)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	// Limit to 10 concurrent downloads
	sem := make(chan struct{}, 10)
	fileCounter := 0

	for _, object := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			fileCounter++
			if err := helpers.DownloadS3Object(s3Client, bucket, *object.Key, dir); err != nil {
				errChan <- err
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	log.Printf("Downloaded %d FHIR data files.", fileCounter)
	return nil
}

// Required configuration values for AWS S3 pull action:
//   - AWS Access Key
//   - AWS Secret Key
//   - AWS Region
//   - S3 Bucket
func validatePullConfig() {
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

	if len(missingEnvs) > 0 {
		log.Println("Missing required config variables for pull action:")
		for _, envVar := range missingEnvs {
			fmt.Printf("%s\n", envVar)
		}
		os.Exit(1)
	}
}
