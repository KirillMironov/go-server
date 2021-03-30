package test

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	_ "github.com/lib/pq"
	"testing"
)

var u = service.NewUsersUsecase()

func TestCreateUser(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	user := &domain.User{Username: "Flanders1", Email: "flanders1@gmail.com", Password: "666"}
	user2 := &domain.User{Username: "Marge11", Email: "marge11@gmail.com", Password: "122"}

	_, err = u.CreateUser(user, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	_, err = u.CreateUser(user2, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserById(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	user := &domain.User{Id: 1}
	user2 := &domain.User{Id: -1}

	err = u.GetUserById(user,  db)
	if err != nil {
		t.Fatal(err)
	}

	err = u.GetUserById(user2,  db)
	if err == nil {
		t.Fatal("User was found using wrong id")
	}
}

func TestGetUserByEmailAndPassword(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	user := &domain.User{Email: "bart@gmail.com", Password: "123"}
	user2 := &domain.User{Email: "Homer", Password: "Homer"}

	err = u.GetUserByEmailAndPassword(user, db)
	if err != nil {
		t.Fatal("Unable to find user")
	}

	err = u.GetUserByEmailAndPassword(user2, db)
	if err == nil {
		t.Fatal("User was found using wrong password")
	}
}

func TestUpdateUsername(t *testing.T)  {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	user := &domain.User{Id: 1}

	err = u.UpdateUsername("Maggie", user, tx)
	if err != nil {
		t.Fatal("Unable to change Username")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
