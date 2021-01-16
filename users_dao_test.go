package main

import (
	"database/sql"
	"testing"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatal(err)
	}

	//Begin tx
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	user := &User{0, "Flanders", "666", "", "flanders@gmail.com"}
	user2 := &User{0, "Marge", "122", "", "marge@gmail.com"}

	err = insert(user, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	err = insert(user2, tx)
	if err != nil {
		t.Fatal("Unable to insert user")
	}

	//Rollback tx
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		t.Fatal(err)
	}
}

func TestFindUserByEmailAndPassword(t *testing.T) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatal(err)
	}

	//Begin tx
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	users, err := findByEmailAndPassword("Homer", "125", tx)
	if err != nil || users.Len() == 0 {
		t.Fatal("Unable to find user")
	}

	//Rollback tx
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		t.Fatal(err)
	}
}

func TestCreateToken(t *testing.T) {
	user := &User{0, "Homer", "666", "", "homer@gmail.com"}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal(err)
	}
}
