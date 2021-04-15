package handler

import (
	"encoding/json"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"log"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	username := r.URL.Query().Get("username")

	password, salt := usecase.GenerateHashedPasswordAndSalt(password)
	user := &domain.User{Username: username, Email: email, Password: password, Salt: salt}

	id, err := usecase.CreateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user.Id = id

	token, err := usecase.GenerateAuthToken(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = usecase.SetTokenInCookies(token, w)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("User created")
}

func signIn(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	user := domain.User{
		Email:    email,
		Password: password,
	}

	err := usecase.GetUserByEmailAndPassword(&user)
	if err != nil {
		log.Println("Wrong email/password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := usecase.GenerateAuthToken(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = usecase.SetTokenInCookies(token, w)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("Successful login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	usecase.RemoveTokenFromCookies(w)
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userId").(int64)

	user := domain.User{
		Id: id,
	}

	err := usecase.GetUserById(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func changeUsername(w http.ResponseWriter, r *http.Request) {
	newUsername := r.URL.Query().Get("username")
	id := r.Context().Value("userId").(int64)

	err := usecase.UpdateUsername(newUsername, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
