package models

type User struct {
	ID    int    `json:"id" orm:"auto"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
