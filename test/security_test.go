package test

import (
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/service"
	"testing"
)

var s = service.NewSecurityUsecase()

func TestGenerateHashedPasswordAndSalt(t *testing.T) {
	hash, salt := s.GenerateHashedPasswordAndSalt("qwerty123")
	if len(hash) == 0 || len(salt) == 0 {
		t.Fatal("Unable to hash a password")
	}
}

func TestGenerateAuthToken(t *testing.T) {
	user := &domain.User{Username: "Homer", Email: "homer@gmail.com"}

	token, err := s.GenerateAuthToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}
}

func TestVerifyAuthToken(t *testing.T) {
	user := &domain.User{Id: 26, Username: "Marge", Email: "marge@gmail.com"}

	token, err := s.GenerateAuthToken(user)
	if err != nil || len(token) == 0 {
		t.Fatal("Unable to create token")
	}

	id, err := s.VerifyAuthToken(token)
	if id != user.Id || err != nil {
		t.Fatal("Token validation error")
	}

	_, err = s.VerifyAuthToken("abc.abc.abc")
	if err == nil {
		t.Fatal("Token validation error")
	}
}
