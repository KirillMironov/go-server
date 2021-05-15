package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type AuthClaims struct {
	Username string
	Id int64
	jwt.StandardClaims
}

type SecurityUsecase struct {
	securityRepo    repository.SecurityRepository
}

func NewSecurityUsecase() repository.SecurityRepository {
	return &SecurityUsecase{
		securityRepo: repository.SecurityRepository(SecurityUsecase{}),
	}
}

func (s SecurityUsecase) GenerateHashedPasswordAndSalt(password string) (string, string) {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	salt := make([]byte, rand.Intn(100 - 50) + 50)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}

	hash := sha256.Sum256([]byte(password + string(salt)))

	return hex.EncodeToString(hash[:]), string(salt)
}

func (s SecurityUsecase) GenerateAuthToken(user *domain.User) (string, error) {
	claims := &AuthClaims{
		Username: user.Username,
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.Security.JWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s SecurityUsecase) VerifyAuthToken(token string) (int64, error) {
	claims := &AuthClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Security.JWTKey), nil
	})
	if err != nil {
		return 0, err
	}

	return claims.Id, nil
}
