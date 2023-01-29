package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Match struct {
	GameId     string    `json:"game_id"`
	GameDate   time.Time `json:"game_date"`
	HomeTeamId string    `json:"home_team_id"`
	HomeScore  *int      `json:"home_score"`
	AwayTeamId string    `json:"away_team_id"`
	AwayScore  *int      `json:"away_score"`
}

func (match *Match) MarshalJSON() ([]byte, error) {
	homeTeam, err := GetTeamById(match.HomeTeamId)
	if err != nil {
		return nil, err
	}
	awayTeam, err := GetTeamById(match.AwayTeamId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&struct {
		Id        string    `json:"id"`
		Date      time.Time `json:"date"`
		HomeTeam  Team      `json:"home_team"`
		HomeScore *int      `json:"home_score"`
		AwayTeam  Team      `json:"away_team"`
		AwayScore *int      `json:"away_score"`
	}{
		Id:        match.GameId,
		Date:      match.GameDate,
		HomeTeam:  homeTeam,
		HomeScore: match.HomeScore,
		AwayTeam:  awayTeam,
		AwayScore: match.AwayScore,
	})
}

type Team struct {
	Id        string   `json:"team_id"`
	FullName  string   `json:"full_name"`
	Acronym   string   `json:"acronym"`
	Wins      int      `json:"wins"`
	Losses    int      `json:"losses"`
	PlayerIds []string `json:"player_ids"`
}

func (team *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id       string `json:"id"`
		FullName string `json:"full_name"`
		Acronym  string `json:"acronym"`
		Wins     int    `json:"wins"`
		Losses   int    `json:"losses"`
	}{
		Id:       team.Id,
		FullName: team.FullName,
		Acronym:  team.Acronym,
		Wins:     team.Wins,
		Losses:   team.Losses,
	})

}

type GameBoxScore struct {
	GameId                 string `json:"game_id"`
	AwayPlayersPerformance []PlayerPerformance
	HomePlayersPerformance []PlayerPerformance
}

type Player struct {
	Id string `json:"id"`
}

type PlayerPerformance struct {
	GameId    string        `json:"game_id"`
	PlayerId  string        `json:"player_id"`
	Minutes   time.Duration `json:"minutes"`
	Points    int           `json:"points"`
	Assists   int           `json:"assists"`
	Rebounds  int           `json:"rebounds"`
	Turnovers int           `json:"turnovers"`
	Steals    int           `json:"steals"`
	Blocks    int           `json:"blocks"`
	Fouls     int           `json:"fouls"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	IsActive bool   `json:"is_active"`
}

//
//func (u *User) UnmarshalJSON(data []byte) error {
//	var values map[string]any
//	if err := json.Unmarshal(data, &values); err != nil {
//		return err
//	}
//
//	id, ok := values["id"].(float64)
//	if ok == true {
//		u.Id = int(id)
//	}
//
//	u.Email = values["email"].(string)
//	u.Password = []byte(values["password"].(string))
//
//	return nil
//}

type Session struct {
	Id     uuid.UUID `json:"id"`
	UserId int       `json:"user_id"`
}

type ActivationCode struct {
	Code    uuid.UUID `json:"code"`
	UserId  int       `json:"user_id"`
	Expires time.Time `json:"expires"`
}
