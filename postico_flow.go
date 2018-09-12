package main

import (
	"fmt"
	"log"
)

func NewPosticoFlow(config *PosticoDatasource) (*PosticoFlow, error) {
	database, err := NewPosticoSQLiteDatabase(config.PosticoSqlitePath)
	if err != nil {
		return nil, err
	}

	return &PosticoFlow{
		PosticoDatasource: *config,
		database:           database,
	}, nil
}

type PosticoFlow struct {
	PosticoDatasource
	database    *PosticoSQLiteDatabase
	credentials *Credentials
}

func (pf *PosticoFlow) Name() string {
	return "POSTICO"
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
	log.Printf("[%s DATASOURCE %s] Writing new username %s and password %s to Keychain", pf.Name(), service, pf.credentials.Username, pf.credentials.Password)
	// Querious 2 finds its item in the keychain based a hashlike combination of the keychain filepath,
	// account, and service. We therefore do not alter any of these things./
	// (connection_settings.keychainItemRefMySQL)
	err = CreateOrUpdateKeychainEntriesForService(service, pf.credentials.Username, pf.credentials.Password, []string{})
	if err != nil {
		log.Print(err)
		log.Fatalf("[%s DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", pf.Name(), service, pf.credentials.Username, pf.credentials.Password)
		return
	}

	return nil
}
