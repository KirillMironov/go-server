package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
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

		hash := sha256.Sum256([]byte(password + user.Salt))

		if user.Password == hex.EncodeToString(hash[:]) {
			rows.Close()
			return user.Id, nil
		}
	}

	return 0, sql.ErrNoRows
}

