package amanar

import (
	"fmt"
	"log"
)

func NewShellFlow(config *ShellDatasource) (*ShellFlow, error) {
	parsedFile, err := NewShellFile(config.Filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open and parse shell file path '%s': %s", config.Filepath, err)
	}

	return &ShellFlow{
		ShellDatasource: *config,
		parsedFile:      parsedFile,
	}, nil
}

type ShellFlow struct {
	ShellDatasource
	credentials *Credentials
	parsedFile  *ShellFile
}

func (sf *ShellFlow) Name() string {
	return "SHELL"
}

func (sf *ShellFlow) UpdateWithCredentials(credentials *Credentials) error {
	log.Printf("[%s DATASOURCE] Updating parsed shell AST %s with new username %s and password %s", sf.Name(), sf.Filepath, credentials.Username, credentials.Password)
	sf.parsedFile.UpdateCredentials(sf.UsernameVariable, sf.PasswordVariable, credentials)
	sf.credentials = credentials
	log.Printf("[%s DATASOURCE] Updated parsed shell AST %s with new username %s and password %s", sf.Name(), sf.Filepath, credentials.Username, credentials.Password)
	return nil
}

func (sf *ShellFlow) PersistChanges() error {
	log.Printf("[%s DATASOURCE] Writing new username %s and password %s to shell script file", sf.Name(), sf.credentials.Username, sf.credentials.Password)
	sf.parsedFile.WriteToDisk()
	log.Printf("[%s DATASOURCE] Successfully wrote new username %s and password %s to shell script file", sf.Name(), sf.credentials.Username, sf.credentials.Password)
	return nil
}
