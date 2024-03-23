package models

import (
	"database/sql"
	"test-echo/db"
	"test-echo/helpers"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string

	conn := db.CreateConn()

	sqlQuery := "SELECT id,email,password FROM users WHERE email = ?"

	err := conn.QueryRow(sqlQuery, email).Scan(&obj.Id, &obj.Email, &pwd)

	if err == sql.ErrNoRows {
		return false, err
	}

	if err != nil {
		return false, err
	}

	match, err := helpers.ComparePassword(password, pwd)
	if !match {
		return false, err
	}

	return true, nil
}
