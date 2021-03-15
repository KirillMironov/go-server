package main

import (
	"database/sql"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"testing"
)

func TestInsertPicture(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	picture := &Picture{Id: 1, ItemId: 1, Picture: "[255, 216, 255]"}

	err = insertPicture(picture, tx)
	if err != nil {
		t.Fatal("Unable to add picture")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
