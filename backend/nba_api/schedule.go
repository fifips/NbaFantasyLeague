package nba_api

import (
	"backend/common"
	db "backend/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func parseGameInfo(gameInfo map[string]interface{}, newGame *db.Match) error {
	newGame.GameId = gameInfo["id"].(string)
	gameDate, err := time.Parse("2006-01-02T15:04:05Z", gameInfo["scheduled"].(string))
	if err != nil {
		return err
	}
	newGame.GameDate = gameDate
	if homeScore, ok := gameInfo["home_points"].(float64); ok {
		homeScoreConverted := int(homeScore)
		newGame.HomeScore = &homeScoreConverted
	}
	if awayScore, ok := gameInfo["away_points"].(float64); ok {
		awayScoreConverted := int(awayScore)
		newGame.AwayScore = &awayScoreConverted
	}

	if homeTeamInfo, ok := gameInfo["home"].(map[string]interface{}); ok != false {
		newGame.HomeTeamId = homeTeamInfo["id"].(string)
	}

	if awayTeamInfo, ok := gameInfo["away"].(map[string]interface{}); ok != false {
		newGame.AwayTeamId = awayTeamInfo["id"].(string)
	}

	return nil
}

func GetSchedule() ([]db.Match, error) {
	url := fmt.Sprintf(common.ScheduleEndpoint, config.locale, config.year, config.seasonType, config.nbaKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(body, &respJson); err != nil {
		return nil, err
	}

	var games []db.Match
	for _, gameInfo := range respJson["games"].([]any) {
		var newGame db.Match

		err := parseGameInfo(gameInfo.(map[string]any), &newGame)
		if err != nil {
			return nil, err
		}

		games = append(games, newGame)
	}
	return games, nil
}

func UpdateSchedule() error {
	newSchedule, err := GetSchedule()
	if err != nil {
		return err
	}
	for _, newGame := range newSchedule {
		if err := db.CreateOrUpdateScheduledGame(newGame); err != nil {
			return err
		}
	}

	return nil
}
