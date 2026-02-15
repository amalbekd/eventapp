package api

import (
	"base/models"
	"base/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	userId, _ := c.Get("user_id")
	event.OrganizerID = userId.(uint)

	if err := service.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func GetEvents(c *gin.Context) {
	events, err := service.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")

	eventID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

	event, err := service.GetEventByID(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}