package middlewares

import (
	"errors"
	"net/http"

	"github.com/aniruddha2000/admybrand-assignment/api/auth"
	"github.com/aniruddha2000/admybrand-assignment/api/responses"
)

// Middleware for the JSON response
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next(rw, r)
	}
}

// Middle for the authentication of user
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(rw, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(rw, r)
	}
}
