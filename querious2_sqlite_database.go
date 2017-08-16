package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewQuerious2SQLiteDatabase(filepath string) (*Querious2SQLiteDatabase, error) {
	database := &Querious2SQLiteDatabase{
		Filepath: filepath,
	}

	return database, nil
}

type Querious2SQLiteDatabase struct {
	Filepath string
}

func (qdb *Querious2SQLiteDatabase) UpdateUsername(databaseUuid string, credentials *Credentials) error {
	db, err := sql.Open("sqlite3", qdb.Filepath)

	statement, err := db.Prepare("UPDATE connection_settings SET user=? WHERE uuid=?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(credentials.Username, databaseUuid)

	defer statement.Close()

	return err
}

