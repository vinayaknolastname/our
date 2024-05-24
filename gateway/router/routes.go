package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpcHandlers "github.com/vinayaknolastname/our/gateway/grpc"
	mediaservice "github.com/vinayaknolastname/our/gateway/media_service"
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

	router.MaxMultipartMemory = 200 << 20 // 8 MB

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{

		user := v1.Group("/user")
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
		user.POST("/createUser", grpcHandlers.ConnectUserServiceGrpcMiddleWare, grpcHandlers.CreateUser)
		user.POST("/getUserAndChats/:userId", grpcHandlers.ConnectUserServiceGrpcMiddleWare, grpcHandlers.GetUserAndChats)
		user.POST("/startChat", grpcHandlers.ConnectUserServiceGrpcMiddleWare, grpcHandlers.StartChat)
		user.POST("/getAllChats", grpcHandlers.ConnectUserServiceGrpcMiddleWare, grpcHandlers.GetAllChats)
		user.POST("/uploadFile/:userId", grpcHandlers.ConnectUserServiceGrpcMiddleWare, mediaservice.FileUpload)

		// user.GET("/getUserChats/:chatId", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.CreateUser)
		// user.GET("/join/:chatId", grpcHandlers.ConnectUserServiceGrpc, grpcHandlers.CreateUser)

	}
	{

		video := v1.Group("/video")
		video.POST("/startVideo", grpcHandlers.ConnectUserServiceGrpcMiddleWare, grpcHandlers.CreateUser)

	}
	{
		ws := router.Group("/ws")
		ws.POST("createRoom", wshandler.CreateRoom)
		ws.GET("startChat/:chatId", wshandler.StartChat)
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
