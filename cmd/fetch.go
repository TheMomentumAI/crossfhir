package cmd

import (
	"crossfhir/internal"
	"sync"

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

	// todo : parametrize
	bucket := "test-fhir-sandbox-synthea-bucket20241011071523475200000001"
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
