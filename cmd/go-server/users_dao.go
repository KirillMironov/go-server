package main

import (
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

func insertUser(user *User, tx *sql.Tx) error {
	sqlStr := "INSERT INTO users (username, password, salt, email) VALUES ($1, $2, $3, $4)"

	_, err := tx.Exec(sqlStr, user.Username, user.Password, user.Salt, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func findUserByEmailAndPassword(email string, password string, db *sql.DB) (int64, error) {
	sqlStr := "SELECT * FROM users WHERE email = $1"
	var user User

	rows, err := db.Query(sqlStr, email)
	if err != nil {
		log.Printf("%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Email)
		if err != nil {
			log.Printf("%v", err)
		}

		if user.Password == hash(password + user.Salt) {
			rows.Close()
			return user.Id, nil
		}
	}

	return 0, sql.ErrNoRows
}

