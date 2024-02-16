package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)
func RegisterRouts(server *gin.Engine) {
    server.GET("/events", getEvents)
    server.GET("/events/:id", getEvent)

    authenticated := server.Group("/")
    authenticated.Use(middlewares.Authenticate)
    authenticated.POST("/events", createEvent)
    authenticated.PUT("/events/:id", updateEvent)
    authenticated.DELETE("/events/:id", deleteEventById)


    server.DELETE("/events", deleteAllEvent)
    server.POST("/signup", signup)
    server.POST("/login", login)

    server.GET("/users", getAllUsers)
}
