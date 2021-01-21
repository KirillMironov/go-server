package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

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
