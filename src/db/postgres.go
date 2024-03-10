package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type PostgresStore struct {
	db *sql.DB
}

func (store PostgresStore) CreateAccount() error {
	return nil
}

func newPostgressStoreOrFatal() *PostgresStore {
	user := os.Getenv("VISTA_VERSE_DB_USER")
	dbname := os.Getenv("VISTA_VERSE_DB_DBNAME")
	password := os.Getenv("VISTA_VERSE_DB_PASSWORD")

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln("Could not open SQL Connection: ", err.Error())
	}

	if err := createTables(db); err != nil {
		log.Fatalln("Could not create Tables: ", err.Error())
	}

	return &PostgresStore{
		db: db,
	}
}

func createTables(db *sql.DB) error {
	query := `create table if not exists account (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		cpf VARCHAR(11) NOT NULL,
		hash VARCHAR(72) NOT NULL,
		UNIQUE(cpf)
	);`

	_, err := db.Query(query)
	return err
}
