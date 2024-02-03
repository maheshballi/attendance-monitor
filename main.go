package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/maheshballi/attendance-monitor/handlers"
	"github.com/maheshballi/attendance-monitor/middleware"
)

var store *sessions.CookieStore

func main() {
	// Create the store once in the main function
	store = sessions.NewCookieStore([]byte("secret-key"))

	r := mux.NewRouter()

	// Allow both GET and POST methods for the /login route
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, store)
	}).Methods("GET", "POST")

	// Example protected route
	r.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.DashboardHandler(w, r, store)
	}))).Methods("GET")

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(w, r, store)
	}).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8081", nil)
}
