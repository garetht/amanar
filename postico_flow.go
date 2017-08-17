package main

import (
	"fmt"
	"log"
)

func NewPosticoFlow(config *PosticoDatasourcesConfig) (*PosticoFlow, error) {
	database, err := NewPosticoSQLiteDatabase(config.PosticoSQLitePath)
	if err != nil {
		return nil, err
	}

	return &PosticoFlow{
		PosticoDatasourcesConfig: *config,
		database:                 database,
	}, nil
}

type PosticoFlow struct {
	PosticoDatasourcesConfig
	database    *PosticoSQLiteDatabase
	credentials *Credentials
}

func (pf *PosticoFlow) UpdateWithCredentials(credentials *Credentials) error {
	pf.credentials = credentials
	return nil
}

func (pf *PosticoFlow) PersistChanges() (err error) {
	err = pf.database.UpdateUsername(pf.DatabaseUUID, pf.credentials.Username)
	if err != nil {
		return
	}

	uuidRow, err := pf.database.GetFavoriteFromUUID(pf.DatabaseUUID)
	if err != nil {
		return
	}

	var host string
	if uuidRow.Host.Valid {
		host = uuidRow.Host.String
	}

	service := fmt.Sprintf("postgresql://%s", host)
	log.Printf("[POSTICO DATASOURCE %s] Writing new username %s and password %s to Keychain", service, pf.credentials.Username, pf.credentials.Password)
	// Querious 2 finds its item in the keychain based a hashlike combination of the keychain filepath,
	// account, and service. We therefore do not alter any of these things./
	// (connection_settings.keychainItemRefMySQL)
	err = CreateOrUpdateKeychainEntriesForService(service, pf.credentials.Username, pf.credentials.Password, []string{})
	if err != nil {
		log.Print(err)
		log.Fatalf("[POSTICO DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", service, pf.credentials.Username, pf.credentials.Password)
		return
	}

	return nil
}

func (pf *PosticoFlow) UpdateCredentials(credentials *Credentials) (err error) {
	err = pf.UpdateWithCredentials(credentials)
	if err != nil {
		return
	}

	err = pf.PersistChanges()
	if err != nil {
		return
	}

	return nil
}
