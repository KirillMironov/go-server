package usecase

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
)

var p = service.NewPicturesUsecase()

func UploadPicture(picture *domain.Picture) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = p.UploadPicture(picture, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
