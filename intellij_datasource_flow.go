package main

import (
	"fmt"

	"log"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type IntellijDatabaseUUID string
type IntellijDatasourceFilepath string

func NewIntellijDatasourceFlow(config *IntellijDatasourceConfig) (*IntellijDatasourceFlow, error) {
	datasourceFile, err := NewIntellijDatasourceFile(config.DatasourceFilePath)
	if err != nil {
		return nil, err
	}

	return &IntellijDatasourceFlow{
		IntellijDatasourceConfig: *config,
		datasourceFile: datasourceFile,
	}, nil
}

// An IntellijDatasourceFlow knows how to update the username
// and password for a single database with a particular IntelliJ UUID.
// It does this by using dataSources.local.xml files.
type IntellijDatasourceFlow struct {
	IntellijDatasourceConfig
	datasourceFile *IntellijDatasourceFile
	credentials *Credentials
}

func (ds *IntellijDatasourceFlow) UpdateWithCredentials(credentials *Credentials) (err error) {
	_, err = ds.datasourceFile.UpdateUsername(ds.DatabaseUUID, credentials.Username)
	if err != nil {
		return
	}

	ds.credentials = credentials
	return nil
}


func (ds *IntellijDatasourceFlow) PersistChanges() (err error) {
	// Writing: side-effecting writes to files and forms of IO and things.
	// In this case, we write to the IntellJ config file and the OSX keychain
	// Should be done sequentially.

	service := fmt.Sprintf("IntelliJ Platform DB — %s", ds.DatabaseUUID)

	err = ds.datasourceFile.WriteToFile()
	if err != nil {
		return err
	}

	log.Printf("[INTELLIJ DATASOURCE %s] Writing new username %s and password %s to Keychain", service, ds.credentials.Username, ds.credentials.Password)
	err = CreateOrUpdateKeychainEntriesForService(service, ds.credentials.Username, ds.credentials.Password, []string{})
	if err != nil {
		log.Print(err)
		log.Fatalf("[INTELLIJ DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", service, ds.credentials.Username, ds.credentials.Password)
		return err
	}

	return
}

// A side effecting function that updates the
func (ds *IntellijDatasourceFlow) UpdateCredentials(credentials *Credentials) (err error) {
	err = ds.UpdateWithCredentials(credentials)
	if err != nil {
		return
	}

	err = ds.PersistChanges()
	if err != nil {
		return
	}

	return nil
}
