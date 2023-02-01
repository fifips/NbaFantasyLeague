package database

func GetAllLeagueIds() ([]int, error) {
	var ids []int
	rows, err := db.Query("SELECT (id) FROM league")
	if err != nil {
		return ids, err
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	return ids, err
}

func CreateLeague(league League) error {
	_, err := db.Exec("INSERT INTO League (owner_id, name, pts_ratio, reb_ratio, ast_ratio) VALUES (?, ?, ?, ?, ?)",
		league.OwnerId, league.Name, league.PtsRatio, league.RebRatio, league.AstRatio)

	return err
}

func GetLeagueById(id int) (League, error) {
	var league League
	err := db.QueryRow("SELECT * FROM league WHERE id = ?", id).Scan(
		&league.Id, &league.OwnerId, &league.Name, &league.PtsRatio, &league.RebRatio, &league.AstRatio)

	return league, err
}

func DeleteLeagueById(id int) error {
	_, err := db.Exec("DELETE FROM league WHERE id = ?", id)
	return err
}
