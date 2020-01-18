package amanar

import (
	"bytes"
	"testing"
)

func TestProcessAmanarWithWriter(t *testing.T) {
	type args struct {
		githubToken string
		a           *Amanar
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{
			name: "Configuration with only constants produces expected output",
			args: args{
				githubToken: "",
				a: &Amanar{
					[]AmanarConfiguration{
						{
							Constant: &Constant{
								Template: stringPointer("First template line"),
							},
						},
						{
							Constant: &Constant{
								Template: stringPointer("Second template line"),
							},
						},
					},
				},
			},
			wantWriter: "First template lineSecond template line",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			ProcessAmanarWithWriter(tt.args.githubToken, tt.args.a, writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("ProcessAmanarWithWriter() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
