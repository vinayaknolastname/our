package main

import (
	"github.com/vinayaknolastname/our/gateway/db"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
)

func main() {
	db, err := db.NewDB()

	if err != nil {

	}

	hub := ws.NewHub()
	wshandler := ws.NewHandler(hub)

	go hub.Run()
	InitRouter(*wshandler)
	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
