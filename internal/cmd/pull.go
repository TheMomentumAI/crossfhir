package cmd

import (
	"crossfhir/internal"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	s3Url string
)

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
	validatePullConfig()
	PullFhirData()

	return nil
}

func PullFhirData() error {
	// s3Url might be passed automatically from `export -p` command or explicitly from `pull` command
	if s3Url != "" {
		cfg.Aws.ExportJobS3Output = s3Url
	}

	bucket, prefix := internal.ParseS3Url(cfg.Aws.ExportJobS3Output)
	objects, err := internal.ListPrefixObjects(s3Client, bucket, prefix)

	log.Printf("Downloading %d FHIR data objects from S3 to local directory %s", len(objects), dir)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	sem := make(chan struct{}, 10) // Limit to 10 concurrent goroutines
	fileCounter := 0

	for _, object := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			fileCounter++
			if err := internal.DownloadS3Object(s3Client, bucket, *object.Key, dir); err != nil {
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
