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

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "rdvigor"

	if os.Getenv("APP_ENV") == "TESTING" {
		port = 8001
	}

	if os.Getenv("APP_ENV") == "PROD" {
		host = os.Getenv("PGHOST")
		user = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname = os.Getenv("PGDATABASE")
		port, _ = strconv.Atoi(os.Getenv("PGPORT"))
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	dbStore.Db = db
	log.Printf("Connected successfully to the database | ENV: %s | host: %s | dbname: %s\n", os.Getenv("APP_ENV"), host, dbname)

	return nil
}

func createMigrations(db *sql.DB) error {

	statement := `CREATE TABLE IF NOT EXISTS users ( 
		id UUID PRIMARY KEY
		, username TEXT NOT NULL
		, first_name TEXT NOT NULL
		, last_name TEXT NOT NULL
		, email TEXT NOT NULL
		, occupation_area TEXT NOT NULL
		, telephone TEXT
		, refer_friend TEXT
		, password TEXT NOT NULL
		, role TEXT NOT NULL DEFAULT 'member'
		, registration_status TEXT NOT NULL DEFAULT 'pending'
		, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		, updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`

	_, err := db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 1... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS tags ( 
		tag_id UUID PRIMARY KEY
		, tag_name TEXT NOT NULL
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 2... Error: %s", err)
	}

	return nil
}
