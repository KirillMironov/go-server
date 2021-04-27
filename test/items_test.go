package test

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	"testing"
	"time"
)

var i = service.NewItemsUsecase()

func TestCreateItem(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	item := &domain.Item{Title: "4K HDMI Cable", Description: "Connects Blu-ray players, Fire TV, Apple TV, PS4",
		Price: 8.86, Attributes: `{"Connector Type": "HDMI","Color": "Black","Brand": "Amazon Basics"}`,
		Picture: "[255, 104, 205]", StatusId: 1, CreatedAt: time.Now()}

	_, err = i.CreateItem(item, tx)
	if err != nil {
		t.Fatal("Unable to add item")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetItemById(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = i.GetItemById(1, db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = i.GetItemById(-1, db)
	if err == nil {
		t.Fatal("Item was found using wrong id")
	}
}

func TestGetItemsByTitleOrDescription(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = i.GetItemsByTitleOrDescription("Fire", db)
	if err != nil {
		t.Fatal("Unable to find item")
	}
}
