package main

import (
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	err := config.ReadConfig()
	if err != nil {
		log.Println(err)
	}
	m.Run()
}
