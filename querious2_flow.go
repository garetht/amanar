package main

import (
	"fmt"
	"log"
	"errors"
)

func NewQuerious2Flow(config *Querious2Datasource) (*Querious2Flow, error) {
	database, err := NewQuerious2SQLiteDatabase(config.Querious2SqlitePath)
	if err != nil {
		return nil, err
	}

	return &Querious2Flow{
		Querious2Datasource: *config,
		database:             database,
	}, nil

}

type Querious2Flow struct {
	Querious2Datasource
	database    *Querious2SQLiteDatabase
	credentials *Credentials
}

func (qf *Querious2Flow) Name() string {
	return "QUERIOUS 2"
}

func (qf *Querious2Flow) UpdateWithCredentials(credentials *Credentials) error {
	qf.credentials = credentials
	return nil
}

func (qf *Querious2Flow) PersistChanges() (err error) {
	if qf.credentials == nil {
		return errors.New("Please provide credentials to update")
	}

	err = qf.database.UpdateUsername(qf.DatabaseUUID, qf.credentials.Username)
	if err != nil {
		return
	}

	service := fmt.Sprintf("MySQL %s", qf.DatabaseUUID)

	log.Printf("[QUERIOUS2 DATASOURCE %s] Writing new username %s and password %s to Keychain", service, qf.credentials.Username, qf.credentials.Password)
	// Querious 2 finds its item in the keychain based a hashlike combination of the keychain filepath,
	// account, and service. We therefore do not alter any of these things./
	// (connection_settings.keychainItemRefMySQL)
	err = CreateOrUpdateKeychainEntriesForService(service, "", qf.credentials.Password, []string{})
	if err != nil {
		log.Print(err)
		log.Fatalf("[QUERIOUS2 DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", service, qf.credentials.Username, qf.credentials.Password)
		return
	}

	return nil
}
