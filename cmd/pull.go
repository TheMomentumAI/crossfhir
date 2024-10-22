package cmd

import (
	"crossfhir/internal"
	"log"
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

	PullCmd.Flags().StringVarP(&s3Url, "url", "u", "", "Directory to save FHIR exported data")
	PullCmd.Flags().StringVarP(&dir, "dir", "d", "./fhir-data", "Directory to save FHIR exported data")

	return PullCmd
}

func Pull(cmd *cobra.Command, args []string) error {
	PullFhirData()

	return nil
}

func PullFhirData() error {
	if s3Url != "" {
		cfg.AwsExportJobS3Output = s3Url
	}

	outputS3Url := cfg.AwsExportJobS3Output

	bucket, prefix := internal.ParseS3Url(outputS3Url)
	objects, err := internal.ListPrefixObjects(s3Client, bucket, prefix)

	log.Printf("Downloading %d FHIR data objects from S3 to local directory %s", len(objects), dir)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	for _, object := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
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
