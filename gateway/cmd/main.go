package main

import (
	"fmt"

	mediaservice "github.com/vinayaknolastname/our/gateway/media_service"
	"github.com/vinayaknolastname/our/gateway/router"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
	"github.com/vinayaknolastname/our/gateway/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	// db, err := db.NewDB()

	if err != nil {
		utils.LogSomething("config", err, 0)
	}
	fmt.Println("sssss")

	hub := ws.NewWsManager()
	wshandler := ws.NewHandler(hub)

	utils.LogSomething("config", config, 0)

	mediaservice.NewMediaDB(config.Cloudinary)

	fmt.Println("sssss")

	go hub.RunWsManager()

	router.InitRouter(*wshandler)

	// if err != nil {
	// 	log.Println("db connection fail", err)
	// }

}
