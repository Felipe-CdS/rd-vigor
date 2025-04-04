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

	statement = `CREATE TABLE IF NOT EXISTS users_tags ( 
		id UUID PRIMARY KEY
		, fk_tag_id UUID NOT NULL
		, FOREIGN KEY(fk_tag_id) REFERENCES tags(tag_id) ON DELETE CASCADE
		, fk_user_id UUID NOT NULL
		, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON DELETE CASCADE
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 3... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS chatrooms ( 
		chatroom_id UUID PRIMARY KEY
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 4... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS chatrooms_users ( 
		fk_user_id UUID NOT NULL
		, fk_chatroom_id UUID NOT NULL
		, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON DELETE CASCADE
		, FOREIGN KEY(fk_chatroom_id) REFERENCES chatrooms(chatroom_id) ON DELETE CASCADE
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 5... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS messages ( 
		message_id UUID PRIMARY KEY
		, fk_sender_id UUID NOT NULL
		, FOREIGN KEY(fk_sender_id) REFERENCES users(id) ON DELETE CASCADE
		, fk_chatroom_id UUID NOT NULL
		, FOREIGN KEY(fk_chatroom_id) REFERENCES chatrooms(chatroom_id) ON DELETE CASCADE
		, content TEXT NOT NULL
		, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 6... Error: %s", err)
	}

	statement = `ALTER TABLE users
		ADD COLUMN IF NOT EXISTS address TEXT DEFAULT '' NOT NULL
		, ADD COLUMN IF NOT EXISTS address2 TEXT DEFAULT '' NOT NULL
		, ADD COLUMN IF NOT EXISTS city TEXT DEFAULT '' NOT NULL
		, ADD COLUMN IF NOT EXISTS state TEXT DEFAULT '' NOT NULL
		, ADD COLUMN IF NOT EXISTS zipcode TEXT DEFAULT '' NOT NULL;`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 7... Error: %s", err)
	}

	statement = `ALTER TABLE users ADD COLUMN IF NOT EXISTS stripe_id TEXT DEFAULT '' NOT NULL;`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 8... Error: %s", err)
	}

	statement = `ALTER TABLE users ADD COLUMN IF NOT EXISTS subscription_status BOOLEAN DEFAULT false NOT NULL;`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 9... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS portifolios ( 
		portifolio_id UUID PRIMARY KEY
		, fk_user_id UUID NOT NULL
		, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON DELETE CASCADE
		, title TEXT  NOT NULL
		, description TEXT NOT NULL
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 10... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS events ( 
		id UUID PRIMARY KEY
		, title TEXT NOT NULL DEFAULT ''
		, description TEXT NOT NULL DEFAULT ''
		, cover_path TEXT NOT NULL DEFAULT '/assets/img/events-1.png'
		, maps_link TEXT NOT NULL DEFAULT ''
		, address TEXT NOT NULL DEFAULT ''
		, date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		, price INTEGER NOT NULL DEFAULT 0
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 11... Error: %s", err)
	}

	statement = `CREATE TABLE IF NOT EXISTS user_event_confirmations ( 
		fk_user_id UUID NOT NULL
		, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON DELETE CASCADE
		, fk_event_id UUID NOT NULL
		, FOREIGN KEY(fk_event_id) REFERENCES events(id) ON DELETE CASCADE
	)`

	_, err = db.Exec(statement)

	if err != nil {
		return fmt.Errorf("failed to create statement 12... Error: %s", err)
	}

	return nil
}
