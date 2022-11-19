package models

type User struct {
	UserId   int64  `json:"userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
