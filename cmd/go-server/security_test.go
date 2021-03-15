package main

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	user := &UserData{0, "Homer", "homer@gmail.com", ""}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}
}

func TestVerifyToken(t *testing.T) {
	user := &UserData{26, "Marge", "marge@gmail.com", ""}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}

	isValid, id, err := verifyToken(token)
	if isValid == false || id != user.Id || err != nil {
		t.Fatal("Token validation error")
	}

	isValid, _, err = verifyToken("saf.saf.saf")
	if isValid == true || err == nil {
		t.Fatal("Token validation error")
	}
}
