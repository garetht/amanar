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
	SequelProDatasources      []SequelProDatasourcesConfig      `json:"sequel_pro_datasources"`
	PosticoDatasources        []PosticoDatasourcesConfig        `json:"postico_datasources"`
}

type IntellijDatasourceConfig struct {
	DatasourceFilePath IntellijDatasourceFilepath `json:"datasource_file_path"`
	DatabaseUUID       IntellijDatabaseUUID       `json:"database_uuid"`
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

type SequelProDatasourcesConfig struct {
	SequelProPlistPath string `json:"sequel_pro_plist_path"`
	DatabaseUUID       string `json:"database_uuid"`
}

type PosticoDatasourcesConfig struct {
	PosticoSQLitePath string `json:"postico_sqlite_path"`
	DatabaseUUID      string `json:"database_uuid"`
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

	for _, sequelProConfig := range configurables.SequelProDatasources {
		log.Printf("[SEQUEL PRO CONFIG] Processing Sequel Pro Plist file at %s", sequelProConfig.SequelProPlistPath)
		err = processSequelProConfig(&sequelProConfig, credentials)
		if err != nil {
			log.Printf("[SEQUEL PRO CONFIG] Could not process Sequel Pro %#v because %s. Skipping ahead.", sequelProConfig, err)
		}
	}

	for _, posticoConfig := range configurables.PosticoDatasources {
		log.Printf("[POSTICO CONFIG] Processing Postico Plist file at %s", posticoConfig.PosticoSQLitePath)
		err = processPosticoConfig(&posticoConfig, credentials)
		if err != nil {
			log.Printf("[POSTICO CONFIG] Could not process Postico %#v because %s. Skipping ahead.", posticoConfig, err)
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

func UpdateCredentials(flows []Flower, credentials *Credentials) (err error) {
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

	return nil
}

func processDatasourceConfig(datasourceConfig *IntellijDatasourceConfig, credentials *Credentials) error {
	flow, err := NewIntellijDatasourceFlow(datasourceConfig)
	if err != nil {
		return err
	}

	return flow.UpdateCredentials(credentials)
}

func processRunConfigurationsConfig(runConfigurationsConfig *IntellijRunConfigurationsConfig, credentials *Credentials) error {
	flow := IntellijRunConfigsFlow{
		IntellijRunConfigurationsConfig: *runConfigurationsConfig,
	}

	return flow.UpdateCredentials(credentials)
}

func processQuerious2Config(querious2Config *Querious2DatasourcesConfig, credentials *Credentials) error {
	flow, err := NewQuerious2Flow(querious2Config)
	if err != nil {
		return err
	}

	return flow.UpdateCredentials(credentials)
}

func processSequelProConfig(sequelProConfig *SequelProDatasourcesConfig, credentials *Credentials) error {
	flow, err := NewSequelProFlow(sequelProConfig)
	if err != nil {
		return err
	}

	return flow.UpdateCredentials(credentials)
}

func processPosticoConfig(posticoConfig *PosticoDatasourcesConfig, credentials *Credentials) error {
	flow, err := NewPosticoFlow(posticoConfig)
	if err != nil {
		return err
	}

	return flow.UpdateCredentials(credentials)
}
