package main

import (
	"database/sql"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"testing"
	"time"
)

func TestInsertItem(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	item := &Item{0, "4K HDMI Cable", "Connects Blu-ray players, Fire TV, Apple TV, PS4",
		8.86, `{"Connector Type": "HDMI","Color": "Black","Brand": "Amazon Basics"}`, "[255, 104, 205]", 1, "", time.Now()}

	_, err = insertItem(item, tx)
	if err != nil {
		t.Fatal("Unable to add item")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindItemById(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = findItemById(4, db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = findItemById(-1, db)
	if err == nil {
		t.Fatal("Item was found using wrong id")
	}
}

func TestFindItemsByTitleOrDescription(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = findItemsByTitleOrDescription("Fire", db)
	if err != nil {
		t.Fatal("Unable to find item")
	}
}
