package main

import "github.com/vinayaknolastname/our/gateway/db"

func main() {
	_, err := db.NewDB()

	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
