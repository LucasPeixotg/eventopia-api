package db

import (
	"database/sql"
	"fmt"
	"log"
	c "vistaverse/src/common"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func newPostgressStoreOrFatal() *PostgresStore {
	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", db_user, db_dbname, db_password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln("Could not open SQL Connection: ", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error while Pinging DB: ", err.Error())
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
	);
	create table if not exists event (
		id SERIAL PRIMARY KEY,
		owner_id INT NOT NULL,
		name VARCHAR(50) NOT NULL,
		datetime TIMESTAMP,
		location VARCHAR(50),
		description TEXT,
		FOREIGN KEY(owner_id) REFERENCES account(id)
	)
	`

	_, err := db.Query(query)
	return err
}

/*
Account and auth related
*/
func (store PostgresStore) CreateAccount(req *c.CreateAccountRequest) (*c.Account, error) {
	// TODO: validate cpf, name and password

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `insert into account (name, cpf, hash) values ($1, $2, $3) returning id, cpf, name;`
	rows, err := store.db.Query(query, req.Name, req.Cpf, string(hash))
	if err != nil {
		return nil, err
	}

	// calling row.Next is needed even to
	// scan the first row.
	rows.Next()

	acc := &c.Account{}
	err = rows.Scan(&acc.ID, &acc.CPF, &acc.Name)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (store PostgresStore) Login(req *c.LoginRequest) (*c.Account, error) {
	query := `select id, name, cpf, hash from account where cpf=$1;`
	rows, err := store.db.Query(query, req.Cpf)
	if err != nil {
		return nil, err
	}

	acc := &c.Account{}
	var hash string

	rows.Next()
	if err := rows.Scan(&acc.ID, &acc.Name, &acc.CPF, &hash); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		return nil, err
	}

	return acc, nil
}

/*
Event related
*/
func (store PostgresStore) CreateEvent(req *c.CreateEventRequest, user_id int) (*c.Event, error) {
	query := `insert into 
		event (name, description, location, datetime, owner_id) 
		values ($1, $2, $3, $4, $5) 
		returning id, owner_id, name, datetime, location, description;`
	rows, err := store.db.Query(query, req.Name, req.Description, req.Location, req.DateTime, user_id)
	if err != nil {
		return nil, err
	}

	ev := &c.Event{}

	rows.Next()
	if err := rows.Scan(&ev.ID, &ev.OwnerID, &ev.Name, &ev.DateTime, &ev.Location, &ev.Description); err != nil {
		return nil, err
	}

	return ev, nil
}

func (store PostgresStore) GetEvents(page int) {
	
}
