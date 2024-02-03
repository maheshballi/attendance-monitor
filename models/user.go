// models/models.go (or wherever you define your user data)
package models

// User represents a user in the system
type User struct {
	Username string
	Password string
}

// Users map stores user data
var Users = map[string]User{
	"user1": {"user1", "password1"},
	"user2": {"user2", "password2"},
	// Add more users as needed
}
