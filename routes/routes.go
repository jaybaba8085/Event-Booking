package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouts(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEventById)
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register", CancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/users", getAllUsers)
	server.GET("/events/register", GetAllRegisteredEvents)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.DELETE("/events", deleteAllEvent)
}
