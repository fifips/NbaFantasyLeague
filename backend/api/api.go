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

	userRoutes.GET("/logout", logout)
	userRoutes.GET("/is_logged_in", isLoggedIn)
}
