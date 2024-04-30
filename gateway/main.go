package main

import (
	"log"

	"github.com/vinayaknolastname/our/db"
)

func main() {
	_, err := db.NewDB()

	if err != nil {
		log.Println("db connection fail", err)
	}

}
