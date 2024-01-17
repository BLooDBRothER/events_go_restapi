package routes

import (
	"net/http"
	"strconv"

	"example.com/events/controller"
	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := controller.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := controller.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No event present"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserId = 1

	insertedId, err := controller.CreateEvent(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": insertedId})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := controller.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No event present"})
		return
	}

	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = controller.UpdateEvent(event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated successfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = controller.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No event present"})
		return
	}

	err = controller.DeleteEvent(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted successfully"})
}
