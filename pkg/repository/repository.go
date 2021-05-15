package repository

import (
	"database/sql"
	"github.com/KirillMironov/go-server/domain"
)

type UsersRepository interface {
	CreateUser(user *domain.User, tx *sql.Tx) (int64, error)
	GetUserById(user *domain.User, db *sql.DB) error
	GetUserByEmailAndPassword(user *domain.User, db *sql.DB) error
	UpdateUsername(newUsername string, id int64, tx *sql.Tx) error
}

type ItemsRepository interface {
	CreateItem(item *domain.Item, tx *sql.Tx) (int64, error)
	GetItemById(id int64, db *sql.DB) (domain.Item ,error)
	GetItemsByTitleOrDescription(query string, db *sql.DB) ([]*domain.Item, error)
}

type SecurityRepository interface {
	GenerateHashedPasswordAndSalt(password string) (string, string)
	GenerateAuthToken(user *domain.User) (string, error)
	VerifyAuthToken(token string) (int64, error)
}

type PicturesRepository interface {
	UploadPicture(picture *domain.Picture, tx *sql.Tx) error
}
