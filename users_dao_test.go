package main

import (
	"database/sql"
	"fmt"
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
		t.Fatal("Unable to insert")
	}

	err = insert(user2, tx)
	if err != nil {
		t.Fatal("Unable to insert")
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

	users, err := findByEmailAndPassword("maggie@gmail.com", "125", tx)
	if err != nil {
		t.Fatal("Unable to find user")
	}

	fmt.Printf("Size = %d", users.Len())

	//Rollback tx
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		t.Fatal(err)
	}
}
