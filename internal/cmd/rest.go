package cmd

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"crossfhir/internal/helpers"

	awsSigner "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	resourceType string
	resourceId   string
	payload			string
)

func RestCmd() *cobra.Command {
	RestCmd := &cobra.Command{
		Use:   "rest",
		Short: "Interact with FHIR REST API",
	}

	RestCmd.AddCommand(GetCmd())
	RestCmd.AddCommand(PutCmd())

	return RestCmd
}

func GetCmd() *cobra.Command {
	GetCmd := &cobra.Command{
		Use:   "get",
		Short: "Get FHIR data from AWS Health Lake",
		RunE:  Get,
	}

	GetCmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Resource type to get, e.g. Patient")
	GetCmd.Flags().StringVarP(&resourceId, "id", "i", "", "Resource ID to get, e.g. 20a70ecf-c423-4318-82c3-40542074d6a8")

	GetCmd.MarkFlagRequired("resource")

	return GetCmd
}

func PutCmd() *cobra.Command {
	PutCmd := &cobra.Command{
		Use:   "put",
		Short: "PUT request for updating resource on FHIR server",
		RunE:  Put,
	}

	PutCmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Resource type to update, e.g. Patient")
	PutCmd.Flags().StringVarP(&resourceId, "id", "i", "", "Resource ID to update, e.g. 20a70ecf-c423-4318-82c3-40542074d6a8")
	PutCmd.Flags().StringVarP(&payload, "payload", "p", "", "Resource payload to update")

	PutCmd.MarkFlagRequired("resource")
	PutCmd.MarkFlagRequired("id")
	PutCmd.MarkFlagRequired("payload")

	return PutCmd
}

func Get(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/%s/%s", cfg.Aws.DatastoreFHIRUrl, resourceType, resourceId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	signRequest(req, "", "healthlake")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to execute request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	// TODO prittify the JSON responses and colorize them
	// TODO operate on JQ style queries (?)
	if resp.StatusCode == http.StatusOK {
		color.Green("Response Status: %s", resp.Status)
	} else {
		color.Red("Response Status: %s", resp.Status)
	}
	helpers.PrintJSON(string(body))

	return nil
}

func Put(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/%s/%s", cfg.Aws.DatastoreFHIRUrl, resourceType, resourceId)

	payload := cmd.Flag("payload").Value.String()

	req, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	signRequest(req, payload, "healthlake")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	if resp.StatusCode == http.StatusOK {
		color.Green("Response Status: %s", resp.Status)
	} else {
		color.Red("Response Status: %s", resp.Status)
	}
	helpers.PrintJSON(string(body))

	return nil
}

// private
func signRequest(request *http.Request, payload string, service string) {
	credentials, err := healthlakeClient.Options().Credentials.Retrieve(context.TODO())
	if err != nil {
		log.Fatalf("failed to retrieve credentials: %v", err)
	}

	signer := awsSigner.NewSigner()

	err = signer.SignHTTP(
		context.TODO(),
		credentials,
		request,
		calculatePayloadHash(payload),
		service,
		cfg.Aws.Region,
		time.Now(),
	)

	if err != nil {
		log.Fatalf("failed to sign request: %v", err)
	}
}

// TODO: how to sign with signer package
func calculatePayloadHash(payload string) string {
	hash := sha256.New()
	hash.Write([]byte(payload))
	return hex.EncodeToString(hash.Sum(nil))
}
