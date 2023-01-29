package database

import (
	"log"
)

func GetSchedule() ([]Match, error) {
	var schedules []Match

	rows, err := db.Query("SELECT * FROM Schedule")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule Match
		if err := rows.Scan(
			&schedule.GameId, &schedule.GameDate, &schedule.HomeTeamId, &schedule.HomeScore, &schedule.AwayTeamId, &schedule.AwayScore); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func UpdateMatch(game Match) error {
	result, err := db.Exec("UPDATE Schedule "+
		"SET game_date = ?, home_team_id = ?, home_score = ?, away_team_id = ?, away_score = ? WHERE game_id = ?",
		game.GameDate.Format("2006-01-02 15:04:05"), game.HomeTeamId, game.HomeScore, game.AwayTeamId, game.AwayScore, game.GameId)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 1 {
		log.Printf("Updated scheduled game %s to: %s vs %s at %s",
			game.GameId, game.HomeTeamId, game.AwayTeamId, game.GameDate)
	}
	return nil
}

func DeleteMatch(game Match) error {
	result, err := db.Exec("DELETE FROM Schedule WHERE game_id = ?", game.GameId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 1 {
		log.Printf("Deleted scheduled game %s: %s vs %s, at %s",
			game.GameId, game.HomeTeamId, game.AwayTeamId, game.GameDate)
	}
	return nil
}

func CreateOrUpdateScheduledGame(game Match) error {
	_, err := db.Exec(`INSERT INTO Schedule 
    						(game_id, game_date, home_team_id, home_score, away_team_id, away_score) 
							VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE game_date = VALUES(game_date),
							home_team_id = VALUES(home_team_id), home_score = VALUES(home_score),
							away_team_id = VALUES(away_team_id), away_score = VALUES(away_score)`,
		game.GameId, game.GameDate.Format("2006-01-02 15:04:05"), game.HomeTeamId, game.HomeScore, game.AwayTeamId,
		game.AwayScore)

	return err
}
