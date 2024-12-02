package converter

// type ViewDefinition struct {
// 	ResourceType string         `json:"resourceType"`
// 	Resource     string         `json:"resource"`
// 	Name         string         `json:"name"`
// 	Status       string         `json:"status"`
// 	Description  string         `json:"description"`
// 	Select       []SelectStruct `json:"select"`
// 	Where        []WhereStruct  `json:"where,omitempty"`
// 	Constant     []Constant     `json:"constant,omitempty"`
// }

type FHIRDefinition struct {
	Type             string   `json:"t"`
	IsArray          bool     `json:"a,omitempty"`
	ReferenceTargets []string `json:"rt,omitempty"`
	ContentReference string   `json:"cr,omitempty"`
}

type Constant struct {
	name  string
	value string
}

type TableGroup struct {
	LeftJoin bool    `json:"leftJoin"`
	Tables   []Table `json:"tables"`
}

type Table struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Parent   *TableRef `json:"parent,omitempty"`
	Criteria []any     `json:"criteria,omitempty"`
	FHIRType string    `json:"fhirType,omitempty"`
	Function string    `json:"fn,omitempty"`
}

type TableRef struct {
	ID     string   `json:"id"`
	Table  string   `json:"t"`
	Fields []string `json:"f"`
}

// Field represents a selected field in the query
type Field struct {
	Table     string   `json:"t"`
	Fields    []string `json:"f"`
	FHIRType  string   `json:"fhirType,omitempty"`
	Function  string   `json:"fnName,omitempty"`
	FuncValue string   `json:"fnValue,omitempty"`
	Alias     string   `json:"alias,omitempty"`
	TypePath  []string `json:"fhirTypePath"`
}

func NewFlatView(definitions map[string]FHIRDefinition) *FlatView {
	return &FlatView{
		definitions: definitions,
		tables:      make([]TableGroup, 0),
		fields:      make([]Field, 0),
		nextTableID: 0,
	}
}
