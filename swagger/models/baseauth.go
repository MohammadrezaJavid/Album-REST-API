package models

type BaseAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
