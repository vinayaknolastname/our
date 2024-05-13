package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
	"github.com/vinayaknolastname/our/gateway/utils"
)

// const routes (
//
//	string intial  ""
//
// )

func InitRouter(wshandler ws.Handler) {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			return origin == "http://localhost:5000"
		},
		MaxAge: 12 * time.Hour}))

	// router := gin.New(	)

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{

		user := v1.Group("/user")
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
		user.POST("/createUser", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.CreateUser)
		user.GET("/getUserAndChats/:userId", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.GetUserAndChats)
		user.POST("/startChat", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.StartChat)
		// user.GET("/getUserChats/:chatId", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.CreateUser)
		// user.GET("/join/:chatId", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.CreateUser)

	}
	{
		ws := router.Group("/ws")
		ws.POST("createRoom", wshandler.CreateRoom)
		ws.GET("startChat/:index", wshandler.StartChat)
		ws.GET("getRooms", wshandler.GetRooms)

	}
	port := utils.GetConfig().GatewayPort

	portToRun := fmt.Sprint(":", port)

	// log.Panicf("err")

	err := router.Run(portToRun)

	if err != nil {
		log.Panicf("err", err)
	}

}
