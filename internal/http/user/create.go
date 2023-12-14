package user

import (
	"net/http"
)

type UserCreateRpcIn struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UserCreateRpcOut struct {
	ID *int64
}

type UserCreateError struct {
	Error string
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
}

func validateUserCreate(bodyBytes []byte) error {
	return nil
}
