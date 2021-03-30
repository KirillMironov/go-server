package main

import (
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/pkg/handler"
	"log"
	"net/http"
)

func main() {
	err := config.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started")

	handler.InitRoutes()
	log.Fatal(http.ListenAndServe(":" + config.Config.Port, http.DefaultServeMux))
}