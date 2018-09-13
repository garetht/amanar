package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func NewJSONFlow(config *JSONDatasource) (*JSONFlow, error) {
	file, err := os.OpenFile(config.Filepath, os.O_RDONLY|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		return nil, fmt.Errorf("could not open and parse JSON file path '%s': %s", config.Filepath, err)
	}

	rawJSON, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read from JSON file '%s': %s", config.Filepath, err)
	}


	existingCredentials, err := UnmarshalJSONCredentials(rawJSON)
	if err != nil {
		existingCredentials = JSONCredentials{}
	}

	return &JSONFlow{
		JSONDatasource: *config,
		parsedFile:     existingCredentials,
	}, nil
}

type JSONFlow struct {
	JSONDatasource
	credentials *Credentials
	parsedFile  JSONCredentials
}

func (sf *JSONFlow) Name() string {
	return "JSON"
}

func (sf *JSONFlow) UpdateWithCredentials(credentials *Credentials) error {
	log.Printf("[%s DATASOURCE] Updating JSON credentials %s with new username %s and password %s", sf.Name(), sf.Filepath, credentials.Username, credentials.Password)
	sf.credentials = credentials

	found := false

	for i, entry := range sf.parsedFile {
		if entry.Identifier == sf.Identifier {
			sf.parsedFile[i].Username = credentials.Username
			sf.parsedFile[i].Password = credentials.Password
			found = true
		}
	}

	if !found {
		sf.parsedFile = append(sf.parsedFile, JSONCredential{
			Identifier: sf.Identifier,
			Username: credentials.Username,
			Password: credentials.Password,
		})
	}

	log.Printf("[%s DATASOURCE] Updated JSON credentials %s with new username %s and password %s", sf.Name(), sf.Filepath, credentials.Username, credentials.Password)
	return nil
}

func (sf *JSONFlow) PersistChanges() error {
	log.Printf("[%s DATASOURCE] Writing new username %s and password %s to JSON file", sf.Name(), sf.credentials.Username, sf.credentials.Password)

	writeData, err := sf.parsedFile.Marshal()
	if err != nil {
		return fmt.Errorf("could not marshal JSON to string: %#v", writeData)
	}

	err = ioutil.WriteFile(sf.Filepath, writeData, 0644)
	if err != nil {
		return fmt.Errorf("could not write JSON to file '%s': %s", sf.Filepath, writeData)
	}

	log.Printf("[%s DATASOURCE] Successfully wrote new username %s and password %s to JSON file", sf.Name(), sf.credentials.Username, sf.credentials.Password)
	return nil
}
