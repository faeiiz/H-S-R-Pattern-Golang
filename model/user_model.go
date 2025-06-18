package model

type Admin struct {
	// ID is the unique identifier for the admin user.
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
