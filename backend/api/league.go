package api

import (
	db "backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllLeagueIds(c *gin.Context) {
	leagueIds, err := db.GetAllLeagueIds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leagueIds)
}
