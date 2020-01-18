package amanar

import (
	"bytes"
	"testing"
)

func TestProcessConstantConfigItem(t *testing.T) {
	type args struct {
		constant Constant
	}

	s := "This is a constant template."
	p := "./fixtures/constant_template.go.md"
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{name: "Can render string from a Constant",
			args: args{
				constant: Constant{
					Template: &s,
				},
			},
			wantWriter: "This is a constant template.",
		},
		{name: "Can render file from a Constant",
			args: args{
				constant: Constant{
					TemplatePath: &p,
				},
			},
			wantWriter: `File Constant
Template
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			ProcessConstantConfigItem(tt.args.constant, writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("ProcessConstantConfigItem() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
