package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouts(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events", deleteEvent)
	server.DELETE("/events/:id", deleteEventById)
}
 