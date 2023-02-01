package api

import (
	db "backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getSchedule retrieves a schedule from a database and returns it in JSON format.
func getSchedule(c *gin.Context) {
	schedule, err := db.GetSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": schedule})
}
