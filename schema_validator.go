package amanar

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateConfiguration(documentLoader gojsonschema.JSONLoader) (err error, re []gojsonschema.ResultError) {
	schema, err := Asset("amanar_config_schema.json")
	if err != nil {
		return fmt.Errorf("could not load schema assets: %w", err), nil
	}

	// We validate the Go document in order to be able to use gojsonschema to validate YAML
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("was not able to run validation on schema: %w", err), nil
	}

	if !result.Valid() {
		re = result.Errors()
		return
	}

	return
}

type SchemaValidator interface {
	loader() gojsonschema.JSONLoader
	Validate() (err error, re []gojsonschema.ResultError)
}

type StructSchemaValidator struct {
	goStruct *Amanar
}

func (s StructSchemaValidator) loader() gojsonschema.JSONLoader {
	return gojsonschema.NewGoLoader(s.goStruct)
}

func (s StructSchemaValidator) Validate() (err error, re []gojsonschema.ResultError) {
	return ValidateConfiguration(s.loader())
}

func NewStructSchemaValidator(goStruct *Amanar) *StructSchemaValidator {
	return &StructSchemaValidator{goStruct: goStruct}
}

type JsonBytesSchemaValidator struct {
	jsonBytes []byte
}

func (j JsonBytesSchemaValidator) loader() gojsonschema.JSONLoader {
	return gojsonschema.NewBytesLoader(j.jsonBytes)
}

func (j JsonBytesSchemaValidator) Validate() (err error, re []gojsonschema.ResultError) {
	return ValidateConfiguration(j.loader())
}

func NewJsonBytesSchemaValidator(jsonBytes []byte) *JsonBytesSchemaValidator {
	return &JsonBytesSchemaValidator{jsonBytes: jsonBytes}
}

