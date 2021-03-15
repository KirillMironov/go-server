package main

import (
	"database/sql"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"testing"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	user := &User{0, "Flanders1", "flanders1@gmail.com", "666", ""}
	user2 := &User{0, "Marge11", "marge11@gmail.com", "122", ""}

	_, err = insertUser(user, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	_, err = insertUser(user2, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindUserByEmailAndPassword(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = findUserByEmailAndPassword("bart@gmail.com", "123", db)
	if err != nil {
		t.Fatal("Unable to find user")
	}

	_, err = findUserByEmailAndPassword("Homer", "Homer", db)
	if err == nil {
		t.Fatal("User was found using wrong password")
	}
}

func TestFindUserById(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	err = findUserById(1,  db)
	if err != nil {
		t.Fatal(err)
	}

	err = findUserById(-1,  db)
	if err == nil {
		t.Fatal("User was found using wrong id")
	}
}

func TestChangeUsername(t *testing.T)  {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	currentUser.Id = 1

	err = updateUsername("Maggie", tx)
	if err != nil {
		t.Fatal("Unable to change Username")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
