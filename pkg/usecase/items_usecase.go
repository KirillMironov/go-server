package usecase

import (
	"database/sql"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
)

var (
	i = service.NewItemsUsecase()
)

func CreateItem(item *domain.Item) (int64, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	id, err := i.CreateItem(item, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_ = tx.Commit()
	return id, err
}

func GetItemsByTitleOrDescription(query string) ([]*domain.Item, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return nil, err
	}

	items, err := i.GetItemsByTitleOrDescription(query, db)
	if err != nil {
		return nil, err
	}

	return items, nil
}
