package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func PrintJSON(res string) {
	var prettyJSON map[string]interface{}

	err := json.Unmarshal([]byte(res), &prettyJSON)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	prettyBytes, err := json.MarshalIndent(prettyJSON, "", "  ")

	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(prettyBytes))

	if err != nil {
		log.Fatal(err)
	}
}

func ExecQuery(conn *pgx.Conn, query string) {
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(values)
	}
}
