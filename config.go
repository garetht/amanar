package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hashicorp/vault/api"
	"github.com/xeipuuv/gojsonschema"
)

type AmanarConfiguration []AmanarConfigItem

type AmanarConfigItem struct {
	VaultPath     VaultPath           `json:"vault_path"`
	VaultRole     VaultRole           `json:"vault_role"`
	Configurables AmanarConfigurables `json:"configurables"`
}

type AmanarConfigurables struct {
	IntellijDatasources       []IntellijDatasourceConfig        `json:"intellij_datasources"`
	IntellijRunConfigurations []IntellijRunConfigurationsConfig `json:"intellij_run_configurations"`
}

type IntellijDatasourceConfig struct {
	DatasourceFilePath  IntellijDatasourceFilepath `json:"datasource_file_path"`
	DatabaseUUID        IntellijDatabaseUUID       `json:"database_uuid"`
	TrustedApplications []string                   `json:"trusted_applications"`
}

type IntellijRunConfigurationsConfig struct {
	RunConfigurationsFolderPath string `json:"run_configurations_folder_path"`
	EnvironmentVariable         string `json:"environment_variable"`
	DatabaseHost                string `json:"database_host"`
}

func ProcessConfigItem(configurables *AmanarConfigurables, secret *api.Secret) {
	var err error
	for _, datasourceConfig := range configurables.IntellijDatasources {
		log.Printf("[DATSOURCE CONFIG] Processing datasource config at %s with UUID %s", datasourceConfig.DatasourceFilePath, datasourceConfig.DatabaseUUID)
		err = processDatasourceConfig(&datasourceConfig, secret)
		if err != nil {
			log.Printf("[DATASOURCE CONFIG] Could not process datasource config %#v because %s. Skipping ahead.", datasourceConfig, err)
		}
	}

	for _, runConfigurationsConfig := range configurables.IntellijRunConfigurations {
		log.Printf("[RUN CONFIGURATIONS CONFIG] Processing run configurations config at %s", runConfigurationsConfig.RunConfigurationsFolderPath)
		err = processRunConfigurationsConfig(&runConfigurationsConfig, secret)
		if err != nil {
			log.Printf("[RUN CONFIGURATIONS CONFIG] Could not process run configurations config %#v because %s. Skipping ahead.", runConfigurationsConfig, err)
		}
	}
	return
}

func LoadConfiguration(configFilepath, schemaAssetPath string) (c AmanarConfiguration, err error, re []gojsonschema.ResultError) {
	bytes, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return
	}

	schema, err := Asset(schemaAssetPath)
	if err != nil {
		return
	}

	documentLoader := gojsonschema.NewBytesLoader(bytes)
	schemaLoader := gojsonschema.NewBytesLoader(schema)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return
	}

	if !result.Valid() {
		re = result.Errors()
		return
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return
	}

	return
}

func processDatasourceConfig(datasourceConfig *IntellijDatasourceConfig, secret *api.Secret) error {
	flow := IntellijDatasourceFlow{
		IntellijDatasourceConfig: *datasourceConfig,
		NewVaultSecret:           secret,
	}
	return flow.UpdateCredentials()
}

func processRunConfigurationsConfig(runConfigurationsConfig *IntellijRunConfigurationsConfig, secret *api.Secret) (err error) {
	flow := IntellijRunConfigsFlow{
		IntellijRunConfigurationsConfig: *runConfigurationsConfig,
		NewVaultSecret:                  secret,
	}

	return flow.UpdateCredentials()
}
