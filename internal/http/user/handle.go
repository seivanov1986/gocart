package user

import (
	"net/http"

	user2 "github.com/seivanov1986/gocart/internal/service/user"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service user2.Service
}

func New(service user2.Service) *user {
	return &user{
		service: service,
	}
}
