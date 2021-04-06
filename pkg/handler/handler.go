package handler

import "net/http"

func InitRoutes() {
	http.HandleFunc("/join", signUp)
	http.HandleFunc("/login", signIn)
	http.Handle("/auth", NewEnsureAuth(nil))
	http.Handle("/getUserData", NewEnsureAuth(getUserData))
	http.Handle("/logout", NewEnsureAuth(logout))
	http.Handle("/changeUsername", NewEnsureAuth(changeUsername))
	http.Handle("/addItem", NewEnsureAuth(addItem))
	http.Handle("/search", NewEnsureAuth(findItems))
	http.Handle("/uploadPicture", NewEnsureAuth(uploadPicture))
}
