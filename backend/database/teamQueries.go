package database

import (
	"log"
)

func CreateTeam(team Team) error {
	_, err := db.Exec("INSERT INTO Team (id, full_name, acronym, wins, losses) VALUES (?, ?, ?, ?, ?)",
		team.Id, team.FullName, team.Acronym, team.Wins, team.Losses)

	return err
}

func GetTeamById(id string) (Team, error) {
	var team Team
	err := db.QueryRow("SELECT * FROM Team WHERE id = ?", id).Scan(
		&team.Id, &team.FullName, &team.Acronym, &team.Wins, &team.Losses)
	return team, err
}

func UpdateTeam(team Team) error {
	result, err := db.Exec("UPDATE Team SET wins = ?, losses = ? WHERE id = ?", team.Wins, team.Losses, team.Id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 1 {
		log.Printf("Updated Team %s to %d/%d (W/L)", team.FullName, team.Wins, team.Losses)
	}

	return nil
}

func DeleteTeam(id string) error {
	_, err := db.Exec("DELETE FROM Team WHERE id = ?", id)
	return err
}

func GetAllTeamIds() ([]string, error) {
	var teamIds []string
	rows, err := db.Query("SELECT (id) FROM Team")
	if err != nil {
		return teamIds, err
	}
	defer rows.Close()

	for rows.Next() {
		var teamId string

		err := rows.Scan(&teamId)
		if err != nil {
			return teamIds, err
		}

		teamIds = append(teamIds, teamId)
	}
	return teamIds, nil
}
