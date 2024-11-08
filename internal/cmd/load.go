package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"crossfhir/internal"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

var (
	pgConn   *pgx.Conn
	dataPath string
	migrate  bool
)

func LoadCmd() *cobra.Command {
	LoadCmd := &cobra.Command{
		Use:   "load",
		Short: "Load pulled FHIR data to PostgreSQL.",
		RunE:  Load,
	}

	LoadCmd.Flags().StringVarP(&dataPath, "data", "d", "", "Path to FHIR data")
	LoadCmd.Flags().BoolVarP(&migrate, "migrate", "m", false, "Run database migration that prepares the database for FHIR data")
	LoadCmd.MarkFlagRequired("data")

	return LoadCmd
}

func Load(cmd *cobra.Command, args []string) error {
	validateLoadConfig()
	initDbConnection()

	if migrate {
		RunFhirMigration(pgConn)
	}

	files, err := os.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := dataPath + "/" + file.Name()
			loadFileData(filePath, file.Name())
		}
	}

	log.Println("Data loaded successfully")

	defer pgConn.Close(context.Background())

	return nil
}

func loadFileData(filePath string, fileName string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var jsonObject interface{}
		line := scanner.Text()

		if err := json.Unmarshal([]byte(line), &jsonObject); err != nil {
			log.Fatal(err)
		}

		loadObject(jsonObject, fileName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func loadObject(v interface{}, fileName string) {
	var resourceMap map[string]interface{}

	parts := strings.Split(fileName, "-")
	if len(parts) < 2 {
		log.Fatal("Invalid file name format - txid not found")
	}

	txid := parts[1]

	resourceJSON, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(resourceJSON, &resourceMap); err != nil {
		log.Fatal(err)
	}

	tableName := strings.ToLower(resourceMap["resourceType"].(string))

	// only allowed table names here
	// sanitize
	// think about versioning
	query := fmt.Sprintf(`
			INSERT INTO %s (id, txid, resource)
			VALUES ($1, $2, $3)
			ON CONFLICT (id) DO UPDATE SET
				txid = EXCLUDED.txid,
				resource = EXCLUDED.resource
	`, tableName)

	// when verbose option is enabled, print the SQL command

	_, err = pgConn.Exec(context.Background(), query, resourceMap["id"], txid, resourceJSON)
	if err != nil {
		log.Fatal(err)
	}
}

func initDbConnection() error {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Database,
	)

	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	pgConn = conn
	return nil
}

func validateLoadConfig() {
	missingEnvs := []string{}

	if cfg.Db.Host == "" {
		cfg.Db.Host = "localhost"
	}

	if cfg.Db.Port == "" {
		cfg.Db.Port = "5432"
	}

	if cfg.Db.Username == "" {
		missingEnvs = append(missingEnvs, "db_username")
	}

	if cfg.Db.Password == "" {
		missingEnvs = append(missingEnvs, "db_password")
	}

	if cfg.Db.Database == "" {
		missingEnvs = append(missingEnvs, "db_database")
	}

	if len(missingEnvs) > 0 {
		log.Println("Missing required config variables for load action:")
		for _, envVar := range missingEnvs {
			fmt.Printf("%s\n", envVar)
		}
		os.Exit(1)
	}
}

func RunFhirMigration(conn *pgx.Conn) error {
	for _, sql := range internal.FhirSQLCommands {
		// when verbose option is enabled, print the SQL command

		_, err := conn.Exec(context.Background(), sql)

		if err != nil {
			log.Fatalf("failed to exec query: %v", err)
		}
	}

	log.Println("FHIR migration completed")

	return nil
}
