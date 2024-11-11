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
