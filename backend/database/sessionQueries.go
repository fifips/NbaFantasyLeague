package database

import (
	"github.com/google/uuid"
)

func GetSessionById(id uuid.UUID) (*Session, error) {
	var s Session

	row := db.QueryRow("SELECT * FROM Session WHERE id = ?", id)
	if err := row.Scan(&s.Id, &s.UserId); err != nil {
		return nil, err
	}

	return &s, nil
}

func CreateOrUpdateSession(s Session) error {
	_, err := db.Exec("INSERT INTO Session (id, user_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE id = ?",
		s.Id, s.UserId, s.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSessionByUserId(id int) error {
	_, err := db.Exec("DELETE FROM Session WHERE user_id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
