package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// RegisterForEvent registers a user for a specific event.
func RegisterForEvent(context *gin.Context) {
	// Extract user ID and event ID from the request
	userID := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Retrieve event by ID
	// existingEvent, err := models.GetEventByID(eventId)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found."})
	// 	return
	// }

	// //Only the owner of an event can modify it
	// if existingEvent.UserID != int(userID) {
	// 	context.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "You are not authorized to perform this action.",
	// 	})
	// 	return
	// }

	// Create a RegisteredEvent instance with user and event IDs
	registeredEvent := &models.RegisteredEvent{
		UserID:  int(userID),
		EventID: int(eventId),
	}

	// Perform registration in the database
	err = registeredEvent.RegisterUserForEvent(int(userID), int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register the user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

// CancelRegistration cancels a user's registration for a specific event.
func CancelRegistration(context *gin.Context) {
	// Extract registration ID from the request
	registrationID := context.PostForm("registration_id")

	// Convert registration ID to an integer
	registrationIDInt, err := strconv.Atoi(registrationID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration ID"})
		return
	}

	// Perform cancellation in the database
	err = models.CancelRegistration(registrationIDInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully"})
}

// GetAllRegisteredEventsHandler returns all registered events.
func GetAllRegisteredEvents(context *gin.Context) {
	// Fetch all registered events from the database
	registeredEvents, err := models.Get_AllRegisteredEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registered events"})
		return
	}

	context.JSON(http.StatusOK, registeredEvents)
}
