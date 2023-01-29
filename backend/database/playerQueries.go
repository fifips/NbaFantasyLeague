package database

import (
	"log"
)

func CreatePlayer(player Player) error {
	_, err := db.Exec("INSERT INTO player (id) VALUES (?)", player.Id)

	return err
}

func GetPlayerById(id string) (Player, error) {
	var player Player
	err := db.QueryRow("SELECT * FROM player WHERE id = ?", id).Scan(
		&player.Id)
	return player, err
}

func UpdatePlayer(player Player) error {
	result, err := db.Exec("UPDATE player SET id = ? WHERE id = ?", player.Id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 1 {
		log.Printf("Updated Player's id from %s to %s", player.Id, player.Id)
	}

	return nil
}

func DeletePlayerById(id string) error {
	_, err := db.Exec("DELETE FROM player WHERE id = ?", id)
	return err
}
