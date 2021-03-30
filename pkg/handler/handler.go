package handler

import "net/http"

func InitRoutes() {
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/register", signUp)
	http.HandleFunc("/login", signIn)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/changeUsername", changeUsername)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/findItems", findItems)
	http.HandleFunc("/uploadPicture", uploadPicture)
}
