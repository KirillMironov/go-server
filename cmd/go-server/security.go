package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"time"
)

var jwtKey = []byte("fijASF!saf=342afAS")

type Claims struct {
	Username string
	Id int64
	jwt.StandardClaims
}

func hash(s string) string {
	bytes := []byte(s)
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

func generateHashAndSalt(password string) (string, string) {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	salt := make([]byte, rand.Intn(100 - 50) + 50)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}

	return hash(password + string(salt)), string(salt)
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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Unable to create token")
		return "", err
	}

	return tokenString, nil
}

func verifyToken(token string) (bool, string) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Printf("%v", err)
	}

	return tkn.Valid, claims.Username
}

