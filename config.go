package main

import (
	"github.com/hashicorp/vault/api"
	"log"
)

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

func processDatasourceConfig(datasourceConfig *IntellijDatasourceConfig, secret *api.Secret) error {
	source := IntellijDatasource{
		DatabaseUUID: datasourceConfig.DatabaseUUID,
		DatasourceFilepath: datasourceConfig.DatasourceFilePath,
		NewVaultSecret: secret,
		TrustedApplications: datasourceConfig.TrustedApplications,
	}
	return source.UpdateCredentials()
}

func processRunConfigurationsConfig(runConfigurationsConfig *IntellijRunConfigurationsConfig, secret *api.Secret) error {
	return nil
}
