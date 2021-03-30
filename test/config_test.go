package test

import (
	"github.com/KirillMironov/go-server/config"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	err := config.LoadConfiguration()
	if err != nil {
		t.Fatalf("Unable to read configuration: %v", err)
	}

	if len(config.Config.Port) == 0 {
		t.Fatal("Wrong configuration")
	}
	if len(config.Config.Database.ConnectionString) == 0 {
		t.Fatal("Wrong configuration")
	}
	if len(config.Config.Security.JWTKey) == 0 {
		t.Fatal("Wrong configuration")
	}
}
