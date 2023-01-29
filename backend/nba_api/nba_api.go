package nba_api

import (
	"backend/common"
	db "backend/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var config nbaApiConfig

func init() {
	config.Init()
}

func runUpdate(timeInterval time.Duration, updateFunc func() error) {
	err := updateFunc()
	if err != nil {
		log.Printf("Error trying to run %s: %#v", common.GetFunctionName(updateFunc), err.Error())
	}
	time.Sleep(timeInterval)
}

func UpdateDatabase() {
	go func() {
		runUpdate(common.UpdateScheduleInterval, UpdateSchedule)
	}()
	go func() {
		runUpdate(common.UpdateStandingsInterval, UpdateStandings)
	}()
}

func parsePlayerPerformance(performanceInfo map[string]any, playerPerformance *db.PlayerPerformance) error {
	playerPerformance.PlayerId = performanceInfo["id"].(string)

	performanceStats := performanceInfo["statistics"].(map[string]any)
	minutesPlayed, err := time.ParseDuration(performanceStats["minutes"].(string))
	if err != nil {
		return err
	}
	playerPerformance.Minutes = minutesPlayed
	playerPerformance.Points = int(performanceStats["points"].(float64))
	playerPerformance.Assists = int(performanceStats["assists"].(float64))
	playerPerformance.Rebounds = int(performanceStats["rebounds"].(float64))
	playerPerformance.Turnovers = int(performanceStats["turnovers"].(float64))
	playerPerformance.Steals = int(performanceStats["steals"].(float64))
	playerPerformance.Blocks = int(performanceStats["blocks"].(float64))
	playerPerformance.Fouls = int(performanceStats["personal_fauls"].(float64))

	return nil
}

func GetBoxScoreByGameId(gameId string) (db.GameBoxScore, error) {
	var resultBoxScore db.GameBoxScore

	url := fmt.Sprintf(common.GameBoxScoreEndpoint, config.locale, gameId, config.nbaKey)
	resp, err := http.Get(url)
	if err != nil {
		return resultBoxScore, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resultBoxScore, err
	}

	var respJson map[string]any
	if err := json.Unmarshal(body, &respJson); err != nil {
		return resultBoxScore, err
	}

	awayTeamBoxScore := respJson["away"].(map[string]any)
	awayTeamPlayerPerformances := awayTeamBoxScore["players"].([]map[string]any)

	for _, performance := range awayTeamPlayerPerformances {
		var newPlayerPerformance db.PlayerPerformance
		newPlayerPerformance.GameId = gameId

		err := parsePlayerPerformance(performance, &newPlayerPerformance)
		if err != nil {
			return resultBoxScore, err
		}

		resultBoxScore.AwayPlayersPerformance = append(resultBoxScore.AwayPlayersPerformance, newPlayerPerformance)
	}

	homeTeamBoxScore := respJson["away"].(map[string]any)
	homeTeamPlayerPerformances := homeTeamBoxScore["players"].([]map[string]any)

	for _, performance := range homeTeamPlayerPerformances {
		var newPlayerPerformance db.PlayerPerformance
		newPlayerPerformance.GameId = gameId

		err := parsePlayerPerformance(performance, &newPlayerPerformance)
		if err != nil {
			return resultBoxScore, err
		}

		resultBoxScore.HomePlayersPerformance = append(resultBoxScore.HomePlayersPerformance, newPlayerPerformance)
	}

	return resultBoxScore, nil
}
