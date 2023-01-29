package main

import (
	"NbaFantasyLeague/api"
	"NbaFantasyLeague/database"
	"NbaFantasyLeague/nba_api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer database.DisconnectDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	api.RegisterRoutes(router)

	go func() {
		nba_api.UpdateDatabase()
	}()

	router.Run("localhost:8080")
}
