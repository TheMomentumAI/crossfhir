package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/marcboeker/go-duckdb"
)

func TestToSQL(t *testing.T) {
	testFileDir := "tests"

	testFiles := []struct {
		name     string
		filePath string
	}{
		{
			name:     "Basic test cases",
			filePath: filepath.Join(testFileDir, "basic.json"),
		},
	}

	tempDir := filepath.Join(testFileDir, "temp")

	// if err := os.RemoveAll(tempDir); err != nil {
	// 	t.Fatalf("Failed to clean temp directory: %v", err)
	// }

	// defer os.RemoveAll(tempDir) // Clean up after tests

	// Recreate temp directory
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	for _, tf := range testFiles {
		t.Run(tf.name, func(t *testing.T) {
			// load test suite from file
			ts, err := LoadTestSuite(tf.filePath)
			if err != nil {
				t.Fatalf("Error loading test suite: %v", err)
			}

			// init DuckDB
			db, err := LoadDuckDB(ts.Resources, tempDir, ts.Title)
			if err != nil {
				t.Fatalf("Error creating DuckDB: %v", err)
			}
			defer db.Close()

			c := NewConverter()

			for _, tc := range ts.Tests {
				t.Run(tc.Title, func(t *testing.T) {

					sql, err := c.ToSQL(tc.View)
					if err != nil {
						t.Fatalf("ToSQL() error = %v", err)
					}

					// t.Logf("Generated SQL:\n%s", sql)

					// for testing purposes overwrite file name to tempDir + ts.Title
					sql = strings.ReplaceAll(
						sql,
						fmt.Sprintf("'%s.ndjson'", strings.ToLower(tc.View.Resource)),
						fmt.Sprintf("'%s'", filepath.Join(tempDir, ts.Title+".ndjson")),
					)

					t.Logf("Modified SQL:\n%s", sql)

					// execute query
					rows, err := db.Query(sql)
					if err != nil {
						t.Fatalf("Query execution error: %v", err)
					}
					defer rows.Close()

					got, err := resultsToJSON(rows)
					if err != nil {
						t.Fatalf("Error getting results: %v", err)
					}

					want, err := json.MarshalIndent(tc.Expect, "", "  ")
					if err != nil {
						t.Fatalf("Error marshaling expected results: %v", err)
					}

					if got != string(want) {
						// t.Errorf("\nExpected:\n%s\nGot:\n%s", string(want), got)
						t.Errorf("Results do not match")
					}
				})
			}
		})
	}
}
