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
	sqlStr := "INSERT INTO users (username, password, salt, email) VALUES ($1, $2, $3, $4)"

	_, err := tx.Exec(sqlStr, user.Username, user.Password, user.Salt, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func findByEmailAndPassword(email string, password string, tx *sql.Tx) (*list.List, error) {
	sqlStr := "SELECT * FROM users WHERE email = $1"

	rows, err := tx.Query(sqlStr, email)
	if err != nil {
		log.Printf("%v", err)
	}
	defer rows.Close()

	users := list.New()

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Email)
		if err != nil {
			log.Printf("%v", err)
		}

		if user.Password == hash(password + user.Salt) {
			users.PushBack(user)
		}
	}

	return users, nil
}
