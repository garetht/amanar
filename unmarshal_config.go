package amanar

import (
	"encoding/json"
	"fmt"
	"github.com/icza/dyno"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
)

func unmarshalConfiguration(bytes []byte) (*Amanar, error) {
	c, err := UnmarshalYamlAmanarConfiguration(bytes)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal amanar configuration: %w", err)
	}

	return &c, err
}

func convertToJson(bytes []byte) ([]byte, error) {
	c, err := DynamicUnmarshalYamlAmanarConfiguration(bytes)

	if err != nil {
		return nil, err
	}

	marshalableMap := dyno.ConvertMapI2MapS(c)

	return json.Marshal(marshalableMap)
}

//go:generate go-bindata -pkg amanar amanar_config_schema.json
func LoadConfiguration(configFilepath string) (*Amanar, error, []gojsonschema.ResultError) {
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("could not read amanar configuration file: %w", err), nil
	}

	configuration, err := convertToJson(bytes)
	if err != nil {
		return nil, fmt.Errorf("could not load amanar configuration to JSON for validation: %w", err), nil
	}

	validator := NewJsonBytesSchemaValidator(configuration)
	err, validationErrors := validator.Validate()

	if err != nil {
		return nil, fmt.Errorf("[JSON CONFIG SCHEMA VALIDATION] could not run validation for JSON: %w", err), nil
	}

	parsedConfiguration, err := unmarshalConfiguration(bytes)
	if err != nil {
		return nil, fmt.Errorf("could not load amanar configuration to struct: %w", err), nil
	}

	return parsedConfiguration, err, validationErrors
}


