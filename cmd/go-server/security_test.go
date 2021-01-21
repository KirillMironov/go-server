package main

import (
	"testing"
)

func TestHash(t *testing.T)  {
	str1 := "abc"
	str2 := "homer231233"

	result, salt := generateHashAndSalt(str1)
	if result != hash(str1+salt) {
		t.Fatal("Hashing error")
	}

	result, salt = generateHashAndSalt(str2)
	if result != hash(str2+salt) {
		t.Fatal("Hashing error")
	}
}

func TestCreateToken(t *testing.T) {
	user := &User{0, "Homer", "666", "", "homer@gmail.com"}

	token, err := createToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal(err)
	}
}
