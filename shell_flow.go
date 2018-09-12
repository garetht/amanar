package main

import (
	"os"
	"fmt"
	"log"
)

func NewShellFlow(config *ShellDatasource) (*ShellFlow, error) {
	file, err := os.Open(config.ScriptPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return &ShellFlow{
		ShellDatasource: *config,
		scriptFile: file,
	}, nil
}

type ShellFlow struct {
	ShellDatasource
	scriptFile  *os.File
	credentials *Credentials
	contents string
}

func (sf *ShellFlow) Name() string {
	return "SHELL"
}

func (sf *ShellFlow) UpdateWithCredentials(credentials *Credentials) error {
	sf.credentials = credentials
	// TODO: if multiple language support is needed, create an if statement
	// here switching on Language
	sf.contents = fmt.Sprintf("export %s=%s\nexport %s=%s", sf.UsernameVariable, sf.credentials.Username, sf.PasswordVariable, sf.credentials.Password)

	return nil
}

func (sf *ShellFlow) PersistChanges() error {
	log.Printf("[%s DATASOURCE] Writing new username %s and password %s to environment variable file", sf.Name(), sf.credentials.Username, sf.credentials.Password)
	_, err := sf.scriptFile.WriteString(sf.contents)
	if err != nil {
		return err
	}

	return nil
}
