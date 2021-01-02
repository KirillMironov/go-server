package main

import (
	"database/sql"
	"log"
	"net/http"
)

var conf Conf

func insertInTx(user *User) {
	db, err := sql.Open("postgres", conf.Database.ConnectionString)
	if err != nil {
		log.Printf("%v", err)
	}

	tx, err := db.Begin()
	if err != nil || tx == nil {
		log.Printf("%v", err)
	}

	err = insert(user, tx)
	if err != nil {
		log.Printf("%v", err)
		err = tx.Rollback()
	} else {
		log.Printf("User inserted")
		err = tx.Commit()
	}
}

func findUser(user *User) bool {
	db, err := sql.Open("postgres", conf.Database.ConnectionString)
	if err != nil {
		log.Printf("%v", err)
	}

	tx, err := db.Begin()
	if err != nil || tx == nil {
		log.Printf("%v", err)
	}

	users, err := findByEmailAndPassword(user.Email, user.Password, tx)
	if err != nil {
		log.Printf("%v", err)
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}

	return users.Len() > 0
}

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	password, salt := generateHashAndSalt(password)

	user := &User{0, email, password, salt, email}
	insertInTx(user)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	user := &User{0, email, password, "", email}
	if findUser(user) {
		log.Printf("Success sign in")
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("Wrong email or password")
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	ReadConfiguration("service.yaml", &conf)

	log.Println("Started")

	http.Handle("/", http.FileServer(http.Dir("../www/")))
	http.HandleFunc("/register/", signUp)
	http.HandleFunc("/login/", signIn)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
