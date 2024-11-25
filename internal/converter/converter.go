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

// parse vd
// build select for each of select object
	// handle .first() .exists() .forEach() etc.
// build from cause
// build where cause


// ToSQL converts a ViewDefinition to a SQL query string
func (c *Converter) ToSQL(vd ViewDefinition) (string, error) {
	if err := c.validate(vd); err != nil {
		return "", err
	}

	var queryParts []string

	// Process SELECT clause
	selectClause, err := c.Select(vd.Select)
	if err != nil {
		return "", fmt.Errorf("error building SELECT clause: %w", err)
	}
	queryParts = append(queryParts, selectClause)

	// FROM clause
	queryParts = append(queryParts, c.buildFromClause(vd.Resource))

	// // WHERE clause
	// if len(vd.Where) > 0 {
	// 	whereClause, err := c.Where(vd.Where)
	// 	if err != nil {
	// 		return "", fmt.Errorf("error building WHERE clause: %w", err)
	// 	}
	// 	queryParts = append(queryParts, whereClause)
	// }

	return strings.Join(queryParts, "\n") + ";", nil
}

func (c *Converter) Select(selects []SelectStruct) (string, error) {
	if len(selects) == 0 {
		return "", fmt.Errorf("at least one select statement is required")
	}

	parts := []string{"SELECT"}
	columns := []string{}

	for _, sel := range selects {
		cols, err := c.processColumn(sel.Column)
		if err != nil {
			return "", err
		}
		columns = append(columns, cols...)
	}

	parts = append(parts, c.indent(strings.Join(columns, ",\n"+c.indent(""))))
	return strings.Join(parts, "\n"), nil
}

func (c *Converter) processColumn(columns []Column) ([]string, error) {
	var sqlColumns []string

	for _, col := range columns {
		// check
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

func (c *Converter) convertPathToSQL(path string) (string, error) {
	switch {
	case strings.Contains(path, ".first()"):
		basePath := strings.TrimSuffix(path, ".first()")
		parts := strings.Split(basePath, ".")

		// for "name.family.first()" return  resource->>'$.name[0].family'
		return fmt.Sprintf("resource->>'$.%s[0].%s'", parts[0], strings.Join(parts[1:], ".")), nil

		// For literal string values like 'A'
	case strings.HasPrefix(path, "'") && strings.HasSuffix(path, "'"):

		return path, nil

	default:
		elements := strings.Split(path, ".")
		return fmt.Sprintf("resource->>'$.%s'", strings.Join(elements, ".")), nil
	}
}

// FROM
func (c *Converter) buildFromClause(resource string) string {
	// todo resource from *?
	return fmt.Sprintf("FROM read_json_auto('%s.ndjson') as resource", strings.ToLower(resource))
}

// HELPERS
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
