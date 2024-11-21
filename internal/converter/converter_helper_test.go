package converter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// TestSuite represents the complete test suite structure
type TestSuite struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	FHIRVersion []string      `json:"fhirVersion"`
	Resources   []interface{} `json:"resources"`
	Tests       []Test        `json:"tests"`
}

// Test represents a single test case
type Test struct {
	Title         string         `json:"title"`
	Tags          []string       `json:"tags"`
	View          ViewDefinition `json:"view"`
	Expect        []interface{}  `json:"expect"`
	ExpectColumns []string       `json:"expectColumns,omitempty"`
}

// // Expect represents the expected output for a test
// type Expect struct {
// 	ID       string `json:"id,omitempty"`
// 	Active   *bool  `json:"active,omitempty"`
// 	LastName string `json:"last_name,omitempty"`
// 	CID      string `json:"c_id,omitempty"`
// 	SID      string `json:"s_id,omitempty"`
// 	A        string `json:"a,omitempty"`
// 	B        string `json:"b,omitempty"`
// 	C        string `json:"c,omitempty"`
// 	D        string `json:"d,omitempty"`
// 	E        string `json:"e,omitempty"`
// 	F        string `json:"f,omitempty"`
// 	G        string `json:"g,omitempty"`
// 	H        string `json:"h,omitempty"`
// }

func LoadTestSuite(filename string) (*TestSuite, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading test file: %w", err)
	}

	var ts TestSuite

	err = json.Unmarshal(data, &ts)
	if err != nil {
		return nil, fmt.Errorf("error parsing test file: %w", err)
	}

	return &ts, nil
}

func LoadDuckDB(resources []interface{}, tempDir string, testName string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return nil, fmt.Errorf("error opening DuckDB: %w", err)
	}

	resourceFile := filepath.Join(tempDir, testName+".ndjson")
	if err := os.MkdirAll(filepath.Dir(resourceFile), 0755); err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating directories: %w", err)
	}

	f, err := os.Create(resourceFile)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating file %s: %w", resourceFile, err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	for _, resource := range resources {
		if err := encoder.Encode(resource); err != nil {
			db.Close()
			return nil, fmt.Errorf("error encoding resource: %w", err)
		}
	}

	return db, nil
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
