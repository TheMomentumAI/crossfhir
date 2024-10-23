package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"encoding/json"
	"io"
	"log"
	"strings"
)

func ConvertCmd() *cobra.Command {
	ConvertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert FHIR data to Postgres data",
		RunE:  Convert,
	}

	return ConvertCmd
}

func Convert(cmd *cobra.Command, args []string) error {
	dataPath := "data"

	files, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := dataPath + "/" + file.Name()
			processFile(filePath, file.Name())
		}
	}

	return nil
}

func processFile(filePath string, fileName string) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	d := json.NewDecoder(strings.NewReader(string(f)))

	for {
		var v interface{}
		err := d.Decode(&v)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		processJSONObject(v, fileName)
	}
}

func processJSONObject(v interface{}, fileName string) {
	var resourceMap map[string]interface{}
	var txid string

	parts := strings.Split(fileName, "-")

	if len(parts) >= 2 {
		txid = parts[1]
	} else {
		log.Fatal("Invalid file name format - txid not found")
	}

	resourceJSON, err := json.Marshal(v)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resourceJSON, &resourceMap)
	if err != nil {
		log.Fatal(err)
	}

	tableName := resourceMap["resourceType"].(string)

	query := fmt.Sprintf(`
			INSERT INTO %s (id, txid, status, resource)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO UPDATE SET
				txid = EXCLUDED.txid,
				status = EXCLUDED.status,
				resource = EXCLUDED.resource
	`, tableName)

	_, err = pgConn.Exec(context.Background(), query, resourceMap["id"], txid, "created", resourceJSON)
	if err != nil {
		log.Fatal(err)
	}
	// var res string
	// pgConn.QueryRow(context.Background(), "select resource from patient limit 1").Scan(&res)
	// internal.PrintJSON(res)

}
