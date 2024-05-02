package main

import (
	"github.com/vinayaknolastname/our/gateway/db"
)

func main() {
	db, err := db.NewDB()

	if err != nil {

	}
	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
