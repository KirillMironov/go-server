package usecase

import (
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	"net/http"
	"strings"
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

func SetTokenInCookies(token string, w http.ResponseWriter) error {
	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		Path: "/",
		Expires: time.Now().Add(24 * time.Hour),
		Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return nil
}

func GetTokenFromCookies(r *http.Request) (string, error) {
	token, err := r.Cookie("jwt")
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return "", http.ErrAbortHandler
	}

	return strings.TrimSpace(splitToken[1]), nil
}

func RemoveTokenFromCookies(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name: "jwt",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}
