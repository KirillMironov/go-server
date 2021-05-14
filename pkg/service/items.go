package service

import (
	"database/sql"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/repository"
	"time"
)

type ItemsUsecase struct {
	itemsRepo    repository.ItemsRepository
}

func NewItemsUsecase() repository.ItemsRepository {
	return &ItemsUsecase{
		itemsRepo: repository.ItemsRepository(ItemsUsecase{}),
	}
}

func (i ItemsUsecase) CreateItem(item *domain.Item, tx *sql.Tx) (int64, error) {
	sqlStr := "INSERT INTO items (title, description, price, attributes, status_id, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var id int64

	err := tx.QueryRow(sqlStr, item.Title, item.Description, item.Price, item.Attributes, item.StatusId, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (i ItemsUsecase) GetItemById(id int64, db *sql.DB) (domain.Item, error) {
	sqlStr := "SELECT items.id, title, description, price, attributes, picture, status, created_at FROM items JOIN item_statuses i on i.id = items.status_id FULL JOIN pictures p on p.item_id = items.id WHERE items.id = $1"
	var itemData domain.Item
	var s sql.NullString

	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return domain.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&itemData.Id, &itemData.Title, &itemData.Description, &itemData.Price, &itemData.Attributes,
			&s, &itemData.Status, &itemData.CreatedAt)
		if err != nil {
			return domain.Item{}, err
		}

		if s.Valid {
			itemData.Picture = s.String
		} else {
			itemData.Picture = ""
		}

		rows.Close()
		return itemData, nil
	}

	return domain.Item{}, sql.ErrNoRows
}

func (i ItemsUsecase) GetItemsByTitleOrDescription(query string, db *sql.DB) ([]*domain.Item, error) {
	sqlStr := "SELECT items.id, title, description, price, attributes, picture, status, created_at from items JOIN item_statuses i on i.id = items.status_id FULL JOIN pictures p on p.item_id = items.id WHERE title @@ to_tsquery($1) or description @@ to_tsquery($1) ORDER BY created_at DESC"
	var items []*domain.Item
	var s sql.NullString

	rows, err := db.Query(sqlStr, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := new(domain.Item)
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Price, &item.Attributes, &s, &item.Status, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		if s.Valid {
			item.Picture = s.String
		} else {
			item.Picture = ""
		}

		items = append(items, item)
	}

	rows.Close()

	return items, nil
}
