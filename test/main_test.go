package test

import (
	"github.com/KirillMironov/go-server/config"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	err := config.LoadConfiguration()
	if err != nil {
		log.Println(err)
	}

	m.Run()
}
