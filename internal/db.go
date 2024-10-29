package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type PgConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func InitConnection() (*pgx.Conn, error) {
	cfg := PgConnection{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "password",
		Database: "postgres",
	}

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	conn, err := pgx.Connect(context.Background(), db_url)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return conn, nil
}

func RunFhirMigration(conn *pgx.Conn) error {
	// parametrize
	schema := "fhir.json"

	var commands []string

	data, err := os.ReadFile(schema)
	if err != nil {
		log.Fatalf("failed to read schema file: %v", err)
	}

	err = json.Unmarshal(data, &commands)
	if err != nil {
		log.Fatalf("failed to unmarshal schema: %v", err)
	}

	for _, sql := range commands {
		_, err = conn.Exec(context.Background(), sql)

		if err != nil {
			log.Fatalf("failed to exec query: %v", err)
		}
	}

	return nil
}

