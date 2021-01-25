package main

import (
	"database/sql"
	"fmt"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"log"
	"net/http"
	"time"
)

func insertInTx(user *User) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		log.Println(err)
	}

	tx, err := db.Begin()
	if err != nil || tx == nil {
		log.Println(err)
	}

	user.Id, err = insertUser(user, tx)
	if err != nil {
		log.Println(err)
		err = tx.Rollback()
	} else {
		log.Println("User inserted")
		err = tx.Commit()
	}
}

func findUser(user *User) bool {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		log.Println(err)
	}

	user.Id, err = findUserByEmailAndPassword(user.Email, user.Password, db)
	if err != nil {
		return false
	}
	return true
}

func setTokenInCookies(user *User, w http.ResponseWriter) {
	token, _ := createToken(user)
	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		Path: "/",
		HttpOnly: true,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}

func auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isValid, _, err := verifyToken(token.Value)
	if err != nil {
		log.Println(err)
	}

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	password, salt := generateHashAndSalt(password)

	user := User{0, email, password, salt, email}
	insertInTx(&user)
	setTokenInCookies(&user, w)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	user := User{0, email, password, "", email}
	if findUser(&user) {
		setTokenInCookies(&user, w)
		log.Println("Success sign in")
	} else {
		log.Println("Wrong email or password")
		w.WriteHeader(http.StatusNoContent)
	}
}

func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println(err)
	}

	isValid, id, err := verifyToken(token.Value)
	if err != nil {
		log.Println(err)
	}

	if isValid {
		_, err = fmt.Fprintf(w, "%v", id)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Println(err)
	}

	log.Println("Started")

	http.Handle("/", http.FileServer(http.Dir("../../../www/")))
	http.HandleFunc("/auth/", auth)
	http.HandleFunc("/register/", signUp)
	http.HandleFunc("/login/", signIn)
	http.HandleFunc("/home/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
