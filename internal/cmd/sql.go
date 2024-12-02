package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"crossfhir/internal/converter"

	_ "github.com/marcboeker/go-duckdb"
	"github.com/spf13/cobra"
)

func configDuckDb() *sql.DB {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		log.Fatal(err)
	}

	// Configure S3 d
	if _, err := db.Exec(`
				INSTALL httpfs;
				LOAD httpfs;
				SET s3_region='` + cfg.Aws.Region + `';
				SET s3_access_key_id='` + cfg.Aws.AccessKey + `';
				SET s3_secret_access_key='` + cfg.Aws.SecretKey + `';
			`); err != nil {
		log.Fatal(err)
	}

	return db
}

func RunQuery(cmd *cobra.Command, args []string) {
	db := configDuckDb()
	defer db.Close()

	patientPath := cfg.Aws.S3Bucket + "802dff948fe4995fe26c029de23e1e0d-FHIR_EXPORT-12dbacc618c59a51a725ec07d8cb65d8/" + "Patient/*.ndjson"

	// Count patients
	var count int
	err := db.QueryRow(`
				SELECT COUNT(*)
				FROM read_json_auto($1,
					format='newline_delimited',
					ignore_errors=true
				);
			`, patientPath).Scan(&count)

	// err := db.QueryRow(query).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total patients: %d\n", count)
}

func Test(cmd *cobra.Command, args []string) {
	converter.Convert()
}

func SqlCmd() *cobra.Command {
	SqlCmd := &cobra.Command{
		Use:   "sql",
		Short: "SQL commands",
		Long:  `SQL commands`,
		Run:   Test,
	}

	return SqlCmd
}
