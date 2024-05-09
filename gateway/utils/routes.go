package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
)

// const routes (
//
//	string intial  ""
//
// )

func InitRouter(wshandler ws.Handler) {

	router := gin.Default()

	// router := gin.New(	)

	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5000"
		},
		MaxAge: 12 * time.Hour}))

	router.Use(gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{

		user := v1.Group("/user")
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
	}
	{
		ws := router.Group("/ws")
		ws.POST("createRoom", wshandler.CreateRoom)
		ws.GET("joinRoom/:roomId", wshandler.JoinRoom)
	}
	port := getConfig().GatewayPort

	portToRun := fmt.Sprint(":", port)

	// log.Panicf("err")

	err := router.Run(portToRun)

	if err != nil {
		log.Panicf("err", err)
	}

}
