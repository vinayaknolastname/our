package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
)

// const routes (
//
//	string intial  ""
//
// )

func InitRouter(wshandler ws.Handler) {

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{

		user := v1.Group("/user")
		user.GET("")
	}
	{

		ws := v1.Group("/ws")
		ws.POST("createRoom", wshandler.CreateRoom)
	}

	router.Run()

}
