package handler

import "net/http"

func InitRoutes() {
	http.Handle("/", http.FileServer(http.Dir("../www/")))
	http.HandleFunc("/join", signUp)
	http.HandleFunc("/login", signIn)
	http.Handle("/logout", NewEnsureAuth(logout))
	http.Handle("/auth", NewEnsureAuth(getUserData))
	http.Handle("/changeUsername", NewEnsureAuth(changeUsername))
	http.Handle("/addItem", NewEnsureAuth(addItem))
	http.Handle("/search", NewEnsureAuth(findItems))
	http.Handle("/uploadPicture", NewEnsureAuth(uploadPicture))
}
