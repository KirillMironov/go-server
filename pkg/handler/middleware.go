package handler

import (
	"context"
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
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	token := usecase.GetTokenFromHeader(r)

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
