package cmd

import (
	"crossfhir/internal"
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

	missingEnvs := []string{}
	missingEnvs = validatePullEnvs(missingEnvs)

	if len(missingEnvs) > 0 {
		log.Println("Missing required environment variables:")
		for _, envVar := range missingEnvs {
			log.Printf("%s\n", envVar)
		}

		os.Exit(1)
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

func PullFhirData() error {
	// s3Url might be passed automatically from `export -p` command or explicitly from `pull` command
	if s3Url != "" {
		cfg.AwsExportJobS3Output = s3Url
	}

	bucket, prefix := internal.ParseS3Url(cfg.AwsExportJobS3Output)
	objects, err := internal.ListPrefixObjects(s3Client, bucket, prefix)

	log.Printf("Downloading %d FHIR data objects from S3 to local directory %s", len(objects), dir)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	sem := make(chan struct{}, 10) // Limit to 10 concurrent goroutines

	for _, object := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}        
			defer func() { <-sem }()

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

	log.Printf("Downloaded FHIR data.")

	return nil
}

func validatePullEnvs(missingEnvs []string) []string {
	cfg.AwsS3Bucket = os.Getenv("AWS_S3_BUCKET")
	if cfg.AwsS3Bucket == "" {
		missingEnvs = append(missingEnvs, "AWS_S3_BUCKET")
	}

	return missingEnvs
}
