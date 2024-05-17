package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	models "github.com/vinayaknolastname/our/services/common/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *sql.DB
}

func NewDB() (*Database, error) {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "ourdb", "postgres")

	db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		fmt.Println("Cannot Open Db %e", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(" Cannot Ping Db %e", err)
	}

	gormAutoMigrate(db, postgresqlDbInfo)

	return &Database{Db: db}, nil

}

func closeDB(d *Database) {
	d.Db.Close()

}

func gormAutoMigrate(d *sql.DB, dbstring string) {
	gormDB, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})

	if err != nil {

		log.Printf("gorm error", err)
	}

	err = gormDB.AutoMigrate(&models.UsersModel{}, &models.ChatsModel{}, &models.MessageModel{}, &models.ReactionOnChatModel{})

	if err != nil {

		log.Printf("autoMigrate error", err)
	}
	log.Printf("autoMigrate error", err)

}

// func (d *Database) getDB() *sql.DB {
// 	return d.db
// }
