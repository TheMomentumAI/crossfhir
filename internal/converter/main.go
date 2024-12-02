package converter

import (
	"encoding/json"
	"fmt"
	"os"
)

func Convert() {
	// load file from vd file
	vdFile, err := os.ReadFile("vd/test.json")
	if err != nil {
		fmt.Errorf("error reading file: %w", err)
		return
	}

	var vd any
	err = json.Unmarshal(vdFile, &vd)
	if err != nil {
		fmt.Errorf("error unmarshalling json: %w", err)
		return
	}

	// fmt.Println(vd)

	fv := NewFlatView(nil)
	fv.LoadViewDefinition(vd.(map[string]any)) // how about vd.(map[string]any

	// TODO: Add example implementation
	// fmt.Println(flatView)
}
