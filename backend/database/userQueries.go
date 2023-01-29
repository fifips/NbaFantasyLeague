package database

func GetUserById(id int) (*User, error) {
	var u User

	row := db.QueryRow("SELECT * FROM User WHERE id = ?", id)
	if err := row.Scan(&u.Id, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUserByEmail(email string) (*User, error) {
	var u User

	row := db.QueryRow("SELECT * FROM User WHERE email = ?", email)
	if err := row.Scan(&u.Id, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return &u, nil
}

func CreateUser(u User) error {
	result, err := db.Exec("INSERT INTO User (email, password) VALUES (?, ?)", u.Email, u.Password)
	if err != nil {
		return err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(insertedId)

	return nil
}

func DeleteUserById(id int) error {
	_, err := db.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
