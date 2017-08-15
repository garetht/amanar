package main

import (
	"log"
	"fmt"
)

type Querious2Flow struct {
	Querious2DatasourcesConfig
	NewCredentials *Credentials
}

func (qf *Querious2Flow) UpdateCredentials() (err error) {
	database, err := NewQuerious2SQLiteDatabase(qf.Querious2SQLitePath)

	if err != nil {
		return
	}

	err = database.UpdateUsername(qf.DatabaseUUID, qf.NewCredentials)
	if err != nil {
		return
	}

	service := fmt.Sprintf("MySQL %s", qf.DatabaseUUID)

	log.Printf("[QUERIOUS2 DATASOURCE %s] Writing new username %s and password %s to Keychain", service, qf.NewCredentials.Username, qf.NewCredentials.Password)
	// Querious 2 finds its item in the keychain based a hashlike combination of the keychain filepath,
	// account, and service. We therefore do not alter any of these things./
	// (connection_settings.keychainItemRefMySQL)
	err = CreateOrUpdateKeychainEntriesForService(service, "", qf.NewCredentials.Password, []string{})
	if err != nil {
		log.Print(err)
		log.Fatalf("[QUERIOUS2 DATASOURCE %s] Could not create the new keychain entry with username %s and password %s", service, qf.NewCredentials.Username, qf.NewCredentials.Password)
		return
	}

	return nil
}
