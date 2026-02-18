package api

import (
	"base/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApplyToEvent(c *gin.Context) {
	eventIDstr := c.Param("id")
	eventID, _ := strconv.ParseUint(eventIDstr, 10, 32)

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