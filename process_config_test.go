package main

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
	}{
		{
			name: "Valid JSON can be loaded",
			args: args{
				configFilepath:  "./example/example_configuration.json",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
		},
		{
			name: "Valid YAML can be loaded",
			args: args{
				configFilepath:  "./example/example_configuration.yml",
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
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
				t.Errorf("ValidateConfiguration() gotRe = %v, hasErrors %t", gotRe, tt.hasErrors)
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
