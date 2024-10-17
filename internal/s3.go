package internal

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func ListPrefixObjects(client *s3.Client, bucket string, prefix string) ([]types.Object, error) {
	var objects []types.Object

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		objects = append(objects, page.Contents...)
	}

	return objects, nil
}

func DownloadS3Object(client *s3.Client, bucket string, key string, destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", destDir, err)
		}
	}

	filePath := filepath.Join(destDir, filepath.Base(key))
	fmt.Println(filePath)
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer outFile.Close()

	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s to %s\n", key, filePath)
	return nil
}
