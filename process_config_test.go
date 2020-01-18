package amanar

import (
	"bytes"
	"reflect"
	"testing"
)

func TestProcessConstantConfigItem(t *testing.T) {
	type args struct {
		constant Constant
	}

	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{name: "Can render string from a Constant",
			args: args{
				constant: Constant{
					Template: stringPointer("This is a constant template."),
				},
			},
			wantWriter: "This is a constant template.",
		},
		{name: "Can render file from a Constant",
			args: args{
				constant: Constant{
					TemplatePath: stringPointer("./fixtures/constant_template.go.md"),
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

func TestNewConfigurationProcessor(t *testing.T) {
	type args struct {
		githubToken string
		ac          AmanarConfiguration
	}
	tests := []struct {
	name          string
	args          args
	wantWriter    string
	want          ConfigurationProcessor
	wantErr       bool
	wantErrString string
}{
		{
			name:       "Will return error if Github token not provided for Vault",
			args:       args{
				githubToken: "",
				ac:          AmanarConfiguration{
					VaultAddress:       stringPointer("https://vault.com"),
					VaultConfiguration: []VaultConfiguration{},
				},
			},
			wantWriter: "",
			want:       nil,
			wantErr:    true,
			wantErrString: "[GITHUB AUTH] Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			got, err := NewConfigurationProcessor(tt.args.githubToken, tt.args.ac, writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfigurationProcessor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr && tt.wantErrString != err.Error() {
				t.Errorf("NewConfigurationProcessor() error string = %s, wantErrString %s", err.Error(), tt.wantErrString)
				return
			}

			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("NewConfigurationProcessor() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigurationProcessor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
