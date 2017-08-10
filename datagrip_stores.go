package main

import (
	"errors"
	"fmt"

	"log"

	"github.com/hashicorp/vault/api"
	"github.com/zalando/go-keyring"
)

// A DatagripStores knows how to update the username
// and password for a single database with a particular IntelliJ UUID.
type DatagripStores struct {
	DatabaseUUID     string
	DatagripFilepath string
	VaultPath        string
	VaultRole        string
	AuthedClient     *VaultGithubAuthClient
}

func (ds *DatagripStores) readNewCredentials() (secret *api.Secret, err error) {
	// Reading: gets new credentials from Vault
	secret, err = ds.AuthedClient.getCredential(ds.VaultPath, ds.VaultRole)
	if err != nil {
		return
	}
	return
}

func (ds *DatagripStores) updateCredentials(datagripConfig *DatagripConfig, secret *api.Secret) (string, error) {
	// Updaters: updates to in-memory data. Should be done
	// sequentially, but no IO is done.
	oldUsername, err := datagripConfig.UpdateUsername(ds.DatabaseUUID, secret)

	if err != nil {
		return "", err
	}

	return oldUsername, nil
}

func (ds *DatagripStores) writeCredentials(config *DatagripConfig, secret *api.Secret, oldUsername string) (err error) {
	// Writing: side-effecting writes to files and forms of IO and things.
	// In this case, we write to the IntellJ config file and the OSX keychain
	// Should be done sequentially.

	service := fmt.Sprintf("IntelliJ Platform DB — %s", ds.DatabaseUUID)

	newUsername, ok := secret.Data["username"].(string)
	if !ok {
		return errors.New("Could not parse username out of Vault secret response.")
	}

	password, ok := secret.Data["password"].(string)
	if !ok {
		return errors.New("Could not parse password out of Vault secret response.")
	}

	err = config.Document.WriteToFile(ds.DatagripFilepath)
	if err != nil {
		return err
	}

	fmt.Println("New username and password", newUsername, password)
	fmt.Println("Service and old username", service, oldUsername)
	fmt.Println(keyring.Get(service, oldUsername))

	err = keyring.Set(service, newUsername, password)
	if err != nil {
		log.Fatalf("Could not create the new keychain entry for service %s with username %s and password %s", service, newUsername, password)
		return err
	}

	if oldUsername != "" {
		err = keyring.Delete(service, oldUsername)
		if err != nil {
			log.Fatalf("Could not delete the old secret for service %s with username %s", service, oldUsername)
			return err
		}
	}

	fmt.Println("Finished changing")

	return
}

// A side effecting function that updates the
func (ds *DatagripStores) RefreshCredentials() (err error) {
	secret, err := ds.readNewCredentials()
	if err != nil {
		return
	}

	datagripConfig, err := NewDatagripConfig(ds.DatagripFilepath)
	if err != nil {
		return
	}

	oldUsername, err := ds.updateCredentials(datagripConfig, secret)
	if err != nil {
		return
	}

	err = ds.writeCredentials(datagripConfig, secret, oldUsername)
	if err != nil {
		return
	}

	return nil
}
