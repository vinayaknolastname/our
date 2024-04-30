package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinayakbot/our/services/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *sql.DB
}

func ConnectDBFnc(dbstring string) *Storage {
	fmt.Println("DbString %e", dbstring)

	db, err := sql.Open("postgres", dbstring)

	if err != nil {
		fmt.Println("Cannot Open Db %e", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(" Cannot Ping Db %e", err)

	}

	gormDB, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})

	if err != nil {

		log.Printf("gorm error", err)
	}

	err = gormDB.AutoMigrate(&models.PgBasicModel{}, &models.PgAddressModel{}, &models.PgFeaturesModel{})

	if err != nil {
		log.Printf("autoMigrate error", err)
	}

	return &Storage{DB: db}

}
