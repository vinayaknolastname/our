package main

import (
	"fmt"

	"github.com/vinayaknolastname/our/gateway/rtc/ws"
	"github.com/vinayaknolastname/our/gateway/utils"
)

func main() {
	// db, err := db.NewDB()

	// if err != nil {

	// }
	fmt.Println("sssss")

	hub := ws.NewHub()
	wshandler := ws.NewHandler(hub)

	utils.InitRouter(*wshandler)
	go hub.Run()

	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
