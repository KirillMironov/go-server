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

func insert(user *User, db *sql.DB) {
	sql:= "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
	_, err := db.Exec(sql, user.Username, user.Password, user.Email)
	if err != nil {
		panic(err)
	}
}