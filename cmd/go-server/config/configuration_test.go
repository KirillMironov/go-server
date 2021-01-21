package config

import (
	"testing"
)

func TestReadConfiguration(t *testing.T)  {
	err := ReadConfig()
	if err != nil {
		t.Fatal("Unable to read config")
	}

	if len(Config.Database.ConnectionString) == 0 {
		t.Fatal("Wrong config")
	}
}
