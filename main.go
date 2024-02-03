package main

import (
	"book-exchange.com/rest/db"
	"book-exchange.com/rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	// server.GET("/ping", func(c *gin.Context){
	// 	c.JSON(http.StatusOK, gin.H{"Message":"Running server..."})
	// })
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
