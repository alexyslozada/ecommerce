package model

// Login model of table logins
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
