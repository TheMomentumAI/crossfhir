package cmd

import (
	"crossfhir/internal"
	"sync"

	"github.com/spf13/cobra"
)

func PullCmd() *cobra.Command {
	PullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull FHIR data from S3 to local",
		RunE:  Pull,
	}

	return PullCmd
}

func Pull(cmd *cobra.Command, args []string) error {

	PullFhirData()

	return nil
}

func PullFhirData() error {
	// todo : parametrize
	bucket := cfg.AwsS3Bucket

	// how to get this from sdk
	prefix := "8699accb152044514abe6bcc49744168-FHIR_EXPORT-28ee898ca886e56a0a28ef73ad75c370/"
	localPath := "./data"

	objects, err := internal.ListPrefixObjects(s3Client, bucket, prefix)
	if err != nil {
		return err
	}

	// check number of goroutines
	var wg sync.WaitGroup
	errChan := make(chan error, len(objects))

	for _, object := range objects {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := internal.DownloadS3Object(s3Client, bucket, *object.Key, localPath); err != nil {
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

	return nil
}
