package database

import (
	. "backend/common"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func GetUserById(id int) (*User, error) {
	var u User

	row := db.QueryRow("SELECT * FROM User WHERE id = ?", id)
	if err := row.Scan(&u.Id, &u.Email, &u.Password, &u.IsActive); err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUserByEmail(email string) (*User, error) {
	var u User

	row := db.QueryRow("SELECT * FROM User WHERE email = ?", email)
	if err := row.Scan(&u.Id, &u.Email, &u.Password, &u.IsActive); err != nil {
		return nil, err
	}

	return &u, nil
}

func ActivateUserById(userId int) error {
	_, err := db.Exec("UPDATE User SET is_active = ? WHERE id = ?", true, userId)
	return err
}

func CreateUser(u *User) error {
	result, err := db.Exec("INSERT INTO User (email, password) VALUES (?, ?)", u.Email, u.Password)
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return CustomError{Message: "This email is already taken."}
			}
			return err
		}
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(insertedId)

	return nil
}

func DeleteUserById(id int) error {
	result, err := db.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("rowsAffected == 0")
	}
	return nil
}
