package cmd

import (
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
	// for each file in data directory read file, convert to json

	f, err := os.ReadFile("tmp/test.ndjson")
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

		outputFile, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer outputFile.Close()

		encoder := json.NewEncoder(outputFile)
		if err := encoder.Encode(v); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
