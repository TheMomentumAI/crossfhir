package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

var (
	pgConn *pgx.Conn
)

func ConvertCmd() *cobra.Command {
	ConvertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert FHIR data to Postgres data. Default PostgreSQL connection",
		RunE:  Convert,
	}

	defer pgConn.Close(context.Background())

	validateConvertEnvs()
	initDbConnection()

	return ConvertCmd
}

func Convert(cmd *cobra.Command, args []string) error {

	pgConn.Exec(context.Background(), "select 1")

	// dataPath := "data"

	// files, err := os.ReadDir(dataPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	if !file.IsDir() {
	// 		filePath := dataPath + "/" + file.Name()
	// 		processFile(filePath, file.Name())
	// 	}
	// }

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

func initDbConnection() error {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbDatabase,
	)

	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	pgConn = conn
	return nil
}

func validateConvertEnvs() {
	cfg.DbHost = os.Getenv("DB_HOST")
	if cfg.DbHost == "" {
		cfg.DbHost = "localhost"
	}

	cfg.DbPort = os.Getenv("DB_PORT")
	if cfg.DbHost == "" {
		cfg.DbHost = "5432"
	}

	cfg.DbUsername = os.Getenv("DB_USERNAME")
	if cfg.DbUsername == "" {
		log.Fatalf("DB_USERNAME is required")
	}

	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	if cfg.DbPassword == "" {
		log.Fatalf("DB_PASSWORD is required")
	}

	cfg.DbDatabase = os.Getenv("DB_DATABASE")
	if cfg.DbDatabase == "" {
		log.Fatalf("DB_DATABASE is required")
	}
}
