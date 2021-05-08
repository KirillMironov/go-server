package domain

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string
	Salt string
	Role string `json:"role"`
}
