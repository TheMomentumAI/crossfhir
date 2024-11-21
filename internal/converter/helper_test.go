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
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FHIRVersion []string   `json:"fhirVersion"`
	Resources   []Resource `json:"resources"`
	Tests       []Test     `json:"tests"`
}

// Resource represents a FHIR resource used in tests
type Resource struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Name         []Name `json:"name,omitempty"`
	Active       bool   `json:"active,omitempty"`
}

// Name represents the name structure in a FHIR resource
type Name struct {
	Family string `json:"family,omitempty"`
}

// Test represents a single test case
type Test struct {
	Title         string         `json:"title"`
	Tags          []string       `json:"tags"`
	View          ViewDefinition `json:"view"`
	Expect        []Expect       `json:"expect"`
	ExpectColumns []string       `json:"expectColumns,omitempty"`
}

// Expect represents the expected output for a test
type Expect struct {
	ID       string `json:"id,omitempty"`
	Active   *bool  `json:"active,omitempty"`
	LastName string `json:"last_name,omitempty"`
	CID      string `json:"c_id,omitempty"`
	SID      string `json:"s_id,omitempty"`
	A        string `json:"a,omitempty"`
	B        string `json:"b,omitempty"`
	C        string `json:"c,omitempty"`
	D        string `json:"d,omitempty"`
	E        string `json:"e,omitempty"`
	F        string `json:"f,omitempty"`
	G        string `json:"g,omitempty"`
	H        string `json:"h,omitempty"`
}

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

func LoadDuckDB(resources []Resource, tempDir string, testName string) (*sql.DB, error) {
	// Create in-memory DuckDB instance
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return nil, fmt.Errorf("error opening DuckDB: %w", err)
	}

	// Write resources to a temporary NDJSON file
	resourceFile := filepath.Join(tempDir, testName+".ndjson")
	f, err := os.Create(resourceFile)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating temp file: %w", err)
	}
	defer f.Close()

	// Write each resource as a separate JSON line
	encoder := json.NewEncoder(f)
	for _, resource := range resources {
		if err := encoder.Encode(resource); err != nil {
			db.Close()
			return nil, fmt.Errorf("error writing resource: %w", err)
		}
	}

	// Create a view for easy querying
	_, err = db.Exec(fmt.Sprintf(`
		CREATE VIEW resource AS
		SELECT * FROM read_json_auto('%s');
	`, resourceFile))

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating resource view: %w", err)
	}

	return db, nil
}

// Helper function to compare test results
func CompareResults(got, want []map[string]interface{}) error {
	if len(got) != len(want) {
		return fmt.Errorf("result length mismatch: got %d, want %d", len(got), len(want))
	}

	// Compare each result
	for i := range got {
		if !reflect.DeepEqual(got[i], want[i]) {
			return fmt.Errorf("result mismatch at index %d:\ngot: %v\nwant: %v", i, got[i], want[i])
		}
	}

	return nil
}
