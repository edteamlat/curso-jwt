package main

import (
	"net/http"
)

type Middleware struct {
	jwt JWT
}

func NewMiddleware(jwt JWT) Middleware {
	return Middleware{jwt: jwt}
}

func (m Middleware) ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := m.jwt.Validate(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token no valido"))
			return
		}

		next(w, r)
	}
}

func (m Middleware) ValidateJWTIsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		userData, err := m.jwt.Validate(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token no valido"))
			return
		}

		if !userData.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("no eres admin"))
			return
		}

		next(w, r)
	}
}
