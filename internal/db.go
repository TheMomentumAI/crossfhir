package internal

import (
	"context"
	"fmt"
	"log"

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
