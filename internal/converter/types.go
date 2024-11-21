package converter

// ViewDefinition represents the main structure for SQL on FHIR view definitions
type ViewDefinition struct {
	ResourceType string         `json:"resourceType"`
	Resource     string         `json:"resource"`
	Name         string         `json:"name"`
	Status       string         `json:"status"`
	Description  string         `json:"description"`
	Select       []SelectStruct `json:"select"`
	Where        []WhereStruct  `json:"where,omitempty"`
	Constant     []Constant     `json:"constant,omitempty"`
}

// Each ViewDefinition instance has a select that specifies the content and names
// for the columns in the view. The content for each column is defined with FHIRPath
// expressions that return specific data elements from the FHIR resources.
type SelectStruct struct {
	Type          string         `json:"type,omitempty"`
	Column        []Column       `json:"column,omitempty"`
	Select        []SelectStruct `json:"select,omitempty"`
	UnionAll      []SelectStruct `json:"unionAll,omitempty"`
	ForEach       string         `json:"forEach,omitempty"`
	ForEachOrNull string         `json:"forEachOrNull,omitempty"`
}

type Column struct {
	Name       string `json:"name,omitempty"`
	Path       string `json:"path"`
	Type       string `json:"type,omitempty"`
	Collection bool   `json:"collection,omitempty"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// The ViewDefinition may include one or more where clauses that may be used to
// further limit, or filter, the resources included in the view. For instance,
// users may have different views for blood pressure observations or other observation types.
type WhereStruct struct {
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
}

// A constant is a value that is injected into a FHIRPath expression through the
// use of a FHIRPath external constant with the same name.
type Constant struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Converter struct {
	indentation int
}
