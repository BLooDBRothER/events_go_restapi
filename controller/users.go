package controller

import (
	"example.com/events/db"
	"example.com/events/models"
	"example.com/events/utils"
)

func CreateUser(user *models.User) (int64, error) {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func ValidateUser(user *models.User) (bool, error) {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, user.Email)

	var reterivedPassword string
	err := row.Scan(&user.Id, &reterivedPassword)

	if err != nil {
		return false, err
	}

	return utils.ComparePasswordHash(user.Password, reterivedPassword), nil
}
