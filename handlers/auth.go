// handlers/auth.go
package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/maheshballi/attendance-monitor/middleware"
)

// ...

// Pass the store variable to LoginHandler
func LoginHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	// Get username from the form or any other source
	username := r.FormValue("username")

	// Perform your authentication logic

	// Set user session
	middleware.SetUserSession(r, username, store)

	// Redirect or perform other actions as needed
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
