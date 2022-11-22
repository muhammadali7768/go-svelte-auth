package middlewares

import (
	"example/books-api/controllers"
	"example/books-api/models"
	"net/http"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *models.User)

type AuthCheck struct {
	handler AuthenticatedHandler
}

func (ea *AuthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := controllers.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}

	ea.handler(w, r, user)
}

func AuthMiddleware(handlerToWrap AuthenticatedHandler) *AuthCheck {
	return &AuthCheck{handlerToWrap}
}
