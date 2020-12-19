package main

import (
	"database/sql"
	"log"
	"net/http"
)

const (
	connectionString = "postgres://postgres:postgres@35.210.228.180:5432/postgres"
)

func insertInTx(user *User) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Unable to open connection")
	}

	tx, err := db.Begin()
	if err != nil || tx == nil {
		log.Printf("Unable to begin transaction")
		return
	}

	err = insert(user, tx)
	if err != nil {
		log.Printf("Unable to insert User")
		err = tx.Rollback()
		if err != nil {
			log.Printf("Unable to Rollback")
		}
	} else {
		err = tx.Commit()
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	log.Println(email, password)

	user := &User{email, password, email}
	insertInTx(user)
}

func main() {
	log.Println("Started")

	http.Handle("/", http.FileServer(http.Dir("../www/")))
	http.HandleFunc("/register/", signUp)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
