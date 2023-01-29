package common

import "time"

// This file contains all project global constant values

// Constants describing all project paths
const (
	ProjectPath = `E:\PracaDyplomowa\NbaFantasyLeague\backend`
	// Database paths

	SqlScriptsPath                      string = ProjectPath + `\database\sqlScripts`
	CreateTestDatabaseSqlScriptPath            = SqlScriptsPath + `\createTestDatabase.sql`
	InsertDataIntoDatabaseSqlScriptPath        = SqlScriptsPath + `\insertTestData.sql`
	DropTestDatabaseSqlScriptPath              = SqlScriptsPath + `\dropTestDatabase.sql`
)

// Constants for authentication tokens duration
const (
	AccessTokenExpiration  = time.Minute * 30
	RefreshTokenExpiration = time.Hour * 24 * 7
)

// ActivationCodeExpiration is a constant for account activation code duration
const ActivationCodeExpiration = time.Minute * 15

// Constants for intervals' durations for updating database
const (
	UpdateScheduleInterval  = time.Hour * 12
	UpdateStandingsInterval = time.Hour * 24
)

// Constant strings representing url nba api endpoints
const (
	//scheduleEndpoint string = "https://data.nba.net/prod/v1/2022/schedule.json"
	ScheduleEndpoint     string = "https://api.sportradar.us/nba/trial/v7/%s/games/%d/%s/schedule.json?api_key=%s"
	GameBoxScoreEndpoint string = "https://api.sportradar.us/nba/trial/v7/%s/games/%s/summary.json?api_key=%s"
	TeamInfoEndpoint     string = "https://api.sportradar.us/nba/trial/v7/%s/teams/%s/profile.json?api_key=%s"
	StandingsEndpoint    string = "https://api.sportradar.us/nba/trial/v7/%s/seasons/%d/%s/standings.json?api_key=%s"
)

// Frontend urls
const ActivationUrl string = "http://localhost:3000/user/activate?activation_code=%s"
