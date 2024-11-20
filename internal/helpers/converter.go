package helpers

import (
	"fmt"
	"strings"
)

type ViewDefinition struct {
	Resource string         `json:"resource"`
	Name     string         `json:"name"`
	Select   []SelectStruct `json:"select"`
	Where    []WhereStruct  `json:"where,omitempty"`
}

type SelectStruct struct {
	Column []Column       `json:"column"`
	Select []SelectStruct `json:"select,omitempty"`
}

type Column struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path"`
}

type WhereStruct struct {
	Path string `json:"path"`
}

type Converter struct {
	indentation int
}

func NewConverter() *Converter {
	return &Converter{
		indentation: 2,
	}
}

func (c *Converter) ToSQL(vd ViewDefinition) (string, error) {
	if vd.Resource == "" || len(vd.Select) == 0 {
		return "", fmt.Errorf("ViewDefinition must have 'resource' and 'select' fields")
	}

	var parts []string

	// Process SELECT clause
	columns := c.processColumns(vd.Select[0].Column)
	parts = append(parts, "SELECT")
	parts = append(parts, c.indent(strings.Join(columns, ",\n"+c.indent(""))))

	// FROM clause with JSON reading
	parts = append(parts, fmt.Sprintf("FROM read_json_auto('%s.ndjson')", strings.ToLower(vd.Resource)))

	// WHERE clause if specified
	if len(vd.Where) > 0 {
		whereClauses := c.processWhere(vd.Where)
		if whereClauses != "" {
			parts = append(parts, "WHERE")
			parts = append(parts, c.indent(whereClauses))
		}
	}

	return strings.Join(parts, "\n") + ";", nil
}

func (c *Converter) processColumns(columns []Column) []string {
	var sqlColumns []string

	for _, col := range columns {
		if col.Path == "" {
			continue
		}

		var sql string

		switch {
		case col.Path == "getResourceKey()":
			sql = "id"
		case strings.Contains(col.Path, "getReferenceKey"):
			// Extract reference key from JSON reference
			pathParts := strings.Split(col.Path, ".")
			referenceField := pathParts[0]
			sql = fmt.Sprintf("json_extract(%s->>'reference', '$.reference')", referenceField)
		default:
			// JSON path extraction using DuckDB's JSON functions
			sql = fmt.Sprintf("json_extract_string(resource, '$.%s')", col.Path)
		}

		if col.Name != "" {
			sql += fmt.Sprintf(" AS %s", col.Name)
		}

		sqlColumns = append(sqlColumns, sql)
	}

	return sqlColumns
}

func (c *Converter) processWhere(whereClauses []WhereStruct) string {
	var conditions []string

	for _, clause := range whereClauses {
		if clause.Path != "" {
			// Basic WHERE clause conversion
			conditions = append(conditions, fmt.Sprintf("resource->'%s' IS NOT NULL", clause.Path))
		}
	}

	return strings.Join(conditions, " AND ")
}

func (c *Converter) indent(text string) string {
	return strings.Repeat(" ", c.indentation) + text
}
