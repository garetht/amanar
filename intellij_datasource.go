package main

import (
	"errors"
	"fmt"

	"log"

	"github.com/hashicorp/vault/api"
)

type IntellijDatabaseUUID string
type IntellijDatasourceFilepath string

// An IntellijDatasource knows how to update the username
// and password for a single database with a particular IntelliJ UUID.
// It does this by using dataSources.local.xml files.
type IntellijDatasource struct {
	DatabaseUUID        IntellijDatabaseUUID
	DatasourceFilepath  IntellijDatasourceFilepath
	TrustedApplications []string
	NewVaultSecret      *api.Secret
}

func (ds *IntellijDatasource) pureUpdateCredentials(datagripConfig *IntellijDatasourceFile) (string, error) {
	// Updaters: updates to in-memory data. Should be done
	// sequentially, but no IO is done.
	oldUsername, err := datagripConfig.UpdateUsername(ds.DatabaseUUID, ds.NewVaultSecret)

	if err != nil {
		return "", err
	}

	return oldUsername, nil
}

func (ds *IntellijDatasource) writeCredentials(config *IntellijDatasourceFile) (err error) {
	// Writing: side-effecting writes to files and forms of IO and things.
	// In this case, we write to the IntellJ config file and the OSX keychain
	// Should be done sequentially.

	service := fmt.Sprintf("IntelliJ Platform DB — %s", ds.DatabaseUUID)

	newUsername, ok := ds.NewVaultSecret.Data["username"].(string)
	if !ok {
		return errors.New("Could not parse username out of Vault secret response.")
	}

	password, ok := ds.NewVaultSecret.Data["password"].(string)
	if !ok {
		return errors.New("Could not parse password out of Vault secret response.")
	}

	err = config.Document.WriteToFile(string(ds.DatasourceFilepath))
	if err != nil {
		return err
	}

	log.Printf("[DATASOURCE %s] Writing new username %s and password %s to Keychain", service, newUsername, password)
	err = CreateOrUpdateKeychainEntriesForService(service, newUsername, password, ds.TrustedApplications)
	if err != nil {
		log.Print(err)
		log.Fatalf("[DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", service, newUsername, password)
		return err
	}

	fmt.Println("Finished changing")

	return
}

// A side effecting function that updates the
func (ds *IntellijDatasource) UpdateCredentials() (err error) {
	datagripConfig, err := NewIntellijDatasourceFile(ds.DatasourceFilepath)
	if err != nil {
		return
	}

	_, err = ds.pureUpdateCredentials(datagripConfig)
	if err != nil {
		return
	}

	err = ds.writeCredentials(datagripConfig)
	if err != nil {
		return
	}

	return nil
}
