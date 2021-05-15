package domain

import "time"

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
