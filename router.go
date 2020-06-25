package main

import (
	api "gin-gonic/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", api.IndexApi)

	router.POST("/person", api.AddPersonApi)

	// router.GET("/persons", GetPersonsApi)

	// router.GET("/person/:id", GetPersonApi)

	// router.PUT("/person/:id", ModPersonApi)

	// router.DELETE("/person/:id", DelPersonApi)

	return router
}
