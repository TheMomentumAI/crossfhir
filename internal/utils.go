package internal

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"encoding/json"
	"github.com/k0kubun/pp/v3"
	"fmt"
)


func PrintJSON(res string) {
	var prettyJSON map[string]interface{}

	err := json.Unmarshal([]byte(res), &prettyJSON)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	pp.Print(prettyJSON)
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
