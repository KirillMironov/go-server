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
