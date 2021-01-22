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

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	user := &User{0, "Flanders", "666", "", "flanders@gmail.com"}
	user2 := &User{0, "Marge", "122", "", "marge@gmail.com"}

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
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = findUserByEmailAndPassword("Homer", "125", db)
	if err != nil {
		t.Fatal("Unable to find user")
	}

	_, err = findUserByEmailAndPassword("Homer", "Homer", db)
	if err == nil {
		t.Fatal("User was found using wrong password")
	}
}
