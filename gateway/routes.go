package main

import "github.com/gin-gonic/gin"

// const routes (
// 	string intial  ""
// )
func InitRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{

		user := v1.Group("/user")
		user.GET("")
	}

}
