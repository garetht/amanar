package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

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
	Querious2Datasources      []Querious2DatasourcesConfig      `json:"querious2_datasources"`
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

type Querious2DatasourcesConfig struct {
	Querious2SQLitePath string `json:"querious2_sqlite_path"`
	DatabaseUUID        string `json:"database_uuid"`
}


func ProcessConfigItem(configurables *AmanarConfigurables, credentials *Credentials) {
	var err error
	for _, datasourceConfig := range configurables.IntellijDatasources {
		log.Printf("[DATSOURCE CONFIG] Processing datasource config at %s with UUID %s", datasourceConfig.DatasourceFilePath, datasourceConfig.DatabaseUUID)
		err = processDatasourceConfig(&datasourceConfig, credentials)
		if err != nil {
			log.Printf("[DATASOURCE CONFIG] Could not process datasource config %#v because %s. Skipping ahead.", datasourceConfig, err)
		}
	}

	for _, runConfigurationsConfig := range configurables.IntellijRunConfigurations {
		log.Printf("[RUN CONFIGURATIONS CONFIG] Processing run configurations config at %s", runConfigurationsConfig.RunConfigurationsFolderPath)
		err = processRunConfigurationsConfig(&runConfigurationsConfig, credentials)
		if err != nil {
			log.Printf("[RUN CONFIGURATIONS CONFIG] Could not process run configurations config %#v because %s. Skipping ahead.", runConfigurationsConfig, err)
		}
	}

	for _, querious2Config := range configurables.Querious2Datasources {
		log.Printf("[QUERIOUS 2 CONFIG] Processing Querious 2 SQLite database at %s", querious2Config.Querious2SQLitePath)
		err = processQuerious2Config(&querious2Config, credentials)
		if err != nil {
			log.Printf("[QUERIOUS 2 CONFIG] Could not process Querious 2 %#v because %s. Skipping ahead.", querious2Config, err)
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

func processDatasourceConfig(datasourceConfig *IntellijDatasourceConfig, credentials *Credentials) error {
	flow := IntellijDatasourceFlow{
		IntellijDatasourceConfig: *datasourceConfig,
		NewCredentials:           credentials,
	}
	return flow.UpdateCredentials()
}

func processRunConfigurationsConfig(runConfigurationsConfig *IntellijRunConfigurationsConfig, credentials *Credentials) error {
	flow := IntellijRunConfigsFlow{
		IntellijRunConfigurationsConfig: *runConfigurationsConfig,
		NewCredentials:                  credentials,
	}

	return flow.UpdateCredentials()
}

func processQuerious2Config(querious2Config *Querious2DatasourcesConfig, credentials *Credentials) error {
	flow := Querious2Flow{
		Querious2DatasourcesConfig: *querious2Config,
		NewCredentials:             credentials,
	}

	return flow.UpdateCredentials()
}
