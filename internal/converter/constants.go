package converter

import (
	"fmt"
)

var constValueFields = []string{
	"value",
	"valueBase64Binary",
	"valueBoolean",
	"valueCanonical",
	"valueCode",
	"valueDate",
	"valueDateTime",
	"valueDecimal",
	"valueId",
	"valueInstant",
	"valueInteger",
	"valueInteger64",
	"valueOid",
	"valueString",
	"valuePositiveInt",
	"valueTime",
	"valueUnsignedInt",
	"valueUri",
	"valueUrl",
	"valueUuid",
}

func ProcessConstant(constant map[string]interface{}) (Constant, error) {
	fmt.Println(constant)

	var processedConstant Constant
	processedConstant.name = constant["name"].(string)

	for _, field := range constValueFields {
		if value, exists := constant[field]; exists {
			processedConstant.value = fmt.Sprintf("%v", value)
			break
		}
	}

	fmt.Println("Processed constant: ", processedConstant)

	return processedConstant, nil
}
