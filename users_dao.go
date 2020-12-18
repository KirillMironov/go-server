package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Username string
	Password string
	Email    string
}

func insert(user *User, tx *sql.Tx) error {
	sqlStr := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"

	_, err := tx.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}
