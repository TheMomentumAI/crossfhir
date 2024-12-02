package converter

import (
	"fmt"
	"strings"
)

type ViewDefinition struct {
	ResourceType string `json:"resourceType"`
	Resource     string `json:"resource"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Description  string `json:"description"`
}

type FlatView struct {
	definitions map[string]FHIRDefinition
	tables      []TableGroup
	fields      []Field
	nextTableID int
}

// type sectionProcessor struct {
// 	fv        *FlatView
// 	constants []struct {
// 		name  string
// 		value string
// 	}
// }

// func (sp *sectionProcessor) processSection(section map[string]any, parent *Table) (*Table, error) {

//		if constants, ok := section["constants"].([]any); ok {
//			for _, c := range constants {
//				if constMap, ok := c.(map[string]any); ok {
//					sp.constants = append(sp.constants, struct {
//						name  string
//						value string
//				}{
//						name:  constMap["name"].(string),
//						value: constMap["value"].(string),
//				})
//			}
//		}
//	}
//
// // LoadViewDefinition loads a view definition into the FlatView
func (fv *FlatView) LoadViewDefinition(viewDefinition map[string]any) error {
	// fmt.Println(viewDefinition)
	processor := Processor{
		fv:        fv,
		constants: make([]Constant, 0),
	}

	err := processor.processVdSections(viewDefinition)
	if err != nil {
		return fmt.Errorf("processing view definition: %w", err)
	}

	return nil
}

// private

type Processor struct {
	fv        *FlatView
	constants []Constant
}

func (p *Processor) applyConst(expr string) string {
	result := expr

	for _, c := range p.constants {
		result = strings.ReplaceAll(result, "%"+c.name, c.value) // expr = expr.replaceAll("%" + c.name, c.value);
	}

	return result
}

func (p *Processor) processVdSections(vd map[string]any) error {

	// fmt.Println(vd["resource"])
	// if (section.constants) constants = section.constants;
	if vd["constant"] != nil {
		for _, c := range vd["constant"].([]any) {
			constant, err := ProcessConstant(c.(map[string]any))
			if err != nil {
				return fmt.Errorf("processing constant: %w", err)
			}
			p.constants = append(p.constants, Constant{
				name:  constant.name,
				value: constant.value,
			})
		}
	}

	fmt.Println(p.constants)

	if vd["resource"] != nil {
		fmt.Println(vd["resource"])
	}

	// if (section.expr || section.expression || section.path)
	// 	this.addFhirPath(pos, applyConst(section.expr || section.expression || section.path), section.name);
	// if (section.forEach)
	// 	pos = this.addFhirPath(pos, applyConst(section.forEach), null, true);
	// if (section.forEachOrNull)
	// 	pos = this.addFhirPath(pos, applyConst(section.forEachOrNull), null, true, true);
	// if (section.from)
	// 	pos = this.addFhirPath(pos, applyConst(section.from), null, true);
	// if (section.where) {
	// 	const exprGroup =
	// 		section.where.map(w => w.expr || w.expression || w.path).join(" AND ");
	// 	this.addWhereCondition(pos, "where(" + applyConst(exprGroup) + ")");
	// }

	// if (section.select) section.select.forEach( item => {
	// 	processSection(item, pos, constants)
	// });
	return nil
}

// parent, expression, alias, notInSelect, forceLeftJoin
// func addFhirPath() {
// 	expr := "Patient.name.given.where(value = 'John')"

// }
