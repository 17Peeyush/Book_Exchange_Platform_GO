package routes

import (
	middlewares "book-exchange.com/rest/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//server.POST("/events", createEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.GET("/allevents", middlewares.Authenticate, getUnfilteredFeed)
	server.GET("/events/:id", getFeed) //events/1, events/2
}