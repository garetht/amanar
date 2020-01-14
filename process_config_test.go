package main

import (
	"reflect"
	"testing"
)

var config = "amanar_config_schema.json"

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
				schemaAssetPath: config,
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
		},
		{
			name: "Valid YAML can be loaded",
			args: args{
				configFilepath:  "./example/example_configuration.yml",
				schemaAssetPath: config,
			},
			hasConfiguration: true,
			wantErr:          nil,
			hasErrors:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, gotErr, gotRe := LoadConfiguration(tt.args.configFilepath, tt.args.schemaAssetPath)
			if (gotC != nil) != tt.hasConfiguration {
				t.Errorf("validateConfiguration() gotC = %v, hasConfiguration %t", gotC, tt.hasConfiguration)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("validateConfiguration() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if (len(gotRe) > 0) != tt.hasErrors {
				t.Errorf("validateConfiguration() gotRe = %v, hasErrors %t", gotRe, tt.hasErrors)
			}
		})
	}
}

func TestJsonYamlCompatibility(t *testing.T) {
	jsonC, _, _ := LoadConfiguration("./example/example_configuration.json", config)
	yamlC, _, _ := LoadConfiguration("./example/example_configuration.yml", config)

	if !reflect.DeepEqual(jsonC, yamlC) {
		t.Errorf("deep equality jsonC = %+v, yamlC %+v", jsonC, yamlC)
	}
}
