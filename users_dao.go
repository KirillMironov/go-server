package main

import (
	"container/list"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	Id int64
	Username string
	Password string
	Salt string
	Email string
}

func insert(user *User, tx *sql.Tx) error {
	sqlStr := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"

	_, err := tx.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func findByEmailAndPassword(email string, password string, tx *sql.Tx) (*list.List, error) {
	sqlStr := "SELECT * FROM users WHERE email = $1 AND password = $2"

	rows, err := tx.Query(sqlStr, email, password)
	if err != nil {
		log.Printf("%v", err)
	}
	defer rows.Close()

	users := list.New()

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Email); err != nil {
			log.Printf("%v", err)
		}

		users.PushBack(user)
	}

	return users, nil
}
