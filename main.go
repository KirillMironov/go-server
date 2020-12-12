package main

import (
	"log"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	log.Println(email, password)
}

func main() {
	log.Println("Started")

	http.Handle("/", http.FileServer(http.Dir("../www/")))
	http.HandleFunc("/register/", signUp)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
