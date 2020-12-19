package main

import (
	"database/sql"
	"log"
	"testing"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//Begin tx
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	user := &User{"Maggie", "125", "maggie@gmail.com"}
	user2 := &User{"Lisa", "123", "lisa@gmail.com"}

	err = insert(user, tx)
	if err != nil {
		t.Fatal("Unable to insert")
	}

	err = insert(user2, tx)
	if err != nil {
		t.Fatal("Unable to insert")
	}

	//Rollback tx
	rollbackErr := tx.Commit()
	if rollbackErr != nil {
		log.Fatal(err)
	}
}
