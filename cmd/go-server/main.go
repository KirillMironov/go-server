package main

import (
	"database/sql"
	"encoding/json"
	"github.com/KirillMironov/go-server/cmd/go-server/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var currentUser UserData

func insertInTx(user *User) (int64, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	id, err := insertUser(user, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_ = tx.Commit()
	return id, nil
}

func findUser(user *User) (int64, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return 0, err
	}

	id, err := findUserByEmailAndPassword(user.Email, user.Password, db)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getUserData(id int64) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	err = findUserById(id, db)
	if err != nil {
		return err
	}

	return nil
}

func findItemsDb(q string) ([]*Item, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return nil, err
	}

	items, err := findItemsByTitleOrDescription(q, db)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func uploadPictureTx(picture *Picture) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = insertPicture(picture, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func setTokenInCookies(w http.ResponseWriter) error {
	token, err := createToken(&currentUser)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		Path: "/",
		HttpOnly: true,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	return nil
}

func changeUsernameTx(username string) error {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = updateUsername(username, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func insertItemTx(item *Item) (int64, error) {
	db, err := sql.Open("postgres", config.Config.Database.ConnectionString)
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	id, err := insertItem(item, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_ = tx.Commit()
	return id, err
}

func auth(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println("jwt cookie not found. Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isValid, id, err := verifyToken(token.Value)
	if err != nil || !isValid {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = getUserData(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	username := r.URL.Query().Get("username")

	password, salt := generateHashAndSalt(password)
	user := User{0, username, email, password, salt}

	id, err := insertInTx(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	err = getUserData(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = setTokenInCookies(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	log.Println("User inserted")
}

func signIn(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	user := User{0, "", email, password, ""}

	id, err := findUser(&user)
	if err != nil {
		log.Println("Wrong email/password")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = getUserData(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNoContent)
	}

	err = setTokenInCookies(w)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	log.Println("Success sign in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "jwt",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func home(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(currentUser)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
	}
}

func changeUsername(w http.ResponseWriter, r *http.Request)  {
	username := r.URL.Query().Get("username")

	err := changeUsernameTx(username)
	if err != nil {
		log.Println(err)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item = new(Item)

	item.Title = r.URL.Query().Get("title")
	item.Description = r.URL.Query().Get("description")
	item.Price, _ = strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	item.Attributes = r.URL.Query().Get("attributes")
	item.StatusId, _ = strconv.ParseInt(r.URL.Query().Get("statusId"), 10, 64)

	id, err := insertItemTx(item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusRequestEntityTooLarge)
	}

	js, err := json.Marshal(id)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
	}

	log.Println("Item added")
}

func findItems(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	items, err := findItemsDb(q)
	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
	}
}

func uploadPicture(w http.ResponseWriter, r *http.Request)  {
	var picture = new(Picture)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		log.Println(err)
	}

	err = uploadPictureTx(picture)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Println(err)
	}

	log.Println("Started")

	http.Handle("/", http.FileServer(http.Dir("../www/")))
	http.HandleFunc("/auth/", auth)
	http.HandleFunc("/register/", signUp)
	http.HandleFunc("/login/", signIn)
	http.HandleFunc("/logout/", logout)
	http.HandleFunc("/home/", home)
	http.HandleFunc("/changeUsername/", changeUsername)
	http.HandleFunc("/addItem/", addItem)
	http.HandleFunc("/findItems/", findItems)
	http.HandleFunc("/uploadPicture/", uploadPicture)

	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
