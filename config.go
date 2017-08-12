package main

type AmanarConfigItem struct {
	VaultPath      string `json:"vault_path"`
	VaultRole      string `json:"vault_role"`
	Configurations struct {
		IntellijDatasources       []IntellijDatasourceConfig        `json:"intellij_datasources"`
		IntellijRunConfigurations []IntellijRunConfigurationsConfig `json:"intellij_run_configurations"`
	}
}

type IntellijDatasourceConfig struct {
	DatasourceFilePath  string   `json:"datasource_file_path"`
	DatabaseUUID        string   `json:"database_uuid"`
	TrustedApplications []string `json:"trusted_applications"`
}

type IntellijRunConfigurationsConfig struct {
	RunConfigurationsFolderPath string `json:"run_configurations_folder_path"`
	EnvironmentVariable         string `json:"environment_variable"`
}
