package handler

import (
	"encoding/json"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"io/ioutil"
	"log"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var credentials domain.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = json.Unmarshal(body, &credentials)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	password, salt := usecase.GenerateHashedPasswordAndSalt(credentials.Password)
	user := &domain.User{Username: credentials.Username, Email: credentials.Email, Password: password, Salt: salt}

	user.Id, err = usecase.CreateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//token, err := usecase.GenerateAuthToken(user)
	//if err != nil {
	//	log.Println(err)
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(user)
	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("User created")
}

func signIn(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	err := json.Unmarshal(body, &keyVal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	email := keyVal["email"]
	password := keyVal["password"]

	user := domain.User{
		Email:    email,
		Password: password,
	}

	err = usecase.GetUserByEmailAndPassword(&user)
	if err != nil {
		log.Println("Wrong email/password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//token, err := usecase.GenerateAuthToken(&user)
	//if err != nil {
	//	log.Println(err)
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(user)
	_, err = w.Write(js)
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
