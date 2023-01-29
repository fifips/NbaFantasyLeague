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
	userRoutes.GET("/is_logged_in", isLoggedIn)

	userRoutes.Use(IsActiveUser)

	userRoutes.GET("/logout", logout)
}
