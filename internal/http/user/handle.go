package user

import (
	"net/http"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type user struct {
}

func New() *user {
	return &user{}
}
