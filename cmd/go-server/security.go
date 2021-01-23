package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"time"
)

type Claims struct {
	Username string
	Id int64
	jwt.StandardClaims
}

func generateHashAndSalt(password string) (string, string) {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	salt := make([]byte, rand.Intn(100 - 50) + 50)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}

	hash := sha256.Sum256([]byte(password + string(salt)))

	return hex.EncodeToString(hash[:]), string(salt)
}

func createToken(user *User) (string, error) {
	claims := &Claims{
		Username: user.Username,
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.Security.JWTKey))
	if err != nil {
		log.Printf("Unable to create token")
		return "", err
	}

	return tokenString, nil
}

func verifyToken(token string) (bool, int64) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Security.JWTKey), nil
	})
	if err != nil {
		log.Printf("%v", err)
	}

	return tkn.Valid, claims.Id
}

