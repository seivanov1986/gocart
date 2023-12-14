package repository

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/repository/user"
)

type Hub interface {
	User() user.Repository
}

type hub struct {
	user user.Repository
}

func New(db sql_client.DataBase) *hub {
	return &hub{
		user: user.New(db),
	}
}

func (h *hub) User() user.Repository {
	return h.user
}
