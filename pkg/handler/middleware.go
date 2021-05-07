package handler

import (
	"context"
	"github.com/KirillMironov/go-server/config"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"log"
	"net/http"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request)

type EnsureAuth  struct {
	handler AuthenticatedHandler
}

func NewEnsureAuth(handlerToWrap AuthenticatedHandler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}

func (rh *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Access-Control-Allow-Origin", config.Config.Security.AllowedOrigin)
	r.Header.Add("Access-Control-Allow-Methods", "POST")
	r.Header.Add("Access-Control-Allow-Methods", "OPTION")
	r.Header.Add("Content-Type", "application/json")

	token, err := usecase.GetTokenFromCookies(r)
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

	// Create a new request context containing the authenticated user id
	ctxWithId := context.WithValue(r.Context(), "userId", id)
	// Create a new request using that new context
	rWithId:= r.WithContext(ctxWithId)
	// Call the real handler, passing the new request
	rh.handler(w, rWithId)
}
