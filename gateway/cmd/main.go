package main

import (
	"fmt"

	"github.com/vinayaknolastname/our/gateway/router"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
)

func main() {
	// db, err := db.NewDB()

	// if err != nil {

	// }
	fmt.Println("sssss")

	hub := ws.NewHub()
	wshandler := ws.NewHandler(hub)

	router.InitRouter(*wshandler)
	go hub.Run()

	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
