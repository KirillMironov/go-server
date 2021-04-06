package usecase

import (
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	"net/http"
	"time"
)

var s = service.NewSecurityUsecase()

func GenerateHashedPasswordAndSalt(password string) (string, string) {
	return s.GenerateHashedPasswordAndSalt(password)
}

func GenerateAuthToken(user *domain.User) (string, error) {
	return s.GenerateAuthToken(user)
}

func VerifyAuthToken(token string) (int64, error) {
	return s.VerifyAuthToken(token)
}

func SetTokenInCookies(cookieName string, token string, w http.ResponseWriter) error {
	cookie := http.Cookie{
		Name: cookieName,
		Value: token,
		Path: "/",
		HttpOnly: true,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	return nil
}

func GetTokenFromCookies(cookieName string, r *http.Request) (string, error) {
	token, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

func RemoveTokenFromCookies(cookieName string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name: cookieName,
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}
