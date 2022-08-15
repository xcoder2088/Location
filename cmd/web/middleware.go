package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteInConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("page is beens Hit ")
		next.ServeHTTP(w, r)
	})
}

// Nosurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and save session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}
