package domain

type User struct {
	Id int64
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Salt string
	Role string
}
