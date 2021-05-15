package test

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	"testing"
)

var p = service.NewPicturesUsecase()

func TestUploadPicture(t *testing.T) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		t.Fatal(err)
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	picture := &domain.Picture{Id: 1, ItemId: 1, Picture: "[255, 216, 255]"}

	err = p.UploadPicture(picture, tx)
	if err != nil {
		t.Fatal("Unable to upload a picture")
	}

	// Rollback transaction
	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
