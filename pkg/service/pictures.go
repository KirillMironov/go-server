package service

import (
	"database/sql"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/repository"
)

type PicturesUsecase struct {
	picturesRepo    repository.PicturesRepository
}

func NewPicturesUsecase() repository.PicturesRepository {
	return &PicturesUsecase{
		picturesRepo: repository.PicturesRepository(PicturesUsecase{}),
	}
}

func (p PicturesUsecase) UploadPicture(picture *domain.Picture, tx *sql.Tx) error {
	sqlStr := "INSERT INTO pictures (item_id, picture) VALUES ($1, $2)"

	_, err := tx.Exec(sqlStr, picture.ItemId, picture.Picture)
	if err != nil {
		return err
	}

	return nil
}
