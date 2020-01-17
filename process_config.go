package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

func ProcessConfigItem(configurables *Configurables, credentials *Credentials) {
	var errs []error
	var flows []Flower

	for _, datasourceConfig := range configurables.IntellijDatasources {
		flow, err := NewIntellijDatasourceFlow(&datasourceConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, runConfigurationsConfig := range configurables.IntellijRunConfigurations {
		flow, err := NewIntellijRunConfigsFlow(&runConfigurationsConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, querious2Config := range configurables.Querious2Datasources {
		flow, err := NewQuerious2Flow(&querious2Config)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)	}

	for _, sequelProConfig := range configurables.SequelProDatasources {
		flow, err := NewSequelProFlow(&sequelProConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, posticoConfig := range configurables.PosticoDatasources {
		flow, err := NewPosticoFlow(&posticoConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, shellConfig := range configurables.ShellDatasources {
		flow, err := NewShellFlow(&shellConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, jsonConfig := range configurables.JSONDatasources {
		flow, err := NewJSONFlow(&jsonConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, templateConfig := range configurables.TemplateDatasources {
		flow, err := NewTemplateFlow(&templateConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}


	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("[FLOW PROCESSING] Encountered errors processing flow: %#v. Processing flows that worked.", err)
		}
	}

	UpdateCredentials(flows, credentials)

	return
}

func unmarshalConfiguration(bytes []byte) (c AmanarConfiguration, err error) {
	c, err = UnmarshalYamlAmanarConfiguration(bytes)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal amanar configuration: %w", err)
	}

	return c, err
}

func LoadConfiguration(configFilepath, schemaAssetPath string) (AmanarConfiguration, error, []gojsonschema.ResultError) {
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("could not read amanar configuration file: %w", err), nil
	}

	configuration, err := unmarshalConfiguration(bytes)
	if err != nil {
		return nil, fmt.Errorf("could not load amanar configuration: %w", err), nil
	}

	err, validationErrors := validateConfiguration(schemaAssetPath, &configuration)

	return configuration, err, validationErrors
}

func validateConfiguration(schemaAssetPath string, configuration *AmanarConfiguration) (err error, re []gojsonschema.ResultError) {
	schema, err := Asset(schemaAssetPath)
	if err != nil {
		return fmt.Errorf("could not load schema assets: %w", err), nil
	}

	documentLoader := gojsonschema.NewGoLoader(configuration)
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

func UpdateCredentials(flows []Flower, credentials *Credentials) {
	var err error
	for _, flow := range flows {
		log.Printf("[%s] Beginning to update flow %#v with credentials %s", flow.Name(), flow, credentials)

		err = flow.UpdateWithCredentials(credentials)
		if err != nil {
			log.Printf("[%s] Error when performing non-write update to flow %#v with credentials %s. Will not try and persist externally. Skipping ahead to next flow.", flow.Name(), flow, credentials)
			log.Print(err)
			continue
		}

		err = flow.PersistChanges()
		if err != nil {
			log.Printf("[%s] Error when persisting changes to to flow %#v with credentials %s. Skipping ahead to next flow.", flow.Name(), flow, credentials)
			log.Print(err)
		}
	}
}
