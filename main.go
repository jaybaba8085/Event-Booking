package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Go Rest-API Project")
	//  gin.Default() -> Default returns an Engine and Server instance with the Logger and Recovery middleware already attached.
	server := gin.Default()
	// server.Run() -> Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	// Note: this method will block the calling goroutine indefinitely unless an error happens.

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") //local host port  8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
	// context.JSON(http.StatusOK, gin.H{"message":"Hello! This is my first Go REST API."}) //gin.h -> map
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
	}

	event.ID = 1
	event.UserID = 1
	event.Saev()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}
