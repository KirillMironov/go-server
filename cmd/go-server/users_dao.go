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
	Email string
	Password string
	Salt string
}

type UserData struct {
	Id int64
	Username string
	Email string
	Role string
}

func insertUser(user *User, tx *sql.Tx) (int64, error) {
	sqlStr := "INSERT INTO users (username, email, password, salt, role) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64

	err := tx.QueryRow(sqlStr, user.Username, user.Email, user.Password, user.Salt, "User").Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func findUserByEmailAndPassword(email string, password string, db *sql.DB) (int64, error) {
	sqlStr := "SELECT id, username, email, password, salt FROM users WHERE email = $1"
	var user User

	rows, err := db.Query(sqlStr, email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Salt)
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

func findUserById(id int64, db *sql.DB) error {
	sqlStr := "SELECT username, email, role FROM users WHERE id = $1"

	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&currentUser.Username, &currentUser.Email, &currentUser.Role)
		if err != nil {
			return err
		}

		currentUser.Id = id
		rows.Close()
		return nil
	}

	return sql.ErrNoRows
}

func updateUsername(username string, tx *sql.Tx) error {
	sqlStr := "UPDATE users SET username = $1 WHERE id = $2"

	_, err := tx.Exec(sqlStr, username, currentUser.Id)
	if err != nil {
		return err
	}

	return nil
}
