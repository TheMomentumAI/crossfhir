package converter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/marcboeker/go-duckdb"
)

func TestToSQL(t *testing.T) {
	// Use the predefined test directory
	testDir := "tests"
	tempDir := filepath.Join(testDir, "temp")

	// Clean temp directory before test
	if err := os.RemoveAll(tempDir); err != nil {
		t.Fatalf("Failed to clean temp directory: %v", err)
	}

	// Recreate temp directory
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after tests

	testFiles := []struct {
		name     string
		filePath string
	}{
		{
			name:     "Basic test cases",
			filePath: filepath.Join(testDir, "basic.json"),
		},
	}

	for _, tf := range testFiles {
		t.Run(tf.name, func(t *testing.T) {
			ts, err := LoadTestSuite(tf.filePath)
			if err != nil {
				t.Fatalf("Error loading test suite: %v", err)
			}

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

					t.Logf("Generated SQL:\n%s", sql)

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
						t.Errorf("\nExpected:\n%s\nGot:\n%s", string(want), got)
					}
				})
			}
		})
	}
}

// Convert sql.Rows to JSON string
func resultsToJSON(rows *sql.Rows) (string, error) {
	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("error getting columns: %w", err)
	}

	// Collect all results
	var results []map[string]interface{}
	for rows.Next() {
		// Create value holders
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		// Scan row
		if err := rows.Scan(valuePointers...); err != nil {
			return "", fmt.Errorf("error scanning row: %w", err)
		}

		// Create row map
		row := make(map[string]interface{})
		for i, col := range columns {
			if values[i] == nil {
				row[col] = nil
				continue
			}

			// Handle specific types
			switch v := values[i].(type) {
			case []byte:
				row[col] = string(v)
			case int64:
				if col == "active" { // Special handling for boolean fields
					row[col] = v == 1
				} else {
					row[col] = v
				}
			default:
				row[col] = v
			}
		}
		results = append(results, row)
	}

	// Convert to JSON string
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling results to JSON: %w", err)
	}

	return string(jsonData), nil
}
