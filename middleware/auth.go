// middleware/auth.go
package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

// AuthMiddleware is a middleware to check user authentication.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Add user information to the request context for later use
		username, ok := session.Values["username"].(string)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SetUserSession sets user session information.
func SetUserSession(r *http.Request, username string, store *sessions.CookieStore) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, nil)
}
