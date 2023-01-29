package nba_api

import "os"

type nbaApiConfig struct {
	nbaKey     string
	locale     string
	year       int
	seasonType string
}

func (config *nbaApiConfig) Init() {
	config.nbaKey = os.Getenv("sport_radar_api_key")
	config.locale = "en"
	config.year = 2022
	config.seasonType = "REG"
}
