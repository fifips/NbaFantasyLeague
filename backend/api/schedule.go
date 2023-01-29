package api

import (
	db "NbaFantasyLeague/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getSchedule(c *gin.Context) {
	schedule, err := db.GetSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": schedule})
}
