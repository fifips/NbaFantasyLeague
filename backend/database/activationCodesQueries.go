package database

// GetActivationCodeByUserId returns ActivationCode struct assigned to User with given userId
func GetActivationCodeByUserId(userId int) (ActivationCode, error) {
	var aC ActivationCode

	row := db.QueryRow("SELECT * FROM activation_code WHERE user_id = ?", userId)
	if err := row.Scan(&aC.UserId, &aC.Code, &aC.Expires); err != nil {
		return aC, err
	}

	return aC, nil
}

// CreateOrUpdateActivationCode creates new ActivationCode row in database. If there already exist one
// it updates its code and expires columns.
func CreateOrUpdateActivationCode(aC ActivationCode) error {
	_, err := db.Exec(`INSERT INTO activation_code (code, user_id, expires) 
							VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE code = ?, expires = ?`,
		aC.Code, aC.UserId, aC.Expires, aC.Code, aC.Expires)
	if err != nil {
		return err
	}
	return nil
}

// DeleteActivationCodeByUserId deletes ActivationCode row assigned to User with given userId from database
func DeleteActivationCodeByUserId(userId int) error {
	_, err := db.Exec("DELETE FROM activation_code WHERE user_id = ?", userId)
	if err != nil {
		return err
	}
	return nil
}
