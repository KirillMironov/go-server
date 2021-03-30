package handler

import (
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"log"
	"net/http"
)

func auth(w http.ResponseWriter, r *http.Request) {
	token, err := usecase.GetTokenFromCookies("jwt", r)
	if err != nil {
		log.Println("JWT token not found. Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, err := usecase.VerifyAuthToken(token)
	if err != nil {
		log.Println("JWT token is not valid / expired. Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO How to use this Id?
	log.Println(id)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	username := r.URL.Query().Get("username")

	password, salt := usecase.GenerateHashedPasswordAndSalt(password)
	user := &domain.User{Username: username, Email: email, Password: password, Salt: salt}

	id, err := usecase.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}
	user.Id = id

	token, err := usecase.GenerateAuthToken(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	err = usecase.SetTokenInCookies("jwt", token, w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
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
		w.WriteHeader(http.StatusNoContent)
		return
	}

	token, err := usecase.GenerateAuthToken(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	err = usecase.SetTokenInCookies("jwt", token, w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	log.Println("Successful login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	usecase.RemoveTokenFromCookies("jwt", w)
}

func changeUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// TODO How to get User data for UpdateUsername function?
	err := usecase.UpdateUsername(username, &domain.User{Id: 1})
	if err != nil {
		log.Println(err)
	}
}
