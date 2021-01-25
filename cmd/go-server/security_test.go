package main

import (
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"testing"
)

func TestCreateToken(t *testing.T) {
	user := &User{0, "Homer", "666", "", "homer@gmail.com"}
	config.Config.Security.JWTKey = jwtKey

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}
}

func TestVerifyToken(t *testing.T) {
	user := &User{26, "Marge", "256", "", "marge@gmail.com"}
	config.Config.Security.JWTKey = jwtKey

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}

	isValid, id, err := verifyToken(token)
	if isValid == false || id != user.Id || err != nil {
		t.Fatal("Token validation error")
	}

	isValid, _, err = verifyToken(invalidToken)
	if isValid == true || err == nil {
		t.Fatal("Token validation error")
	}
}
