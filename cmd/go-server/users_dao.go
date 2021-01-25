package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	_ "github.com/lib/pq"
)

type User struct {
	Id int64
	Username string
	Password string
	Salt string
	Email string
}

func insertUser(user *User, tx *sql.Tx) (int64, error) {
	sqlStr := "INSERT INTO users (username, password, salt, email) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int64

	err := tx.QueryRow(sqlStr, user.Username, user.Password, user.Salt, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func findUserByEmailAndPassword(email string, password string, db *sql.DB) (int64, error) {
	sqlStr := "SELECT * FROM users WHERE email = $1"
	var user User

	rows, err := db.Query(sqlStr, email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Email)
		if err != nil {
			return 0, err
		}

		hash := sha256.Sum256([]byte(password + user.Salt))

		if user.Password == hex.EncodeToString(hash[:]) {
			rows.Close()
			return user.Id, nil
		}
	}
	return 0, sql.ErrNoRows
}

