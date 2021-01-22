package main

import (
	"database/sql"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"log"
	"net/http"
	"time"
)

func insertInTx(user *User) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		log.Printf("%v", err)
	}

	tx, err := db.Begin()
	if err != nil || tx == nil {
		log.Printf("%v", err)
	}

	user.Id, err = insertUser(user, tx)
	if err != nil {
		log.Printf("%v", err)
		err = tx.Rollback()
	} else {
		log.Printf("User inserted")
		err = tx.Commit()
	}
}

func findUser(user *User) bool {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		log.Printf("%v", err)
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
		log.Printf("Success sign in")
	} else {
		log.Printf("Wrong email or password")
		w.WriteHeader(http.StatusNoContent)
	}
}

func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token, err := r.Cookie("jwt")
	if err != nil {
		log.Printf("%v", err)
	}

	isValid, id := verifyToken(token.Value)

	if isValid {
		_, err = w.Write(make([]byte, id))
		if err != nil {
			log.Printf("%v", err)
		}
	}
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Started")

	http.Handle("/", http.FileServer(http.Dir("../../../www/")))
	http.HandleFunc("/register/", signUp)
	http.HandleFunc("/login/", signIn)
	http.HandleFunc("/home/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
