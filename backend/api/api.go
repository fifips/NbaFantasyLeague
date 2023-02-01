package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/schedule", getSchedule)

	router.POST("/login", login)
	router.POST("/register", register)

	userRoutes := router.Group("/user")
	userRoutes.Use(Authenticate)

	userRoutes.POST("/activate", activateUser)
	userRoutes.GET("/logged_in", isLoggedIn)
	userRoutes.GET("/logout", logout)

	userRoutes.GET("/leagues", getAllLeagueIds)
	userRoutes.Use(IsActiveUser)

}
