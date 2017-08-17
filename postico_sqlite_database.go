package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"errors"
	"fmt"
)

type PosticoFavorite struct {
	UUID     string
	User     sql.NullString
	Nickname sql.NullString
	Host     sql.NullString
	Database sql.NullString
	Port     sql.NullInt64
}

func NewPosticoSQLiteDatabase(filepath string) (*PosticoSQLiteDatabase, error) {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		return nil, err
	}

	database := &PosticoSQLiteDatabase{
		Database: db,
	}

	return database, nil
}

type PosticoSQLiteDatabase struct {
	Database *sql.DB
}

func (pdb *PosticoSQLiteDatabase) GetFavoriteFromUUID(uuid string) (*PosticoFavorite, error) {
	favorites := []PosticoFavorite{}
	rows, err := pdb.Database.Query("SELECT ZUUID, ZUSER, ZNICKNAME, ZHOST, ZDATABASE, ZPORT FROM ZPGEFAVORITE WHERE ZUUID=?", uuid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var f PosticoFavorite
		err = rows.Scan(&f.UUID, &f.User, &f.Nickname, &f.Host, &f.Database, &f.Port)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, f)
	}

	if len(favorites) != 1 {
		return nil, errors.New(fmt.Sprintf("There was not exactly one row found for UUID %s", uuid))
	}

	return &favorites[0], nil
}

func (pdb *PosticoSQLiteDatabase) UpdateUsername(databaseUuid string, username string) error {
	statement, err := pdb.Database.Prepare("UPDATE ZPGEFAVORITE SET ZUSER=? WHERE ZUUID=?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(username, databaseUuid)

	defer statement.Close()

	return err
}
