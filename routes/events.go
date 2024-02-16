package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	// Parse event ID from URL parameter
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Retrieve event by ID
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// Return event as JSON response
	context.JSON(http.StatusOK, event)
}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event}) //gin.h -> map
}

func updateEvent(context *gin.Context) {
    // Parse event ID from URL parameter
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
        return
    }

    // Retrieve event by ID
    existingEvent, err := models.GetEventByID(eventId)
    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
        return
    }

    // Parse updated event data from request body
    var updatedEvent models.Event
    if err := context.ShouldBindJSON(&updatedEvent); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body."})
        return
    }

    // Update event fields with new data
    existingEvent.Name = updatedEvent.Name
    existingEvent.Description = updatedEvent.Description
    existingEvent.Location = updatedEvent.Location
    existingEvent.DateTime = updatedEvent.DateTime
    existingEvent.UserID = updatedEvent.UserID

    // Save the updated event to the database
    if err := existingEvent.Update(); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
        return
    }

    // Return success response
    context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": existingEvent})
}

func deleteEvent(context *gin.Context) {
    // Delete all events
    if err := models.DeleteAllEvents(); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete events. Try again later."})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "All events deleted successfully"})
}

func deleteEventById(context *gin.Context) {
    // Parse event ID from URL parameter
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
        return
    }

    // Delete event by ID
    if err := models.DeleteEventByID(eventId); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later."})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}