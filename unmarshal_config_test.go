package amanar

import (
	"reflect"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	type args struct {
		configFilepath  string
		schemaAssetPath string
	}
	tests := []struct {
		name             string
		args             args
		hasConfiguration bool
		wantErr          error
		hasErrors        bool
		errors           []string
	}{
		{
			name: "Valid JSON can be loaded",
			args: args{
				configFilepath: "./example/example_configuration.json",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
		},
		{
			name: "Valid YAML can be loaded",
			args: args{
				configFilepath: "./example/example_configuration.yml",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
		},
		{
			name: "Invalid YAML cannot be loaded: top-level property that do not begin with x-",
			args: args{
				configFilepath: "./fixtures/invalid_properties.yml",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        true,
			errors:           []string{"(root): Additional property invalid_prop is not allowed"},
		},
		{
			name: "Invalid YAML cannot be loaded: anchor resolves to invalid type",
			args: args{
				configFilepath: "./fixtures/invalid_resolved_anchor.yml",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        true,
			errors:           []string{"amanar_configuration.0.vault_configuration.0.vault_path: Invalid type. Expected: string, given: integer"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, gotErr, gotRe := LoadConfiguration(tt.args.configFilepath)
			if (gotC != nil) != tt.hasConfiguration {
				t.Errorf("ValidateConfiguration() gotC = %v, hasConfiguration %t", gotC, tt.hasConfiguration)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("ValidateConfiguration() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if (len(gotRe) > 0) != tt.hasErrors {
				t.Errorf("ValidateConfiguration() got errors = %v, but wanted has errors %t", gotRe, tt.hasErrors)
			}

			if tt.hasErrors {
				for i, actualError := range gotRe {
					assertErr := tt.errors[i]
					if actualError.String() != assertErr {
						t.Errorf("ResultErrors got error = %s, but wanted error to be %s", actualError.String(), assertErr)
					}
				}
			}
		})
	}
}

func TestJsonYamlCompatibility(t *testing.T) {
	jsonC, _, _ := LoadConfiguration("./example/example_configuration.json")
	yamlC, _, _ := LoadConfiguration("./example/example_configuration.yml")

	if !reflect.DeepEqual(jsonC, yamlC) {
		t.Errorf("deep equality jsonC = %+v, yamlC %+v", jsonC, yamlC)
	}
}
