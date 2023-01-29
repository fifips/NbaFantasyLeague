package database

func GetTeamById(id string) (Team, error) {
	var team Team
	err := db.QueryRow("SELECT * FROM Team WHERE id = ?", id).Scan(
		&team.Id, &team.FullName, &team.Acronym, &team.Wins, &team.Losses)
	return team, err
}

func CreateOrUpdateTeam(team Team) error {
	_, err := db.Exec(`INSERT INTO Team (id, full_name, acronym, wins, losses) VALUES (?, ?, ?, ?, ?)
                                	ON DUPLICATE KEY UPDATE full_name = VALUES(full_name), acronym = VALUES(acronym),
                                	wins = VALUES(wins), losses = VALUES(losses)`,
		team.Id, team.FullName, team.Acronym, team.Wins, team.Losses)
	return err
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
