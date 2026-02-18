package api

import (
	"base/service"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ApplyToEvent(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, _ := strconv.ParseUint(eventIDStr, 10, 32)

	userID := c.MustGet("user_id").(uint)

	reg, err := service.RegisterToEvent(userID, uint(eventID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reg)
}

func GetMyRegistrations(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	regs, err := service.GetMyEvents(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regs)
}

// GET /api/events/:id/participants
func GetParticipants(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDStr, 10, 32)
	if err != nil {
		fmt.Println("Ошибка парсинга ID:", err) // Проверь, не пустая ли строка
	}
	userID := c.MustGet("user_id").(uint)
	

	regs, err := service.GetParticipants(userID, uint(eventID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regs)
}

// PATCH /api/registrations/:id/status
func UpdateStatus(c *gin.Context) {
	regIDStr := c.Param("id")
	regID, _ := strconv.ParseUint(regIDStr, 10, 32)
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	reg, err := service.UpdateApplicationStatus(userID, uint(regID), input.Status)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reg)
}