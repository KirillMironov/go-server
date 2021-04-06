package domain

type Picture struct {
	Id int64
	ItemId int64 `json:",string"`
	Picture string
}
