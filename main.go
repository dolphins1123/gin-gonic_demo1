package main

import (
	//api "gin-gonic/apis"
	api "gin-gonic/apis"
	db "gin-gonic/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/hello/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.JSON(200, gin.H{
	// 		"message": "hello " + name,
	// 	})
	// })
	// r.Run()

	defer db.SqlDB.Close()
	router := initRouter1()
	router.Run(":8001")
}

func initRouter1() *gin.Engine {
	router := gin.Default()

	router.GET("/", api.IndexApi)

	router.POST("/person", api.AddPersonApi)

	//router.GET("/list", api.GetPersons)

	// router.GET("/person/:id", GetPersonApi)

	// router.PUT("/person/:id", ModPersonApi)

	// router.DELETE("/person/:id", DelPersonApi)

	return router
}
