// handlers/auth.go
package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/maheshballi/attendance-monitor/middleware"
)

// DashboardHandler handles requests to the dashboard.
func DashboardHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	// Your implementation for the dashboard handler
	username := r.Context().Value("username").(string)
	w.Write([]byte("Welcome to the dashboard, " + username + "!"))
}

// LogoutHandler handles logout requests.
func LogoutHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	// Perform any cleanup or logout actions here
	middleware.SetUserSession(r, "", store)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
