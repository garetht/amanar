// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    amanarConfiguration, err := UnmarshalAmanarConfiguration(bytes)
//    bytes, err = amanarConfiguration.Marshal()

package main

import "encoding/json"

type AmanarConfiguration []AmanarConfigurationElement

func UnmarshalAmanarConfiguration(data []byte) (AmanarConfiguration, error) {
	var r AmanarConfiguration
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AmanarConfiguration) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AmanarConfigurationElement struct {
	VaultAddress       string               `json:"vault_address"`      // The address to a particular vault. Vault addresses usually differ for different; environments. For example, we may have one vault address for production and another for; staging.
	VaultConfiguration []VaultConfiguration `json:"vault_configuration"`
}

// A list of vault roles and paths and configuration options for output to data sources
// within a particular vault environment.
type VaultConfiguration struct {
	Configurables Configurables `json:"configurables"`
	VaultPath     string        `json:"vault_path"`   // The path representing the datastore in the Vault. This is equivalent to $VAULT_PATH in; the CLI command `vault read $VAULT_PATH/creds/$VAULT_ROLE`.
	VaultRole     string        `json:"vault_role"`   // The role representing the permissions that are sought to the Vault datastore. This is; equivalent to $VAULT_ROLE in the CLI command `vault read $VAULT_PATH/creds/$VAULT_ROLE`.
}

type Configurables struct {
	IntellijDatasources       []IntellijDatasource       `json:"intellij_datasources"`       // Allows IntelliJ datasource usernames and passwords to be changed. Most useful for; DataGrip and databases within IntelliJ Ultimate.
	IntellijRunConfigurations []IntellijRunConfiguration `json:"intellij_run_configurations"`// Allows changes to database access credentials within IntelliJ run configurations.
	JSONDatasources           []JSONDatasource           `json:"json_datasources"`           // Allows a JSON file to be generated containing usernames and passwords.
	PosticoDatasources        []PosticoDatasource        `json:"postico_datasources"`        // Allows changes to database access credentials stored in a Postico SQLite database.
	Querious2Datasources      []Querious2Datasource      `json:"querious2_datasources"`      // Allows changes to database access credentials stored in a Querious 2 SQLite database.
	SequelProDatasources      []SequelProDatasource      `json:"sequel_pro_datasources"`     // Allows changes to database access credentials for Sequel Pro plists.
	ShellDatasources          []ShellDatasource          `json:"shell_datasources"`          // Allows a file to be generated in a shell script that contains exports of environment; variables containing the new credentials.
}

type IntellijDatasource struct {
	DatabaseUUID       string `json:"database_uuid"`       // The IntelliJ UUID for the database you want to update. You can find this by examining the; dataSources.local.xml file.
	DatasourceFilePath string `json:"datasource_file_path"`// The path to IntelliJ data sources file. The file is typically called; dataSources.local.xml.
}

type IntellijRunConfiguration struct {
	DatabaseHost                string `json:"database_host"`                 // The username and password for the URL will only be updated if the host of URL in the; environment variable matches this string.
	EnvironmentVariable         string `json:"environment_variable"`          // The environment variable in the run configuration under which the database connection
	RunConfigurationsFolderPath string `json:"run_configurations_folder_path"`// A directory containing all IntelliJ run configurations to be examined. Usually located in; .idea/runConfigurations. Run configurations may need to be shared before becoming visible; in this folder.
}

type JSONDatasource struct {
	Filepath   string `json:"filepath"`  // The path the JSON file should be generated to.
	Identifier string `json:"identifier"`// The name of this vault role and vault path pair to be used as an identifier for this JSON; object.
}

type PosticoDatasource struct {
	DatabaseUUID      string `json:"database_uuid"`      // The unique identifier for the Postico database to update. Can be found by looking in the; SQLite database.
	PosticoSqlitePath string `json:"postico_sqlite_path"`// Path to the SQLite database in which Postico stores its data. The file is typically; called ConnectionFavorites.db
}

type Querious2Datasource struct {
	DatabaseUUID        string `json:"database_uuid"`        // The unique identifier for the Querious database to update. Can be found by looking in the; SQLite database.
	Querious2SqlitePath string `json:"querious2_sqlite_path"`// Path to the SQLite database in which Querious 2 stores its data. The file is typically; called Connections.sqlite.
}

type SequelProDatasource struct {
	DatabaseUUID       string `json:"database_uuid"`        // The unique identifier for the Sequel Pro database to update. Can be found by looking in; the plist.
	SequelProPlistPath string `json:"sequel_pro_plist_path"`// Path to the plist in which Sequel Pro stores its data. The file is typically called; Favorites.plist
}

type ShellDatasource struct {
	Filepath         string `json:"filepath"`         // The path the shell script should be generated to.
	PasswordVariable string `json:"password_variable"`// The name of the environment variable that should contain the password
	UsernameVariable string `json:"username_variable"`// The name of the environment variable that should contain the username
}
