package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Store struct {
	Db *sql.DB
}

func NewStore() Store {

	dbStore := Store{}

	if err := dbStore.getConnection(); err != nil {
		log.Fatalf("failed to connect to the database... Error: %s", err)
	}

	if err := createMigrations(dbStore.Db); err != nil {
		log.Fatalln(err)
	}

	return dbStore
}

func (dbStore *Store) getConnection() error {

	if dbStore.Db != nil {
		return nil
	}

	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "rdvigor"
	)

	if os.Getenv("DEPLOY_TYPE") == "PRODUCTION" {
		host = os.Getenv("PGHOST")
		user = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		port, _ = strconv.Atoi(os.Getenv("PGPORT"))
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Printf("%s\n", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	dbStore.Db = db
	log.Println("Connected successfully to the database...")

	return nil
}

func createMigrations(db *sql.DB) error {

	statement := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
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
