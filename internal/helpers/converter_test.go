package helpers

import (
	"testing"
)

func TestToSql(t *testing.T) {
	tests := []struct {
		name string
		vd   ViewDefinition
		want string
	}{
		{
			name: "basic patient demographics",
			vd: ViewDefinition{
				Resource: "Patient",
				Name:     "patient_demographics",
				Select: []SelectStruct{{
					Column: []Column{
						{Name: "patient_id", Path: "getResourceKey()"},
						{Name: "gender", Path: "gender"},
						{Name: "dob", Path: "birthDate"},
					},
				}},
			},
			want: `SELECT
	id AS patient_id,
	resource->>'gender' AS gender,
	resource->>'birthDate' AS dob
	FROM patient;`,
		},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ToSQL(tt.vd)
			if err != nil {
				t.Errorf("ToSQL() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("ToSQL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
