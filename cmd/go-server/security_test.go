package main

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	user := &User{0, "Homer", "666", "", "homer@gmail.com"}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}
}

func TestVerifyToken(t *testing.T) {
	user := &User{26, "Marge", "256", "", "marge@gmail.com"}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}

	isValid, id := verifyToken(token)
	if isValid == false || id != user.Id {
		t.Fatal("Token validation error")
	}

	wrongToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IkhvbWVyIiwiSWQiOjAsImV4cCI6MTYxMTMzNDMyNX0.2aFT3DJVekoGONKHI1S-Ga0aKXqs_zTCrS54fsyutZQ"
	isValid, id = verifyToken(wrongToken)
	if isValid == true {
		t.Fatal("Token validation error")
	}
}
