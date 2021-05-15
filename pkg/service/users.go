package service

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/repository"
)

type UsersUsecase struct {
	usersRepo    repository.UsersRepository
}

func NewUsersUsecase() repository.UsersRepository {
	return &UsersUsecase{
		usersRepo: repository.UsersRepository(UsersUsecase{}),
	}
}

func (u UsersUsecase) CreateUser(user *domain.User, tx *sql.Tx) (int64, error) {
	sqlStr := "INSERT INTO users (username, email, password, salt, role) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64

	err := tx.QueryRow(sqlStr, user.Username, user.Email, user.Password, user.Salt, "User").Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u UsersUsecase) GetUserById(user *domain.User, db *sql.DB) error {
	sqlStr := "SELECT username, email, role FROM users WHERE id = $1"

	rows, err := db.Query(sqlStr, user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Email, &user.Role)
		if err != nil {
			return err
		}
		rows.Close()
		return nil
	}

	return sql.ErrNoRows
}

func (u UsersUsecase) GetUserByEmailAndPassword(user *domain.User, db *sql.DB) error {
	sqlStr := "SELECT id, username, email, password, salt FROM users WHERE email = $1"
	var tempUser domain.User

	rows, err := db.Query(sqlStr, user.Email)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tempUser.Id, &tempUser.Username, &tempUser.Email, &tempUser.Password, &tempUser.Salt)
		if err != nil {
			return err
		}

		hash := sha256.Sum256([]byte(user.Password + tempUser.Salt))

		if tempUser.Password == hex.EncodeToString(hash[:]) {
			rows.Close()
			user.Id = tempUser.Id
			user.Username = tempUser.Username
			user.Role = tempUser.Role
			user.Password = ""
			return nil
		}
	}

	return sql.ErrNoRows
}

func (u UsersUsecase) UpdateUsername(newUsername string, id int64, tx *sql.Tx) error {
	sqlStr := "UPDATE users SET username = $1 WHERE id = $2"

	_, err := tx.Exec(sqlStr, newUsername, id)
	if err != nil {
		return err
	}

	return nil
}
