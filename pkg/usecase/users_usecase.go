package usecase

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	_ "github.com/lib/pq"
)

var u = service.NewUsersUsecase()

func CreateUser(user *domain.User) (int64, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	id, err := u.CreateUser(user, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	err = db.Close()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserById(user *domain.User) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	err = u.GetUserById(user, db)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmailAndPassword(user *domain.User) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	err = u.GetUserByEmailAndPassword(user, db)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func UpdateUsername(newUsername string, id int64) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = u.UpdateUsername(newUsername, id, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
