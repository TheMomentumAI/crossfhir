package converter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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

func resultsToJSON(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("failed to get columns: %w", err)
	}

	results := make([]map[string]interface{}, 0)
	types, err := rows.ColumnTypes()

	// // print types as strings
	// for _, t := range types {
	// 	fmt.Println(t.DatabaseTypeName())
	// }

	if err != nil {
		return "", fmt.Errorf("failed to get column types: %w", err)
	}

	for rows.Next() {
		row, err := scanRow(rows, columns, types)
		if err != nil {
			return "", err
		}
		results = append(results, row)
	}

	if err = rows.Err(); err != nil {
		return "", fmt.Errorf("error iterating rows: %w", err)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(jsonData), nil
}

func scanRow(rows *sql.Rows, columns []string, types []*sql.ColumnType) (map[string]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	if err := rows.Scan(values...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	row := make(map[string]interface{}, len(columns))
	for i, col := range columns {
		val := reflect.ValueOf(values[i]).Elem().Interface()
		if val == nil {
			row[col] = nil
			continue
		}

		switch v := val.(type) {
		case []byte:
			// Handle text fields
			row[col] = string(v)
		case int64:
			row[col] = v
		case float64:
			row[col] = v
		case string:
			if v == "true" {
				row[col] = v == "true"
			} else if v == "false" {
				row[col] = v == "true"
			} else {
				row[col] = v
			}
		default:
			row[col] = v
		}
	}

	return row, nil
}
