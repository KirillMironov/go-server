package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type Item struct {
	Id int64
	Title string
	Description string
	Price float64
	Attributes string
	Picture string
	StatusId int64
	Status string
	CreatedAt time.Time
}

func insertItem(item *Item, tx *sql.Tx) (int64, error) {
	sqlStr := "INSERT INTO items (title, description, price, attributes, status_id, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var id int64

	err := tx.QueryRow(sqlStr, item.Title, item.Description, item.Price, item.Attributes, item.StatusId, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func findItemById(id int64, db *sql.DB) (Item ,error) {
	sqlStr := "SELECT items.id, title, description, price, attributes, picture, status, created_at FROM items JOIN item_statuses i on i.id = items.status_id FULL JOIN pictures p on p.item_id = items.id WHERE items.id = $1"
	var itemData Item
	var s sql.NullString

	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&itemData.Id, &itemData.Title, &itemData.Description, &itemData.Price, &itemData.Attributes,
			&s, &itemData.Status, &itemData.CreatedAt)
		if err != nil {
			return Item{}, err
		}

		if s.Valid {
			itemData.Picture = s.String
		} else {
			itemData.Picture = ""
		}

		rows.Close()
		return itemData, nil
	}

	return Item{}, sql.ErrNoRows
}

func findItemsByTitleOrDescription(q string, db *sql.DB) ([]*Item, error) {
	sqlStr := "SELECT items.id, title, description, price, attributes, picture, status, created_at from items JOIN item_statuses i on i.id = items.status_id FULL JOIN pictures p on p.item_id = items.id WHERE title @@ to_tsquery($1) or description @@ to_tsquery($1) ORDER BY created_at DESC"
	var items []*Item
	var s sql.NullString

	rows, err := db.Query(sqlStr, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := new(Item)
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

	return items, nil
}

