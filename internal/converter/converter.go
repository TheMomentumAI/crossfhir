package converter

import (
	"fmt"
	"strings"
)

func NewConverter() *Converter {
	return &Converter{
		indentation: 2,
	}
}

// ToSQL converts a ViewDefinition to a SQL query string
func (c *Converter) ToSQL(vd ViewDefinition) (string, error) {
	if err := c.validate(vd); err != nil {
		return "", err
	}

	var queryParts []string

	// Process SELECT clause
	selectClause, err := c.buildSelectClause(vd.Select)
	if err != nil {
		return "", fmt.Errorf("error building SELECT clause: %w", err)
	}
	queryParts = append(queryParts, selectClause)

	// FROM clause
	queryParts = append(queryParts, c.buildFromClause(vd.Resource))

	// WHERE clause
	if len(vd.Where) > 0 {
		whereClause, err := c.buildWhereClause(vd.Where)
		if err != nil {
			return "", fmt.Errorf("error building WHERE clause: %w", err)
		}
		queryParts = append(queryParts, whereClause)
	}

	return strings.Join(queryParts, "\n") + ";", nil
}

func (c *Converter) buildSelectClause(selects []SelectStruct) (string, error) {
	if len(selects) == 0 {
		return "", fmt.Errorf("at least one select statement is required")
	}

	parts := []string{"SELECT"}
	columns := []string{}

	for _, sel := range selects {
		cols, err := c.processColumns(sel.Column)
		if err != nil {
			return "", err
		}
		columns = append(columns, cols...)
	}

	parts = append(parts, c.indent(strings.Join(columns, ",\n"+c.indent(""))))
	return strings.Join(parts, "\n"), nil
}

func (c *Converter) processColumns(columns []Column) ([]string, error) {
	var sqlColumns []string

	for _, col := range columns {

		//????
		if col.Path == "" {
			continue
		}

		sql, err := c.convertPathToSQL(col.Path)
		if err != nil {
			return nil, fmt.Errorf("error converting path '%s': %w", col.Path, err)
		}

		if col.Name != "" {
			sql += fmt.Sprintf(" AS %s", col.Name)
		}

		sqlColumns = append(sqlColumns, sql)
	}

	return sqlColumns, nil
}

func (c *Converter) buildFromClause(resource string) string {
	return fmt.Sprintf("FROM read_json_auto('%s.ndjson') as resource", strings.ToLower(resource))
}

func (c *Converter) buildWhereClause(whereClauses []WhereStruct) (string, error) {
	if len(whereClauses) == 0 {
		return "", nil
	}

	var conditions []string
	for _, clause := range whereClauses {
		condition, err := c.convertWhereExpression(clause.Path)
		if err != nil {
			return "", err
		}
		conditions = append(conditions, condition)
	}

	return "WHERE " + c.indent(strings.Join(conditions, " AND\n"+c.indent(""))), nil
}

func (c *Converter) convertPathToSQL(path string) (string, error) {
	switch {
	case path == "getResourceKey()":
		return "resource['id']", nil
	case strings.Contains(path, "getReferenceKey"):
		parts := strings.Split(path, ".")
		referenceField := parts[0]
		return fmt.Sprintf("json_extract(resource['%s']['reference'], '$.reference')", referenceField), nil
	case strings.Contains(path, ".first()"):
		basePath := strings.TrimSuffix(path, ".first()")
		parts := strings.Split(basePath, ".")
		jsonPath := strings.Join(parts, "']['")
		return fmt.Sprintf("list_extract(resource['%s'], 1)", jsonPath), nil
	default:
		parts := strings.Split(path, ".")
		jsonPath := strings.Join(parts, "']['")
		return fmt.Sprintf("resource['%s']", jsonPath), nil
	}
}

func (c *Converter) convertWhereExpression(expr string) (string, error) {
	if strings.Contains(expr, " and ") {
		parts := strings.Split(expr, " and ")
		conditions := make([]string, len(parts))
		for i, part := range parts {
			converted, err := c.convertWhereExpression(strings.TrimSpace(part))
			if err != nil {
				return "", err
			}
			conditions[i] = converted
		}
		return strings.Join(conditions, " AND "), nil
	}

	if strings.HasSuffix(expr, ".exists()") {
		path := strings.TrimSuffix(expr, ".exists()")
		parts := strings.Split(path, ".")
		jsonPath := strings.Join(parts, "']['")
		return fmt.Sprintf("resource['%s'] IS NOT NULL", jsonPath), nil
	}

	if strings.Contains(expr, " = ") {
		parts := strings.Split(expr, " = ")
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		leftParts := strings.Split(left, ".")
		jsonPath := strings.Join(leftParts, "']['")

		if strings.HasPrefix(right, "'") && strings.HasSuffix(right, "'") {
			return fmt.Sprintf("resource['%s'] = %s", jsonPath, right), nil
		}

		if right == "true" || right == "false" {
			return fmt.Sprintf("CAST(resource['%s'] AS BOOLEAN) = %s", jsonPath, right), nil
		}

		return fmt.Sprintf("resource['%s'] = '%s'", jsonPath, right), nil
	}

	return "", fmt.Errorf("unsupported where expression: %s", expr)
}

func (c *Converter) validate(vd ViewDefinition) error {
	if vd.Resource == "" {
		return fmt.Errorf("resource is required")
	}
	if len(vd.Select) == 0 || len(vd.Select[0].Column) == 0 {
		return fmt.Errorf("at least one column in select is required")
	}
	return nil
}

func (c *Converter) indent(text string) string {
	return strings.Repeat(" ", c.indentation) + text
}
