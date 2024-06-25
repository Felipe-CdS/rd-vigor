package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

func NewStore(dbName string) Store {

	dbStore := Store{}

	if err := dbStore.getConnection(dbName); err != nil {
		log.Fatalf("failed to connect to the database... Error: %s", err)
	}

	if err := createMigrations(dbName, dbStore.Db); err != nil {
		log.Fatalln(err)
	}

	return dbStore
}

func (dbStore *Store) getConnection(dbName string) error {

	if dbStore.Db != nil {
		return nil
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	dbStore.Db = db
	log.Println("Connected successfully to the database...")

	return nil
}

func createMigrations(dbName string, db *sql.DB) error {

	statement := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
        occupation_area TEXT NOT NULL,
        created_at INTEGER NOT NULL
    )`

	_, err := db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 1... Error: %s", err)
	}

	return nil

}
