package nba_api

import (
	"NbaFantasyLeague/common"
	"NbaFantasyLeague/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func parseTeamInfo(teamInfo map[string]any, team *database.Team) error {
	team.Id = teamInfo["id"].(string)
	team.FullName = fmt.Sprintf("%s %s", teamInfo["market"].(string), teamInfo["name"].(string))
	team.Acronym = teamInfo["alias"].(string)

	for _, playerInfo := range teamInfo["players"].([]map[string]any) {
		team.PlayerIds = append(team.PlayerIds, playerInfo["id"].(string))
	}

	return nil
}

func GetTeamInfo(teamId string) (database.Team, error) {
	var team database.Team

	url := fmt.Sprintf(common.TeamInfoEndpoint, config.locale, teamId, config.nbaKey)
	resp, err := http.Get(url)
	if err != nil {
		return team, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return team, err
	}

	var respJson map[string]any
	if err := json.Unmarshal(body, &respJson); err != nil {
		return team, err
	}

	err = parseTeamInfo(respJson, &team)
	if err != nil {
		return team, err
	}

	return team, nil
}

func parseStandings(leagueInfo map[string]any, teams *[]database.Team) {
	for _, conference := range leagueInfo["conferences"].([]any) {
		for _, division := range conference.(map[string]any)["divisions"].([]any) {
			for _, teamInfo := range division.(map[string]any)["teams"].([]any) {
				var team database.Team

				teamInfoConverted := teamInfo.(map[string]any)

				team.Id = teamInfoConverted["id"].(string)
				team.FullName = fmt.Sprintf("%s %s", teamInfoConverted["market"].(string), teamInfoConverted["name"].(string))
				team.Wins = int(teamInfoConverted["wins"].(float64))
				team.Losses = int(teamInfoConverted["losses"].(float64))

				*teams = append(*teams, team)
			}
		}
	}

}

func GetStandings() ([]database.Team, error) {
	var teams []database.Team

	client := &http.Client{}
	url := fmt.Sprintf(common.StandingsEndpoint, config.locale, config.year, config.seasonType, config.nbaKey)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return teams, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return teams, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return teams, err
	}

	var respJson map[string]any
	if err := json.Unmarshal(body, &respJson); err != nil {
		return teams, err
	}

	parseStandings(respJson, &teams)

	return teams, nil
}

func UpdateStandings() error {
	updatedStandings, err := GetStandings()
	if err != nil {
		return err
	}
	for _, updatedTeam := range updatedStandings {
		err = database.UpdateTeam(updatedTeam)
		if err != nil {
			return err
		}
	}

	return nil
}
