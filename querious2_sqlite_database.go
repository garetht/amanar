package main

import (
	"database/sql"
	"github.com/hashicorp/vault/api"
)

func NewQuerious2SQLiteDatabase(filepath string) (*Querious2SQLiteDatabase, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	database := &Querious2SQLiteDatabase{
		Database: db,
	}

	return database, nil
}

type Querious2SQLiteDatabase struct {
	Database *sql.DB
}


func (qdb *Querious2SQLiteDatabase) WriteToDatabase(uuid string, secret *api.Secret) error {
	statement, err := qdb.Database.Prepare("UPDATE connection_settings set user=? where uuid=?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(newUsername, uuid)


}
