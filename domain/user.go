package domain

type User struct {
	Id int64
	Username string
	Email string
	Password string
	Salt string
	Role string
}
