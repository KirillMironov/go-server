package domain

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Salt string `json:"-"`
	Role string `json:"role"`
}
