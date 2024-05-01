package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDB() (*Database, error) {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "ourDB", "ourDB")

	db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		fmt.Println("Cannot Open Db %e", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(" Cannot Ping Db %e", err)
	}

	return &Database{db: db}, nil

}

func closeDB(d *Database) {
	d.db.Close()

}

func (d *Database) getDB() *sql.DB {
	return d.db
}
