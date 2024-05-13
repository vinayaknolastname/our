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

	hub := ws.NewWsManager()
	wshandler := ws.NewHandler(hub)

	fmt.Println("sssss")

	go hub.RunWsManager()

	router.InitRouter(*wshandler)

	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
