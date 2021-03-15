package main

import (
	"database/sql"
)

type Picture struct {
	Id int64
	ItemId int64 `json:",string"`
	Picture string
}

func insertPicture(picture *Picture, tx *sql.Tx) error {
	sqlStr := "INSERT INTO pictures (item_id, picture) VALUES ($1, $2)"

	_, err := tx.Exec(sqlStr, picture.ItemId, picture.Picture)
	if err != nil {
		return err
	}

	return nil
}

